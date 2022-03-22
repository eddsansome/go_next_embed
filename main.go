package main

import (
	"bytes"
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"net/http"
)

//go:embed client/dist
//go:embed client/dist/_next/static
//go:embed client/dist/_next/static/chunks/pages/*.js
//go:embed client/dist/_next/static/*/*.js
var nextFS embed.FS


type User struct {
	Username string `json:"username"`
	Email string `json:"email"`
}

func main() {
	distFS, err := fs.Sub(nextFS, "client/dist")
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/", http.FileServer(http.FS(distFS)))
	http.HandleFunc("/api/users", handleUsers)
	log.Println("Starting server on port 4000")
	http.ListenAndServe(":4000", nil)
}


func handleUsers(w http.ResponseWriter, r *http.Request){
	  log.Println("users endpoint called")
		u := User{Username: "eddie", Email: "eddie@gmail.com"}
		var buf bytes.Buffer
		if err := json.NewEncoder(&buf).Encode(u); err != nil {
			fmt.Fprintln(w, "you done goofed")
			return
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, &buf)
}
