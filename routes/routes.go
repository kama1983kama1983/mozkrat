package routes

import (
	"log"
	"mozkrat/control"
	"net/http"
)

func MainHandler() {
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/", control.Dashboard) // Assuming you have a Dashboard function

	log.Println("Starting Server on http://localhost:8777")

	err := http.ListenAndServe(":8777", nil)

	if err != nil {

		log.Fatalf("Error starting server: %v", err)

	}

}

// MoadHandler sets up routes for Moad operations

func MoadHandler() {

	http.HandleFunc("/AddMoad", control.AddMoad)

	http.HandleFunc("/GetAllMoad", control.GetAllMoad)

	http.HandleFunc("/EditMoad", control.EditMoad)

	http.HandleFunc("/DeleteMoad", control.DelMoad)

}

// QanonHandler sets up routes for Qanon operations

func QanonHandler() {

	http.HandleFunc("/CreateQanon", control.CreateQanon)

	http.HandleFunc("/AddQanon", control.AddQanon)

	http.HandleFunc("/GetAllQanon", control.GetAllQanon)

	http.HandleFunc("/EditQanon", control.EditQanon)

	http.HandleFunc("/UpQanon", control.UpQanon)

	http.HandleFunc("/DelQanon", control.DelQanon)

}

// Daw3OrMozHandler sets up routes for Daw3 operations

func Daw3OrMozHandler() {

	http.HandleFunc("/createDaw3", control.CreateDaw3)

	http.HandleFunc("/AddDaw3", control.AddDaw3)

	http.HandleFunc("/GetAllD3w", control.GetAllD3w)

	http.HandleFunc("/EditDaw3", control.EditDaw3)

	http.HandleFunc("/UpDaw3", control.UpDaw3)

	http.HandleFunc("/DelDaw3", control.DelDaw3)

}

func Al3aqodHandler() {

	http.HandleFunc("/al3qod/CreateAl3aqod", control.CreateAl3aqod)

	http.HandleFunc("/al3qod/AddAl3aqod", control.AddAl3aqod)

	http.HandleFunc("/al3qod/GetAllAl3aqod", control.GetAllAl3aqod)

	http.HandleFunc("/al3qod/DetailsAl3aqod", control.DetailsAl3aqod)

	http.HandleFunc("/al3qod/EditAl3aqod", control.EditAl3aqod)

	http.HandleFunc("/al3qod/UpAl3aqod", control.UpAl3aqod)

	http.HandleFunc("/al3qod/DelAl3aqod", control.DelAl3aqod)

}

func CasesHandler() {
	http.HandleFunc("/cases/add", control.CreateCase)
	http.HandleFunc("/cases/add/submit", control.AddCase)
	http.HandleFunc("/cases", control.GetAllCases)
	http.HandleFunc("/cases/edit", control.EditCase)
	http.HandleFunc("/cases/edit/submit", control.UpCase)
	http.HandleFunc("/cases/delete", control.DeleteCase)
}

// InitializeRoutes initializes all routes

func InitializeRoutes() {

	MoadHandler()

	QanonHandler()

	Daw3OrMozHandler()

	Al3aqodHandler()

	MainHandler()
}
