package main

// continuing with the "Hello World" Application
import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

// Defining Port Number globally:
const portNumber = ":8080"

// HomePage is the About Page Handler:
func HomePage(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "This is the Home Page!")
}

// DividePage is the Divide Page Handler:
func DividePage(w http.ResponseWriter, req *http.Request) {
	// calling a function which does division:
	divide, err := divdeValues(100.0, 0.0)
	if err != nil {
		fmt.Fprintf(w, "Cannot divide by zero!")
		return // as we want to stop executing this function as soon as the error occurs
	}
	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", 100.0, 0.0, divide))
}

func divdeValues(x, y float32) (float32, error) {
	// checking for error here:
	if y <= 0 {
		err := errors.New("Cannot divide by zero!")
		return 0, err
	}

	result := x / y
	return result, nil
}

func main() {

	// creating HandleFunctions for the above handler functions:
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/divide", DividePage)

	log.Printf("Server listing on PORT %s", portNumber)

	// listening to the server:
	_ = http.ListenAndServe(portNumber, nil)

}
