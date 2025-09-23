// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// type msg string

// func (m msg) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
// 	fmt.Fprint(resp, m)
// }

// func main() {
// 	var msgHandler http.Handler = msg("Hello from Web Server in Go")
// 	fmt.Println("Server is listening...")
// 	http.ListenAndServe("localhost:8181", msgHandler)

// }

package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "Hello World!") })
	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
}
