package main
 
import (
    "fmt"
    "net/http"
)
 

 func main() {
        // handle route using handler function
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello World")
    })
 
        // listen to port
    http.ListenAndServe(":5050", nil)
}
