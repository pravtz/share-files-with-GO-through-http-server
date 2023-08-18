package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Uso: go run main.go <directory> <port>")
		os.Exit(1)
	}
	httpDir := os.Args[1]
	port := os.Args[2]
	fs := http.FileServer(http.Dir(httpDir))
	fmt.Printf("server http run in port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, fs))
}
