package control

import (
	"log"
	"mozkrat/middleware"
	"net/http"
)

// CreateAl3aqod renders the form to create a new al3aqod
func CreateAl3aqod(w http.ResponseWriter, r *http.Request) {
	if err := middleware.RenderTemplate(w, "al3aqod-add.html", nil); err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		log.Printf("Error rendering template: %v", err)
	}
}

// AddAl3aqod handles the form submission to add a new al3aqod
func AddAl3aqod(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		fields := []string{"title", "details", "Typeof", "caseid", "notes"}
		values := map[string]interface{}{
			"title":           r.FormValue("title"),
			"details":         r.FormValue("details"),
			"date_contract":   r.FormValue("date_contract"),
			"first_side":      r.FormValue("first_side"),
			"secound_side":    r.FormValue("secound_side"),
			"num_of_contract": r.FormValue("num_of_contract"),
			"notes":           r.FormValue("notes"),
		}
		result := middleware.Dbinsert("al3aqod", fields, values)
		if result {
			http.Redirect(w, r, "/GetAllAl3aqod", http.StatusSeeOther)
			return
		}
	}
	http.Error(w, "هناك خطأ فى إضافة البيانات", http.StatusInternalServerError)
}

// DetailsAl3aqod retrieves and displays the details of a specific al3aqod
func DetailsAl3aqod(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	al3aqod, err := middleware.DbGetOne("al3aqod", id)
	if err != nil {
		http.Error(w, "هناك خطأ فى استرجاع البيانات", http.StatusInternalServerError)
		log.Printf("Error retrieving data: %v", err)
		return
	}
	if err := middleware.RenderTemplate(w, "al3aqod-details.html", al3aqod); err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		log.Printf("Error rendering template: %v", err)
	}
}

// EditAl3aqod retrieves the al3aqod record for editing
func EditAl3aqod(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	al3aqod, err := middleware.DbGetOne("al3aqod", id)
	if err != nil {
		log.Printf("هناك خطأ %v", err)
		http.Error(w, "لا يوجد بيانات", http.StatusNotFound)
		return
	}
	if err := middleware.RenderTemplate(w, "al3aqod-edit.html", al3aqod); err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		log.Printf("Error rendering template: %v", err)
	}
}

// UpAl3aqod handles the form submission to update an existing al3aqod
func UpAl3aqod(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		id := r.FormValue("id")
		fields := []string{"title", "details", "Typeof", "caseid", "notes"}
		values := map[string]interface{}{
			"title":           r.FormValue("title"),
			"details":         r.FormValue("details"),
			"date_contract":   r.FormValue("date_contract"),
			"first_side":      r.FormValue("first_side"),
			"secound_side":    r.FormValue("secound_side"),
			"num_of_contract": r.FormValue("num_of_contract"),
			"notes":           r.FormValue("notes"),
		}
		result := middleware.DbUpdate("al3aqod", id, fields, values)
		if result {
			http.Redirect(w, r, "/GetAllAl3aqod", http.StatusSeeOther)
			return
		}
		http.Error(w, "هناك خطأ فى تحديث البيانات", http.StatusInternalServerError)
	}
}

// GetAllAl3aqod retrieves and displays all al3aqod records
func GetAllAl3aqod(w http.ResponseWriter, r *http.Request) {
	data, err := middleware.DbGetAll("al3aqod")
	if err != nil {
		http.Error(w, "هناك خطأ فى استرجاع البيانات", http.StatusInternalServerError)
		log.Printf("Error retrieving data: %v", err)
	}
	if err := middleware.RenderTemplate(w, "al3aqodView.html", data); err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		log.Printf("Error rendering template: %v", err)
	}
}

// DelAl3aqod deletes a specific al3aqod record
func DelAl3aqod(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "رقم التعريف غير موجود", http.StatusBadRequest)
		return
	}
	result := middleware.DbDelete("al3aqod", id)
	if result {
		http.Redirect(w, r, "/GetAllAl3aqod", http.StatusSeeOther)
	} else {
		http.Error(w, "هناك خطأ فى حذف البيانات", http.StatusInternalServerError)
	}
}
