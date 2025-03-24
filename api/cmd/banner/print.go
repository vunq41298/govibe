package banner

import (
	"log"
	"os"
)

// Print prints the banner to the console
func Print() {
	log.New(os.Stdout, "", log.LstdFlags).Printf("Starting project: %s\n", os.Getenv("PROJECT_NAME"))
}
