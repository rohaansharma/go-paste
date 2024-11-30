package models

import "time"

type Paste struct {
	ID      string
	Content string
	Created time.Time
}
