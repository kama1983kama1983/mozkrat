package middleware

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

// Helper function to join strings
func Join(arr []string, sep string) string {
	result := ""
	for i, s := range arr {
		if i > 0 {
			result += sep
		}
		result += s
	}
	return result
}

func Err404(err error, w http.ResponseWriter, message string) {
	if err != nil {
		http.Error(w, message, http.StatusInternalServerError)
	}
}

func InitDb() *sql.DB {
	db, err := sql.Open("sqlite3", "./mozkrat.db")
	if err != nil {
		log.Fatalf("لا يوجد اتصال بقاعدة البيانات حاول مرة اخرى : %v", err)
	}
	return db
}

func RenderTemplate(w http.ResponseWriter, templateName string, data interface{}) error {
	// Parse the templates
	t, err := template.ParseFiles("views/base.html", "views/"+templateName+".html")
	if err != nil {
		http.Error(w, "مشكلة فى تنفيذ التمبلت: "+err.Error(), http.StatusInternalServerError)
		return err
	}
	// Execute the template and write to the response
	if err := t.Execute(w, data); err != nil {
		http.Error(w, "مشكلة فى تنفيذ التمبلت: "+err.Error(), http.StatusInternalServerError)
		return err
	}
	return nil
}
