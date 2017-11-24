// Data Reprensentation and Query Project 
// Author: Garry Cummins


package main
import (
    //"os"
    "fmt"
    "net/http"
    "math/rand"
    "regexp"
    
)

func userinputhandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hello %s, how are you today? ", r.URL.Query().Get("value")) //.Path[1:])
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
      "I was wondering that myself",
      "Could you go into further detail? ",
    }

    // Checking for keywords to return unique responses

    // (?i). for flag, \b before and after 
    
    // Keyword hello response
    if matched, _ := regexp.MatchString(`(?i).*\bhello\b.*`, input); matched {
        return "Hello human, how are you doing today?"
     }

    // Keyword father response
     if matched, _ := regexp.MatchString(`(?i).*\bfather\b.*`, input); matched {
        return "Why don’t you tell me more about your father?"
     }

    // Keyword mother response
    if matched, _ := regexp.MatchString(`(?i).*\bmother\b.*`, input); matched {
        return "Why don’t you tell me more about your mother?"
    }

    // Keyword question response
    if matched, _ := regexp.MatchString(`(?i).*\bquestion\b.*`, input); matched {
        return "Tell me about where you come from?"
    }

    // Keyword town response
    if matched, _ := regexp.MatchString(`(?i).*\btown\b.*`, input); matched {
        return "Are there any hotspots in your town that you'll like to discuss?"
    }

     // Keywords star wars response
    if matched, _ := regexp.MatchString(`(?i).*\bstar wars\b.*`, input); matched {
        return "Im glad you noticed the background, are you looking forward to The Last Jedi?"
    }

    


    // Ends program if goodbye is entered 
    if matched, _ := regexp.MatchString(`(?i).*\bgoodbye\b.*`, input); matched {
        return "Thank you for the conversation, until next time!"
        //os.Exit(3))

    }

return responses[rand.Intn(len(responses))]
}// Eliza

func main() {
    
    // Serves files in static folder
    sta := http.Dir("./static")
    fileServer := http.FileServer(sta)

	http.Handle("/", fileServer)

	http.HandleFunc("/uInput", userinputhandler)
	// 127.0.1.8080
	http.ListenAndServe(":8080", nil)
} 