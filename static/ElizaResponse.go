// Data Reprensentation and Query Project 
// Author: Garry Cummins


package main
import (
    "fmt"
    "net/http"
    "math/rand"
    "regexp"
    
)

func userinputhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s, how are you today? ", r.URL.Query().Get("value")) //.Path[1:])
    uInput := r.URL.Query().Get("userInput")
    response := ElizaResponse(uInput)
    fmt.Fprintf(w, response)
}// User

func ElizaResponse(input string) string {

   //Default responses for Eliza if input doesn't contain a keyword I am looking for 
   responses := []string{
      "I'm not sure what you’re trying to say. Could you explain it to me?",
      "How does that make you feel?",
      "Why do you say that?",
    }

     if matched, _ := regexp.MatchString(`(?i).*\bfather\b.*`, input); matched {
        return "Why don’t you tell me more about your father?"
    }

     if matched, _ := regexp.MatchString(`(?i).*\bmother\b.*`, input); matched {
        return "Why don’t you tell me more about your ?"
    }

     if matched, _ := regexp.MatchString(`(?i).*\bquestion\b.*`, input); matched {
        return "Tell me about where you come from?"
    }

    return responses[rand.Intn(len(responses))]
}// Eliza

func main() {
    
   // rand.Seed(time.Now().UTC().UnixNano())

    dir := http.Dir("./static")
    fileServer := http.FileServer(dir)

    
	http.Handle("/", fileServer)

	http.HandleFunc("/uInput", userinputhandler)
	// 127.0.1.8080
	http.ListenAndServe(":8080", nil)
} 