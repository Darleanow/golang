// Command exercise-5a-nethttp exposes a CRUD REST API for a collection of
// items using only the standard library net/http package.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/google/uuid"
)

type Item struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// RWMutex lets multiple readers proceed concurrently while serialising writes -
// appropriate here because reads (GET) outnumber writes.
var (
	mu    sync.RWMutex
	items = []Item{
		{ID: "1", Name: "Clavier mécanique", Description: "Switch Cherry MX Red"},
		{ID: "2", Name: "Souris gaming", Description: "12 000 DPI, sans fil"},
	}
)

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Printf("writeJSON: %v", err)
	}
}

func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]string{"error": msg})
}

func itemsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		mu.RLock()
		defer mu.RUnlock()
		writeJSON(w, http.StatusOK, items)

	case http.MethodPost:
		var item Item
		if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
			writeError(w, http.StatusBadRequest, "corps JSON invalide")
			return
		}
		if item.Name == "" {
			writeError(w, http.StatusBadRequest, "le champ 'name' est requis")
			return
		}
		item.ID = uuid.NewString()

		mu.Lock()
		items = append(items, item)
		mu.Unlock()

		writeJSON(w, http.StatusCreated, item)

	default:
		writeError(w, http.StatusMethodNotAllowed, "méthode non autorisée")
	}
}

func itemByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/items/")
	if id == "" {
		writeError(w, http.StatusBadRequest, "id manquant dans l'URL")
		return
	}

	switch r.Method {
	case http.MethodGet:
		mu.RLock()
		defer mu.RUnlock()
		for _, it := range items {
			if it.ID == id {
				writeJSON(w, http.StatusOK, it)
				return
			}
		}
		writeError(w, http.StatusNotFound, fmt.Sprintf("item '%s' introuvable", id))

	case http.MethodPut:
		var update Item
		if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
			writeError(w, http.StatusBadRequest, "corps JSON invalide")
			return
		}

		mu.Lock()
		defer mu.Unlock()
		for i, it := range items {
			if it.ID == id {
				if update.Name != "" {
					items[i].Name = update.Name
				}
				if update.Description != "" {
					items[i].Description = update.Description
				}
				writeJSON(w, http.StatusOK, items[i])
				return
			}
		}
		writeError(w, http.StatusNotFound, fmt.Sprintf("item '%s' introuvable", id))

	case http.MethodDelete:
		mu.Lock()
		defer mu.Unlock()
		for i, it := range items {
			if it.ID == id {
				items = append(items[:i], items[i+1:]...)
				w.WriteHeader(http.StatusNoContent)
				return
			}
		}
		writeError(w, http.StatusNotFound, fmt.Sprintf("item '%s' introuvable", id))

	default:
		writeError(w, http.StatusMethodNotAllowed, "méthode non autorisée")
	}
}

func main() {
	mux := http.NewServeMux()
	// Trailing slash routes to itemByIDHandler; exact path routes to the collection handler.
	mux.HandleFunc("/items", itemsHandler)
	mux.HandleFunc("/items/", itemByIDHandler)

	addr := ":8080"
	log.Printf("Serveur démarré sur http://localhost%s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("ListenAndServe : %v", err)
	}
}
