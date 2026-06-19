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

var (
	mu    sync.Mutex
	items = []Item{
		{ID: "1", Name: "Clavier mécanique", Description: "Switch Cherry MX Red"},
		{ID: "2", Name: "Souris gaming", Description: "12 000 DPI, sans fil"},
	}
)

func sendJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Printf("sendJSON encode error: %v", err)
	}
}

func itemsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		mu.Lock()
		defer mu.Unlock()
		sendJSON(w, http.StatusOK, items)

	case http.MethodPost:
		var item Item
		if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
			sendJSON(w, http.StatusBadRequest, map[string]string{"error": "corps JSON invalide"})
			return
		}
		if item.Name == "" {
			sendJSON(w, http.StatusBadRequest, map[string]string{"error": "le champ 'name' est requis"})
			return
		}
		item.ID = uuid.NewString()

		mu.Lock()
		items = append(items, item)
		mu.Unlock()

		sendJSON(w, http.StatusCreated, item)

	default:
		sendJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "méthode non autorisée"})
	}
}

func itemByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/items/")
	if id == "" {
		sendJSON(w, http.StatusBadRequest, map[string]string{"error": "id manquant"})
		return
	}

	switch r.Method {
	case http.MethodGet:
		mu.Lock()
		defer mu.Unlock()
		for _, it := range items {
			if it.ID == id {
				sendJSON(w, http.StatusOK, it)
				return
			}
		}
		sendJSON(w, http.StatusNotFound, map[string]string{"error": fmt.Sprintf("item '%s' introuvable", id)})

	case http.MethodPut:
		var update Item
		if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
			sendJSON(w, http.StatusBadRequest, map[string]string{"error": "corps JSON invalide"})
			return
		}

		mu.Lock()
		found := false
		for i, it := range items {
			if it.ID == id {
				if update.Name != "" {
					items[i].Name = update.Name
				}
				if update.Description != "" {
					items[i].Description = update.Description
				}
				result := items[i]
				mu.Unlock()
				sendJSON(w, http.StatusOK, result)
				found = true
				break
			}
		}
		if !found {
			mu.Unlock()
			sendJSON(w, http.StatusNotFound, map[string]string{"error": fmt.Sprintf("item '%s' introuvable", id)})
		}

	case http.MethodDelete:
		mu.Lock()
		for i, it := range items {
			if it.ID == id {
				items = append(items[:i], items[i+1:]...)
				mu.Unlock()
				w.WriteHeader(http.StatusNoContent)
				return
			}
		}
		mu.Unlock()
		sendJSON(w, http.StatusNotFound, map[string]string{"error": fmt.Sprintf("item '%s' introuvable", id)})

	default:
		sendJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "méthode non autorisée"})
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/items", itemsHandler)
	mux.HandleFunc("/items/", itemByIDHandler)

	addr := ":8080"
	log.Printf("Serveur démarré sur http://localhost%s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("ListenAndServe : %v", err)
	}
}
