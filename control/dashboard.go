package control

import (
	"html/template"
	"log"
	"net/http"
)

func Dashboard(w http.ResponseWriter, h *http.Request) {
	t, err := template.ParseFiles("views/base.html")
	if err != nil {
		log.Printf("Error parsing template: %v", err) // Log the error in English
		http.Error(w, "Error parsing template: "+err.Error(), http.StatusInternalServerError)
		return
	}
	// Execute the template and write to the response
	if err := t.Execute(w, nil); err != nil {
		log.Printf("Error executing template: %v", err) // Log the error in English
		http.Error(w, "Error executing template: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
