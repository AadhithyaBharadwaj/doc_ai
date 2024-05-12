package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type FormData struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
	Symptoms  string `json:"symptoms"`
}

func submitDetailsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse JSON request body
	var formData FormData
	err := json.NewDecoder(r.Body).Decode(&formData)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Process the data
	fmt.Println("Received details:")
	fmt.Println("First Name:", formData.FirstName)
	fmt.Println("Last Name:", formData.LastName)
	fmt.Println("Age:", formData.Age)
	fmt.Println("Symptoms:", formData.Symptoms)

	// Return a simple response
	response := response(formData)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, response)
}

func response(formData FormData) string {
	if formData.Symptoms == "fine" ||
		formData.Symptoms == "great" ||
		formData.Symptoms == "nice" ||
		formData.Symptoms == "okay" {
		return fmt.Sprintf("Well, %s, I'm glad to hear that you're feeling fine! If you have any concerns or questions in the future, feel free to reach out. Take care!\nThank you for using Doc AI. If you encounter any bugs or errors, please feel free to report to asdfghjk@yahoo.com so that we can squash them.\nIf you feel you need personal advice, please contact the below doctors:\nBabs: 34567894567", formData.FirstName)
	} else {
		return fmt.Sprintf("Well, %s, the solution is to take rest.", formData.FirstName)
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	submitDetailsHandler(w, r)
}

func main() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)
}
