// package main

// import (
// 	"fmt"
// 	"github.com/tyler-sommer/stick"
// 	"log"
// 	"net/http"
// 	"os"
// )

// func check(e error) {
// 	// Log message with specified arguments.

// 	if e != nil {
// 		panic(e)
// 	}
// }

// func getPort() string {
// 	// Get the Port from the environment so we can run on Heroku (more of this later)
// 	var port = os.Getenv("PORT")

// 	// Set a default port if there is nothing in the environment
// 	if port == "" {
// 		port = "80"
// 		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
// 	}

// 	return ":" + port
// }

// func main() {
// 	fsRoot, err := os.Getwd() // Templates are loaded relative to this directorys.
// 	check(err)
// 	env := stick.New(stick.NewFilesystemLoader(fsRoot))
// 	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
// 		env.Execute("bar.html", w, nil)
// 	})

// 	fmt.Println("Listening...")
// 	log.Fatal(http.ListenAndServe(getPort(), nil))
// }
