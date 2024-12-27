package control

import (
    "net/http"
    "log"
    "mozkrat/middleware"
)

// getAllMoad retrieves and displays all Moad records
func GetAllMoad(w http.ResponseWriter, r *http.Request) {
    data, err := middleware.DbGetAll("moad")
    if err != nil {
        http.Error(w, "Error retrieving data", http.StatusInternalServerError)
        return
    }
    middleware.RenderTemplate(w,"moad-view",data)
}

// addMoad handles the form submission to add a new Moad
func AddMoad(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        fields := []string{"mada", "details", "notes"}
        values := map[string]interface{}{
            "mada":    r.FormValue("mada"),
            "details": r.FormValue("details"),
            "notes":   r.FormValue("notes"),
        }

        result := middleware.Dbinsert("moad", fields, values)
        if result {
            http.Redirect(w, r, "/GetAllMoad", http.StatusSeeOther)
            return
        }
    }
    //middleware.RenderTemplate(w,"moad-add")
}

// upMoad retrieves the Moad record for editing
func UpMoad(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id")
    data, err := middleware.DbGetOne("moad", id)
    if err != nil {
        log.Printf("Error retrieving Moad: %v", err)
        http.Error(w, "لا يوجد بيانات", http.StatusNotFound)
        return
    }
    middleware.RenderTemplate(w,"moad-edit",data)
}

// editMoad handles the form submission to update an existing Moad
func EditMoad(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id")
    if r.Method == http.MethodPost {
        fields := []string{"mada", "details", "notes"}
        values := map[string]interface{}{
            "mada":    r.FormValue("mada"),
            "details": r.FormValue("details"),
            "notes":   r.FormValue("notes"),
        }
        result := middleware.DbUpdate("moad", id, fields, values)
        if result {
            http.Redirect(w, r, "/GetAllMoad", http.StatusSeeOther)
            return
        }
    }

    // If not a POST request, render the edit form
    UpMoad(w, r)
}

// DelMoad deletes a specific Moad record
func DelMoad(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id")
    if id == "" {
        http.Error(w, "رقم التعريف غير موجود", http.StatusBadRequest)
        return
    }
    result := middleware.DbDelete("moad", id)
    if result {
        http.Redirect(w,r, "/GetAllMoad", http.StatusSeeOther)
    } else {
        http.Error(w, "هناك خطأ فى حذف البيانات", http.StatusInternalServerError)
    }
}