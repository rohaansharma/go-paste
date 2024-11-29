const API_URL = process.env.REACT_APP_API_URL;

export const createPaste = async (content) => {
    const response = await fetch(`${API_URL}/api/paste`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ content }),
    });
    if (!response.ok) {
        const errorText = await response.text();
        throw new Error(`Error ${response.status}: ${errorText}`);
    }
    return response.json();
};

export const getPaste = async (id) => {
    const response = await fetch(`${API_URL}/api/paste/${id}`);
    if (!response.ok) {
        const errorText = await response.text();
        throw new Error(`Error ${response.status}: ${errorText}`);
    }
    return response.json();
};