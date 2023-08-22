package main

import (
	"fmt"
	auth "github.com/abbot/go-http-auth"
	"log"
	"net/http"
	"os"
)

func Secret(user, realm string) string {
	//simple form as an example
	//generate password https://unix4lyfe.org/crypt/

	if user == "john" {
		// password is "hello"
		return "$1$z/TK7/gF$9.GRY43DWKVvVm8uy13WX0"
	}
	return ""
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Uso: go run main.go <directory> <port>")
		os.Exit(1)
	}
	httpDir := os.Args[1]
	port := os.Args[2]

	authenticator := auth.NewBasicAuthenticator("meuserver.com", Secret)
	http.HandleFunc("/", authenticator.Wrap(func(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
		http.FileServer(http.Dir(httpDir)).ServeHTTP(w, &r.Request)
	}))

	fmt.Printf("server http run in port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
