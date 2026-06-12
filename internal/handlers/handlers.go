package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type Handlers struct {
	db *sql.DB
}

func New(db *sql.DB) *Handlers {
	return &Handlers{db: db}
}

// Healthz indicates if the container is running (Liveness)
func (h *Handlers) Healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// Ready indicates if the app can accept traffic by checking DB connection (Readiness)
func (h *Handlers) Ready(w http.ResponseWriter, r *http.Request) {
	if err := h.db.Ping(); err != nil {
		http.Error(w, "Database not ready", http.StatusServiceUnavailable)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Ready"))
}

func (h *Handlers) GetBookmarks(w http.ResponseWriter, r *http.Request) {
	// TODO: Fetch from h.db
	json.NewEncoder(w).Encode(map[string]string{"status": "impl pending"})
}

func (h *Handlers) CreateBookmark(w http.ResponseWriter, r *http.Request) {
	// TODO: Insert into h.db
	w.WriteHeader(http.StatusCreated)
}