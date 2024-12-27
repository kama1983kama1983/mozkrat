package control

import (
    "net/http"
    "log"
    "mozkrat/middleware"
)

// Cases retrieves all cases from the database
func Cases() []map[string]interface{} {
    records, err := middleware.DbGetAll("cases")
    if err != nil {
        log.Println("هناك خطأ فى البيانات!")
        return []map[string]interface{}{} // Return an empty slice if there's an error
    }
    return records
}

// CreateDaw3 renders the form to create a new Daw3
func CreateDaw3(w http.ResponseWriter, r *http.Request) {
    data := Cases() // Get cases data
    if len(data) == 0 {
        http.Error(w, "لا توجد بيانات متاحة", http.StatusNotFound)
        return
    }
    middleware.RenderTemplate(w,"da3ormozAdd",data)   
}

// AddDaw3 handles the form submission to add a new Daw3
func AddDaw3(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        fields := []string{"title", "details", "typeof", "caseid", "notes"}
        values := map[string]interface{}{
            "title":    r.FormValue("title"),
            "details":  r.FormValue("details"),
            "typeof":   r.FormValue("typeof"),
            "caseid":   r.FormValue("caseid"),
            "notes":    r.FormValue("notes"),
        }
        result := middleware.Dbinsert("da3ormoz", fields, values)
        if result {
            http.Redirect(w, r, "/GetAllD3w", http.StatusSeeOther)
            return
        }
    }
    http.Error(w, "Method not allowed, please use POST", http.StatusMethodNotAllowed)
}

// EditDaw3 retrieves the Daw3 record for editing
func EditDaw3(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id")
    d3we, err := middleware.DbGetOne("da3ormoz", id)
    if err != nil {
        log.Printf("هناك خطأ %v", err)
        http.Error(w, "لا يوجد بيانات", http.StatusNotFound)
        return
    }
    middleware.RenderTemplate(w,"da3ormoz-edit",d3we)
}

// UpDaw3 handles the form submission to update an existing Daw3
func UpDaw3(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
    id := r.FormValue("id") // Get the ID from the form
    if id == "" {
        http.Error(w, "ID is required for updating", http.StatusBadRequest)
        return
    }
        fields := []string{"title", "details", "typeof", "caseid", "notes"}
        values := map[string]interface{}{
            "title":    r.FormValue("title"),
            "details":  r.FormValue("details"),
            "typeof":   r.FormValue("typeof"),
            "caseid":   r.FormValue("caseid"),
            "notes":    r.FormValue("notes"),
        }
        result := middleware.DbUpdate("da3ormoz", id, fields, values)
        if result {
            http.Redirect(w,r, "/GetAllD3w", http.StatusSeeOther)
            return
        }
    }
    http.Error(w, "هناك خطأ فى تحديث البيانات", http.StatusInternalServerError)
}

// GetAllD3w retrieves and displays all Daw3 records
func GetAllD3w(w http.ResponseWriter, r *http.Request) {
    data, err := middleware.DbGetAll("da3ormoz")
    if err != nil {
        http.Error(w, "هناك خطأ فى عرض التمبلت 1", http.StatusInternalServerError)
        log.Printf("Error parsing template: %v", err)
        return
    }
    middleware.RenderTemplate(w,"da3ormozView",data)
}

// DelDaw3 deletes a specific Daw3 record
func DelDaw3(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id")
    if id == "" {
        http.Error(w, "رقم التعريف غير موجود", http.StatusBadRequest)
        return
    }
    result := middleware.DbDelete("da3ormoz", id)
    if result {
        http.Redirect(w,r, "/GetAllD3w", http.StatusSeeOther)
    } else {
        http.Error(w, "هناك خطأ فى حذف البيانات", http.StatusInternalServerError)
    }
}