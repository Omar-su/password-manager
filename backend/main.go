package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type PassWordRequest struct {
	Password string `json:"password"`;
	UserName string `json:"username"`;
	Site string `json:"site"`
}

func passwordHandler(w http.ResponseWriter, r *http.Request){

	// CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")

	// Handle preflight request
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	
	// Check if allowed req method
	if r.Method != http.MethodPost {
		http.Error(w, "Only post allowed", http.StatusMethodNotAllowed)
		return
	}


	var pwr PassWordRequest
	err := json.NewDecoder(r.Body).Decode(&pwr)
	if err != nil {
		http.Error(w,"Invalid body", http.StatusBadRequest)
		return
	}
	
	file, err := os.OpenFile("passwords.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		http.Error(w, "Failed to open file", http.StatusInternalServerError)
		return
	}

	defer file.Close()

	entry := map[string]string{
		"username": pwr.UserName,
		"site":     pwr.Site,
		"password": pwr.Password,
	}
	
	jsonEntry, _ := json.Marshal(entry)
	
	if _, err = file.WriteString(string(jsonEntry) + "\n"); err != nil {
		http.Error(w, "Failed to write to the file", http.StatusInternalServerError)
		return
	}

	fmt.Println("Recieved request")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "Password saved!!"})

}


func main() {

	http.HandleFunc("/api/pw", passwordHandler)

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
