package main

import "net/http"
import "fmt"
import "log"
import "os"

func main() {
	args := os.Args[1:]
	port := ":9000"
	if len(args) > 0 {
		port = ":" + args[0]
	}
	log.Println("Staring listening on", port)
	ConnectToDatabase()
	SetupHTTPRouting()
	log.Fatal(http.ListenAndServe(port, nil))
	fmt.Print("ended")
}
