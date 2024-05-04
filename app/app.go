package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pablo-costanzo/go-banking/domain"
	"github.com/pablo-costanzo/go-banking/service"
)

func Start() {
	router := mux.NewRouter()

	//Inyect
	customerRepositoryDb := domain.NewCustomerRepositoryStub()
	ch := CustomerHandlers{service.NewCustomerService(customerRepositoryDb)}

	// define routes
	router.
		HandleFunc("/customers", ch.getAllCustomers).
		Methods(http.MethodGet).
		Name("GetAllCustomers")

	//Start the server
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
