package control

import (
	"log"
	"mozkrat/middleware"
	"net/http"
)

// CreateCase renders the form to create a new case
func CreateCase(w http.ResponseWriter, r *http.Request) {
	middleware.RenderTemplate(w, "cases-add.html", nil)
}

// AddCase handles the form submission to add a new case
func AddCase(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		fields := []string{"mod3", "mod3le", "date_of_session", "tawkel"}
		values := map[string]interface{}{
			"mod3":            r.FormValue("mod3"),
			"mod3le":          r.FormValue("mod3le"),
			"date_of_session": r.FormValue("date_of_session"),
			"tawkel":          r.FormValue("tawkel"),
		}
		result := middleware.Dbinsert("cases", fields, values)
		if result {
			http.Redirect(w, r, "/GetAllCases", http.StatusSeeOther)
			return
		}
	}
	http.Error(w, "هناك خطأ فى إضافة البيانات", http.StatusInternalServerError)
}

// GetAllCases retrieves and displays all cases
func GetAllCases(w http.ResponseWriter, r *http.Request) {
	data, err := middleware.DbGetAll("cases")
	if err != nil {
		http.Error(w, "Error retrieving data", http.StatusInternalServerError)
		return
	}
	middleware.RenderTemplate(w, "cases-view.html", data)
}

// EditCase retrieves the case record for editing
func EditCase(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	caseData, err := middleware.DbGetOne("cases", id)
	if err != nil {
		log.Printf("Error retrieving case: %v", err)
		http.Error(w, "لا يوجد بيانات", http.StatusNotFound)
		return
	}
	middleware.RenderTemplate(w, "cases-edit.html", caseData)
}

// UpCase handles the form submission to update an existing case
func UpCase(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		id := r.FormValue("id")
		fields := []string{"mod3", "mod3le", "date_of_session", "tawkel"}
		values := map[string]interface{}{
			"mod3":            r.FormValue("mod3"),
			"mod3le":          r.FormValue("mod3le"),
			"date_of_session": r.FormValue("date_of_session"),
			"tawkel":          r.FormValue("tawkel"),
		}
		result := middleware.DbUpdate("cases", id, fields, values)
		if result {
			http.Redirect(w, r, "/GetAllCases", http.StatusSeeOther)
			return
		}
	}
	http.Error(w, "هناك خطأ فى تحديث البيانات", http.StatusInternalServerError)
}

// DeleteCase handles the deletion of a case
func DeleteCase(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	result := middleware.DbDelete("cases", id)
	if result {
		http.Redirect(w, r, "/GetAllCases", http.StatusSeeOther)
	} else {
		http.Error(w, "هناك خطأ فى حذف البيانات", http.StatusInternalServerError)
	}
}
