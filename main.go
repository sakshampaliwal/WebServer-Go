package main

import (
	"fmt"
	"log"
	"net/http"
)

// func formHandler(w http.ResponseWriter, r *http.Request) { //w is a variable that has been declared to hold a value of type http.ResponseWriter and r is a variable that has been declared to hold a pointer to a value of type http.Request.
// 	if err := r.ParseForm(); err != nil {
// 		fmt.Fprintf(w, "ParseForm() err: %v", err)
// 		return
// 	}
// 	fmt.Fprintf(w, "POST request successful")
// 	name := r.FormValue("name")
// 	address := r.FormValue("address")
// 	fmt.Fprintf(w, "Name = %s\n", name)
// 	fmt.Fprintf(w, "Address = %s\n", address)
// }
// Or use this:

func formHandler(w http.ResponseWriter, r *http.Request) {
	// Try to parse the form data from the request.
	// If there is an error, write an error message to the response and return.
	err := r.ParseForm()
	if err != nil { //nil is used to represent the absence of an error. 
		errorMessage := "Error: Could not parse form data."
		http.Error(w, errorMessage, http.StatusInternalServerError)
		return
	}

	// Extract the values of the "name" and "address" fields from the form data.
	name := r.FormValue("name")
	address := r.FormValue("address")

	// Write a success message and the extracted form values to the response.
	fmt.Fprintf(w, "Form submitted successfully!\n")
	fmt.Fprintf(w, "Name: %s\n", name)
	fmt.Fprintf(w, "Address: %s\n", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" { //checks whether the requested URL path is equal to /hello and if not then return 404 error
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" { //checks whether the HTTP request method is GET. If it's not, it returns a method is not supported error
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello!")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static")) //file server handler will serve static files(like image,css) to the user's browser. The http.Dir("./static") part specifies the directory on the server where the static files are located.
	http.Handle("/", fileServer) //telling the Go HTTP server to associate that handler with a specific URL path pattern.
	http.HandleFunc("/form", formHandler) //when someone visits the /form URL of the website, the formHandler function will be called to handle that request.
	http.HandleFunc("/hello", helloHandler) //when someone visits the /hello URL of the website, the helloHandler function will be called to handle that request.

	fmt.Printf("Starting server at port 8080\n")
	err := http.ListenAndServe(":8080", nil) //nil means that we are using the default handler provided by the http package.
	if err != nil { //nil is used to represent the absence of an error. 
    	log.Fatal(err)
	}

}