package routing

import (
	"fmt"
	"os"
)

func Serve() {
	serverHost := os.Getenv("SERVER_HOST")
	serverPort := os.Getenv("SERVER_PORT")

	fmt.Printf("this is the new port 2 %s", fmt.Sprintf("%s:%s", serverHost, serverPort))
	router.Run(fmt.Sprintf("%s:%s", serverHost, serverPort))
}
