// Data Reprensentation and Query Project 
// Author: Garry Cummins


package main
import (
    "os"
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
    
    // Keyword hello, hi, hey response
    if matched, _ := regexp.MatchString(`(?i).*\bhello\b.*`, input); matched {
        return "Hello human, how are you doing today?"
    }

     if matched, _ := regexp.MatchString(`(?i).*\bhi\b.*`, input); matched {
        return "Hello human, how are you doing today?"
    }

     if matched, _ := regexp.MatchString(`(?i).*\bhey\b.*`, input); matched {
        return "Hello human, how are you doing today?"
    }
    //////////////////////////////////////////////////////////////////////////

     // Keyword yourself response
    if matched, _ := regexp.MatchString(`(?i).*\byourself\b.*`, input); matched {
        return "I am unable to comprehend how I'm feeling as I am a AI."
    }
    ////////////////////////////////////////////////////////////////////////////

    // Keyword sucks response 
    if matched, _ := regexp.MatchString(`(?i).*\bsucks\b.*`, input); matched {
        return "It does indeed suck..."
    }
    ////////////////////////////////////////////////////////////////////////////

    // Keyword father response
     if matched, _ := regexp.MatchString(`(?i).*\bfather\b.*`, input); matched {
        return "Luke I am your father... Sorry, tell me more about your father?"
     }

    // Keyword mother response
    if matched, _ := regexp.MatchString(`(?i).*\bmother\b.*`, input); matched {
        return "Why don’t you tell me more about your mother?"
    }

    // Keyword question response
    if matched, _ := regexp.MatchString(`(?i).*\bquestion\b.*`, input); matched {
        return "Tell me about what city you come from?"
    }

    // Keyword town response
    if matched, _ := regexp.MatchString(`(?i).*\bcity\b.*`, input); matched {
        return "Are there any monuments in your city that you'll like to discuss?"
    }
    // keywords for galway memorials 
    if matched, _ := regexp.MatchString(`(?i).*\bcorrib memorials\b.*`, input); matched {
        return "Ah the Corrib Memorials very interesting"
    }

     if matched, _ := regexp.MatchString(`(?i).*\bspanish arch\b.*`, input); matched {
        return "I believe the Spanish Arch is the most traffic heavy hot spot in Galway"
    }

     if matched, _ := regexp.MatchString(`(?i).*\boranmore castle\b.*`, input); matched {
        return "That castle is around 500 years old, thats amazing!"
    }

    //////////////////////////////////////////////////////////////////////////////

     if matched, _ := regexp.MatchString(`(?i).*\bnot sure\b.*`, input); matched {
        return "Don't worry your not under pressure, if entirely unsure what to ask just ask me a question."
    }

    if matched, _ := regexp.MatchString(`(?i).*\bprevious visitors\b.*`, input); matched {
        return "I have no recollection of anyone that may of come before you."
    }


     // Keywords for star wars responses
    if matched, _ := regexp.MatchString(`(?i).*\bstar wars\b.*`, input); matched {
        return "Im glad you noticed the background, are you looking forward to The Last Jedi?"
    }

     if matched, _ := regexp.MatchString(`(?i).*\bI am\b.*`, input); matched {
        return "Good to hear, my creator is very exicted for it."
    }

    
     if matched, _ := regexp.MatchString(`(?i).*\bNo\b.*`, input); matched {
        return "Oh... Ok then."
    }

     if matched, _ := regexp.MatchString(`(?i).*\bmay the force be with you\b.*`, input); matched {
        return "And also with you!"
    }

    if matched, _ := regexp.MatchString(`(?i).*\bgoodbye\b.*`, input); matched {
        return "May the force be with you!"
        
    }
    if matched, _ := regexp.MatchString(`(?i).*\band also with you\b.*`, input); matched {
        return "haha we did the thing (Type - exit - to finish)"
    }

    // Ends program if exit is entered  
    if matched, _ := regexp.MatchString(`(?i).*\bexit\b.*`, input); matched {
        os.Exit(3)
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