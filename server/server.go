package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Game running in http://127.0.0.1:8080/wasm_exec.html")
	if err := http.ListenAndServe(":8080", http.FileServer(http.Dir("./index"))); err != nil {
		log.Fatal(err)
	}
}
