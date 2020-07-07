package main

<<<<<<< HEAD
import (
	"log"
	"net"
	"net/http"

	"github.com/istherepie/dropoff/internal/webserver"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")

	if err != nil {
		log.Fatal(err)
	}

	server := webserver.New()

	serveErr := http.Serve(listener, server)

	log.Fatal(serveErr)
=======
import "fmt"

func main() {
	fmt.Println("Begin project: Dropoff")
>>>>>>> 16d9b83... Setup go project
}
