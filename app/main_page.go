package app

import (
	"ScaleSync/pkg/api"
	"ScaleSync/pkg/database"
	"ScaleSync/pkg/models"
	"encoding/json"
	"fmt"
	"image/color"
	"io"
	"log"
	"net/http"

	//"ScaleSync/app/sign_up"

	"fyne.io/fyne/dialog"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/gorilla/mux"
)

func showMainPage(win fyne.Window) {

	SmallSpacer := canvas.NewText(" ", color.White)
	SmallSpacer.TextSize = 15

	AppName := canvas.NewText("ScaleSync", color.White)
	AppName.TextSize = 45
	AppName.TextStyle.Bold = true
	AppName.Alignment = fyne.TextAlignCenter

	Spacer := canvas.NewText(" ", color.White)
	Spacer.TextSize = 40

	IDEntry := widget.NewEntry()
	IDEntry.SetPlaceHolder("Enter userID")

	PasswordEntry := widget.NewPasswordEntry()
	PasswordEntry.SetPlaceHolder("Enter password")

	sign_in := widget.NewButton("Sign In", func() {
		fetchUsers(IDEntry.Text, PasswordEntry.Text)
	})

	or := canvas.NewText("----------------- OR -----------------", color.White)
	or.TextStyle.Monospace = true
	or.Alignment = fyne.TextAlignCenter

	sign_up := widget.NewButton("Sign Up", func() {
		showSignUpPage(win)
	})

	form := container.NewVBox(SmallSpacer,
		AppName,
		Spacer,
		IDEntry,
		PasswordEntry,
		SmallSpacer,
		sign_in,
		SmallSpacer,
		or,
		SmallSpacer,
		sign_up,
	)

	win.SetContent(form)
}

func Login() {
	database.Pool = database.InitDB() // Initialize the database connection
	defer database.Pool.Close()

	r := mux.NewRouter()

	r.HandleFunc("/users", api.CreateUser).Methods("POST")
	r.HandleFunc("/users", api.GetAllUsers).Methods("GET")

	r.HandleFunc("/users/{id}", api.GetUser).Methods("GET")
	r.HandleFunc("/users/{id}", api.UpdateUser).Methods("PATCH")
	r.HandleFunc("/users/{id}", api.DeleteUser).Methods("DELETE")

	r.HandleFunc("/login", api.LoginUser).Methods("POST")

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func fetchUsers() {
	resp, err := http.Get("http://localhost:8080/users")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		var users []models.User
		json.Unmarshal(body, &users)
		userList := ""
		for _, user := range users {
			userList += fmt.Sprintf("Name: %s, Email: %s\n", user.Name, user.Email)
		}
		dialog.ShowInformation("Users", userList, nil)
	} else {
		dialog.ShowError(fmt.Errorf("Could not fetch users"), nil)
	}
}
