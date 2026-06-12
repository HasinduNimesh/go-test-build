package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"bookmark-api/internal/handlers"
	_ "github.com/lib/pq"
)

func main() {
	// 1. Load config from environment (set by Docker/K8s)
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL environment variable is required")
	}

	// 2. Connect to Database
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	// 3. Initialize Handlers
	h := handlers.New(db)

	// 4. Setup Routing (requires Go 1.22+)
	mux := http.NewServeMux()
	
	// K8s Probes
	mux.HandleFunc("GET /healthz", h.Healthz) // Liveness
	mux.HandleFunc("GET /ready", h.Ready)     // Readiness
	
	// API Endpoints
	mux.HandleFunc("GET /bookmarks", h.GetBookmarks)
	mux.HandleFunc("POST /bookmarks", h.CreateBookmark)

	// 5. Start Server
	port := ":8080"
	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}