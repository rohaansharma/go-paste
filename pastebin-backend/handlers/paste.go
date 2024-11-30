package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"net/http"
	"pastebin-backend/db"
	"pastebin-backend/models"
	"pastebin-backend/utils"
	"strings"
	"time"
)

func CreatePasteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not allowed")
	}

	var req struct {
		Content string `json:"content"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid request body")
		return
	}

	id, err := utils.GenerateID(8)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to generate ID")
		return
	}

	paste := models.Paste{
		ID:      id,
		Content: req.Content,
		Created: time.Now(),
	}

	query := `INSERT INTO pastes (id, content, created_at) VALUES ($1, $2, $3)`

	_, err = db.DB.Exec(context.Background(), query, paste.ID, paste.Content, paste.Created)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to create paste")
		return
	}

	response := map[string]string{"id": paste.ID}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetPasteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, "Method not allowed")
		return
	}

	pathPrefix := "/api/paste/"
	if !strings.HasPrefix(r.URL.Path, pathPrefix) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Invalid paste ID")
		return
	}

	id := strings.TrimPrefix(r.URL.Path, pathPrefix)
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Invalid paste ID")
		return
	}

	query := `SELECT content FROM pastes WHERE id = $1`
	var content string
	err := db.DB.QueryRow(context.Background(), query, id).Scan(&content)
	if err != nil {
		if err == pgx.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "Paste not found")
			return
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "Internal server error")
			log.Printf("Error querying paste: %v\n", err)
			return
		}
	}

	response := map[string]string{"content": content}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
