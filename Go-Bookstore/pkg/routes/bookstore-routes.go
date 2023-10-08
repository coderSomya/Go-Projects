package routes

import(
	"github.com/gorilla/mux"
	"go-bookstore/pkg/controllers"
)


var RegisterBookstoreRoutes = func(router *mux.Router){

	router.HandleFunc("/book", controllers.createBook).Methods("POST")
	router.HandleFunc("/book", controllers.getBooks).Methods("GET")
	router.HandleFunc("/book{bookid}", controllers.getBookById).Methods("GET")
	router.HandleFunc("/book{bookid}", controllers.deleteBook).Methods("DELETE")
	router.HandleFunc("/book{bookid}", controllers.updateBook).Methods("PUT")
}