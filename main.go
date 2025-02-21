package main

import (
	"encoding/json"
	"fmt"
	"go-language/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func isPalindrome(text string) bool {
	text = strings.ToLower(strings.ReplaceAll(text, " ", ""))
	reversed := ""
	for i := len(text) - 1; i >= 0; i-- {
		reversed += string(text[i])
	}
	return text == reversed
}


func palindromeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		text := r.URL.Query().Get("text")
		if text == "" {
			http.Error(w, "Text query parameter is required", http.StatusBadRequest)
			return
		}

		if isPalindrome(text) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Palindrome"))
		} else {
			http.Error(w, "Not palindrome", http.StatusBadRequest)
		}
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello Go developers"))
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}


func GetlanguageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(models.GetLanguages())
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func languageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil || id < 0 || id >= len(models.Languages) {
			http.Error(w, "Language not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(models.Languages[id])
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}


func addLanguageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var newLanguage models.Language
		if err := json.NewDecoder(r.Body).Decode(&newLanguage); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		models.AddLanguage(newLanguage)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newLanguage)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func patchLanguageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPatch {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil || id < 0 || id >= len(models.Languages) {
			http.Error(w, "Language not found", http.StatusNotFound)
			return
		}

		var updatedLanguage models.Language
		if err := json.NewDecoder(r.Body).Decode(&updatedLanguage); err != nil {
			http.Error(w, "Invalid JSON data", http.StatusBadRequest)
			return
		}

		// Mengubah data language di index yang sesuai
		models.Languages[id] = updatedLanguage

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(updatedLanguage)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func deleteLanguageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil || id < 0 || id >= len(models.Languages) {
			http.Error(w, "Language not found", http.StatusNotFound)
			return
		}

		models.Languages = append(models.Languages[:id], models.Languages[id+1:]...)

		w.WriteHeader(http.StatusNoContent)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
func main() {
	r := mux.NewRouter()

	// Menangani route untuk setiap endpoint
	r.HandleFunc("/", helloHandler).Methods(http.MethodGet)
	r.HandleFunc("/language", GetlanguageHandler).Methods(http.MethodGet) 
	r.HandleFunc("/language/{id:[0-9]+}", languageHandler).Methods(http.MethodGet) 
	r.HandleFunc("/palindrome", palindromeHandler).Methods(http.MethodGet)
	r.HandleFunc("/language/add", addLanguageHandler).Methods(http.MethodPost)
	r.HandleFunc("/language/{id:[0-9]+}", patchLanguageHandler).Methods(http.MethodPatch) 
	r.HandleFunc("/language/{id:[0-9]+}", deleteLanguageHandler).Methods(http.MethodDelete) 

	port := ":8080"
	fmt.Println("Server is running on http://localhost" + port)
	if err := http.ListenAndServe(port, r); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
