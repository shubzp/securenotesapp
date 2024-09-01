package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"securenotesapp.com/internal/database"
	"securenotesapp.com/internal/notes"
	"securenotesapp.com/internal/users"
	"securenotesapp.com/internal/utililty"
)

type Service struct {
	database database.IDatabase
}

func MakeService() *Service {
	return &Service{
		database: database.InitializeDatabase(),
	}
}

func (service *Service) CreateServer() {
	service.RequestHandler()
	log.Println("Creating server")
	http.ListenAndServe(":8082", nil)
}

func (service *Service) RequestHandler() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		service.HealthCheck()
	})

	http.HandleFunc("/v1/note/add", service.noteHandler)
	http.HandleFunc("/v1/note/get", service.noteHandler)
	http.HandleFunc("/v1/user/register", service.userHandler)
	http.HandleFunc("/v1/user/login", service.userHandler)
	http.HandleFunc("/v1/user/update", service.userHandler)
}

func (service *Service) noteHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	responseBody := utililty.MakeResponseBody(w)
	switch path[1] {
	case "v1":
		switch path[3] {
		case "add":
			if r.Method != http.MethodPost {
				responseBody.AddKeyValue("message", "Invalid method")
				responseBody.WriteResponseHeader(http.StatusMethodNotAllowed)
				break
			}
			log.Println("Adding notes")
			type ReqBody struct {
				Content string
			}
			var body ReqBody
			err := json.NewDecoder(r.Body).Decode(&body)
			if err != nil {
				log.Println("error parsing request body", err)
			}
			note := notes.CreateNote(body.Content)
			err = service.database.SaveNote(note.GetNote())
			if err != nil {
				log.Printf("error while saving note to database : %v", err)
			}
			responseBody.AddKeyValue("message", fmt.Sprintf("Note added with %s", note.GetId()))
			responseBody.WriteResponseHeader(http.StatusOK)
		case "get":
			log.Println("Getting notes")
			if r.Method != http.MethodGet {
				responseBody.AddKeyValue("message", "Invalid method")
				responseBody.WriteResponseHeader(http.StatusMethodNotAllowed)
				break
			}
			noteId := r.URL.Query().Get("noteId")
			note, err := service.database.GetNote(noteId)
			if err != nil {
				log.Println("note")
			}
			responseBody.AddKeyValue("message", "Note fetched")
			responseBody.AddKeyValue("noteContent", note.Content)
			responseBody.WriteResponseHeader(http.StatusOK)
		default:
			fmt.Println("Path not defined")
		}
	}
	responseBody.WriteResponse()
}

func (service *Service) userHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	responseBody := utililty.MakeResponseBody(w)
	switch path[1] {
	case "v1":
		switch path[3] {
		case "register":
			if r.Method != http.MethodPost {
				responseBody.AddKeyValue("message", "Invalid method")
				responseBody.WriteResponseHeader(http.StatusMethodNotAllowed)
				break
			}
			type UserReqBody struct {
				Name     string
				PhoneNo  string
				UserName string
				Password string
			}
			var userDetails UserReqBody
			err := json.NewDecoder(r.Body).Decode(&userDetails)
			if err != nil {
				log.Println("error decoding request body")
				responseBody.AddKeyValue("message", "some problem with request body")
				responseBody.WriteResponseHeader(http.StatusMethodNotAllowed)
				break
			}
			user := users.MakeNewUser(userDetails.Name, userDetails.PhoneNo, userDetails.UserName, userDetails.Password)
			service.database.AddUser(user.GetUser())
			responseBody.AddKeyValue("message", "user added successfully")
			responseBody.WriteResponseHeader(http.StatusOK)
		case "login":
			if r.Method != http.MethodGet {
				responseBody.AddKeyValue("message", "Invalid method")
				responseBody.WriteResponseHeader(http.StatusMethodNotAllowed)
				break
			}
			username := r.URL.Query().Get("username")
			password := r.URL.Query().Get("password")
			areCredsOK, err := service.database.VerifyUserCredentials(username, password)
			if err != nil {
				log.Printf("error verifying user credentials : %v", err)
				responseBody.AddKeyValue("message", "error verifying user credential")
				responseBody.WriteResponseHeader(http.StatusOK)
				break
			}
			if !areCredsOK {
				responseBody.AddKeyValue("message", "wrong credentials")
				responseBody.WriteResponseHeader(http.StatusOK)
				break
			}
			responseBody.AddKeyValue("token", "some token")
			responseBody.AddKeyValue("message", "login successful")
			responseBody.WriteResponseHeader(http.StatusOK)
		case "update":
			if r.Method != http.MethodPost {
				responseBody.AddKeyValue("message", "Invalid method")
				responseBody.WriteResponseHeader(http.StatusMethodNotAllowed)
				break
			}
		default:
			fmt.Println("Path not defined")
		}
	}
	responseBody.WriteResponse()
}

func (svc *Service) HealthCheck() {
	fmt.Println("health is ok")
}
