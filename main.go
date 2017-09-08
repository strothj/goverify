package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	log.Println("Starting server at http://localhost:" + port)

	http.HandleFunc("/verify", handleRequest)
	http.ListenAndServe(":"+port, nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	email := r.URL.Query().Get("email")

	if email == "" {
		http.Error(w, "You need to pass and email address to verify.", 500)
		return
	}

	verify_result := VerifyResult{Email: email}

	verify_result.Verify()

	json.NewEncoder(w).Encode(verify_result)
}
