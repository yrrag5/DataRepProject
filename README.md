# DataRepProject

## Creating a chatbot guided by the ELIZA program

My name is Garry Cummins, I'm a software development student in GMIT. This chatbot was created for a project as part of the module Data Reprensentation and Query.

## Running Eliza
The first step is to click clone or download located at the top right of this repository. Copy the URl and proceed to your cmd prompt. Now its expected to have go installed, cd into a folder you wish to clone your repo than type the command (git clone url).

Cd into the repo than type the command (go run ElizaResponse.go), allow access than go to localhost:8080 on your browser, the chatbot should now be visable.

## Functionality
ElizaResponse.go serves all files located in the static folder. I combined my CSS and JS in my index.html so only one file exists. I use a jpeg for my chatbot background in the img folder. index uses jquery and ajax in js src in order to be served by the ElizaRespons.go file.

I made the chatbot star wars themed were Eliza will respond to certain quotes or mentions to it with regular expressions (May the force be with you, star wars, etc...). There are also general words such as city, father, hello, that Eliza will catch on to continue the conversation. There are standard responses if Eliza doesn't pick up a keyword. The (exit) keyword is used to end the local host serve and end the conversation.


## Resources

https://golang.org/pkg/regexp/

https://stackoverflow.com/questions/26559557/how-do-you-serve-a-static-html-file-using-a-go-web-server

https://www.w3schools.com/jquery/

https://stackoverflow.com/questions/23197206/how-to-get-ajax-post-request-value-in-go-lang

