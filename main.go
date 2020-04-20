package main

import (
	"travel/router"
)

func main() {
	router.Initial()

	// r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
	// 	vars := mux.Vars(r)
	// 	title := vars["title"]
	// 	page := vars["page"]

	// 	fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	// })

	// user := &model.User{
	// 	FullName:    "tony",
	// 	PhoneNumber: "0703022780",
	// 	Password:    "123123",
	// }
	// repository.InsertUser(user)
	// repository.GetUser("0703022780", "123123")
}
