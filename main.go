// package main

// import (
// 	"ScaleSync/pkg/models"
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"log"
// 	"net/http"
// 	"sync"

// 	"ScaleSync/pkg/api"
// 	"ScaleSync/pkg/database"

// 	"fyne.io/fyne/v2/app"
// 	"fyne.io/fyne/v2/container"
// 	"fyne.io/fyne/v2/dialog"
// 	"fyne.io/fyne/v2/widget"
// 	"github.com/gorilla/mux"
// 	"github.com/joho/godotenv"
// )

// func init() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// }

// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(1)

// 	// Start the HTTP server in a goroutine
// 	go func() {
// 		defer wg.Done()
// 		Login() // Start the server
// 	}()

// 	// Start the GUI in the main goroutine
// 	startGUI()

// 	wg.Wait()
// }

package main

import (
	"ScaleSync/app"
)

func main() {
	app.App()
}
