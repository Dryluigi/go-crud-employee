package routes

import (
	"database/sql"
	"net/http"

	"github.com/Dryluigi/crud-employee-go/controller"
)

func MapRoutes(server *http.ServeMux, db *sql.DB) {
	server.HandleFunc("/", controller.NewHelloWorldController())
	server.HandleFunc("/employee", controller.NewIndexEmployee(db))
	server.HandleFunc("/employee/create", controller.NewCreateEmployeeController(db))
	server.HandleFunc("/employee/update", controller.NewUpdateEmployeeController(db))
	server.HandleFunc("/employee/delete", controller.NewDeleteEmployeeController(db))
}
