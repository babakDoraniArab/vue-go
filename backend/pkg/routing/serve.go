package routing

import (
	"fmt"
	"os"
)

func Serve() {
	serverHost := os.Getenv("SERVER_HOST")
	serverPort := os.Getenv("SERVER_PORT")
	// Add CORS middleware
	// config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"http://127.0.0.1:5173"} // replace with your frontend URL
	// router.Use(cors.New(config))

	fmt.Printf("this is the new port 2 %s", fmt.Sprintf("%s:%s", serverHost, serverPort))
	router.Run(fmt.Sprintf("%s:%s", serverHost, serverPort))
}
