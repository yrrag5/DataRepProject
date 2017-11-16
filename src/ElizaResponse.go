// Data Reprensentation and Query Project 
// Author: Garry Cummins


package main
import (
    "fmt"
    "net/http"
    "math/rand"
    //"time"
    //"regexp"
    
)

func userinputhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Query().Get("value")) //.Path[1:])
}// User

func ElizaResponse(input string) string {

   responses := []string{
      "I'm not sure what you’re trying to say. Could you explain it to me?",
      "How does that make you feel?",
      "Why do you say that?",
    }
    return responses[rand.Intn(len(responses))]

}// Eliza

func main() {
    fs := http.FileServer(http.Dir("src"))
	http.Handle("/", fs)

	http.HandleFunc("/user-input", userinputhandler)
	// 127.0.1.8080
	http.ListenAndServe(":8080", nil)
}