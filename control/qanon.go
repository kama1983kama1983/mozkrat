package control

import (
	"log"
	"mozkrat/middleware"
	"net/http"
)

// CreateQanon renders the form to create a new Qanon
func CreateQanon(w http.ResponseWriter, r *http.Request) {
	middleware.RenderTemplate(w, "qanon-add", nil)
}

// AddQanon handles the form submission to add a new Qanon
func AddQanon(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		fields := []string{"title", "details", "typeof", "caseid", "notes"}
		values := map[string]interface{}{
			"title":   r.FormValue("title"),
			"details": r.FormValue("details"),
			"typeof":  r.FormValue("typeof"),
			"caseid":  r.FormValue("caseid"),
			"notes":   r.FormValue("notes"),
		}
		result := middleware.Dbinsert("qanon", fields, values)
		if result {
			http.Redirect(w, r, "/GetAllQanon", http.StatusSeeOther)
			return
		}
	}
	http.Error(w, "Method not allowed, please use POST", http.StatusMethodNotAllowed)
}

// DetailsQanon retrieves and displays the details of a specific Qanon
func DetailsQanon(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	qanon, err := middleware.DbGetOne("qanon", id)
	if err != nil {
		log.Printf("هناك خطأ %v", err)
		http.Error(w, "لا يوجد بيانات", http.StatusNotFound)
		return
	}
	middleware.RenderTemplate(w, "qanon-details", qanon)
}

// EditQanon retrieves the Qanon record for editing
func EditQanon(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	qanon, err := middleware.DbGetOne("qanon", id)
	if err != nil {
		log.Printf("هناك خطأ %v", err)
		http.Error(w, "لا يوجد بيانات", http.StatusNotFound)
		return
	}
	middleware.RenderTemplate(w, "qanon-edit", qanon)
}

// UpQanon handles the form submission to update an existing Qanon
func UpQanon(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		id := r.FormValue("id") // Get the ID from the form
		if id == "" {
			http.Error(w, "ID is required for updating", http.StatusBadRequest)
			return
		}
		fields := []string{"title", "date_of_publish"}
		values := map[string]interface{}{
			"title":           r.FormValue("title"),
			"date_of_publish": r.FormValue("date_of_publish"),
		}
		result := middleware.DbUpdate("qanon", id, fields, values)
		if result {
			http.Redirect(w, r, "/GetAllQanon", http.StatusSeeOther)
			return
		}
	}
	http.Error(w, "هناك خطأ فى تحديث البيانات", http.StatusInternalServerError)
}

// GetAllQanon retrieves and displays all Qanon records
func GetAllQanon(w http.ResponseWriter, r *http.Request) {
	data, err := middleware.DbGetAll("qanon")
	if err != nil {
		http.Error(w, "Error retrieving data", http.StatusInternalServerError)
		return
	}
	middleware.RenderTemplate(w, "qanon-view", data)
}

// DelQanon deletes a specific Qanon record
func DelQanon(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "رقم التعريف غير موجود", http.StatusBadRequest)
		return
	}
	result := middleware.DbDelete("qanon", id)
	if result {
		http.Redirect(w, r, "/GetAllQanon", http.StatusSeeOther)
	} else {
		http.Error(w, "هناك خطأ فى حذف البيانات", http.StatusInternalServerError)
	}
}
