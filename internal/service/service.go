package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"securenotesapp.com/internal/notes"
	"securenotesapp.com/internal/utililty"
)

type Service struct {
}

func MakeService() *Service {
	return &Service{}
}

func (service *Service) CreateServer() {
	service.RequestHandler()
	// http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
	// 	service.HealthCheck()
	// })
	fmt.Println("Creating server")
	http.ListenAndServe(":8082", nil)
}

func (service *Service) RequestHandler() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		service.HealthCheck()
	})

	http.HandleFunc("/v1/notes/add", noteHandler)
	http.HandleFunc("/v1/notes/get", noteHandler)
}

func noteHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")

	switch path[1] {
	case "v1":
		switch path[3] {
		case "add":
			fmt.Println("Adding notes")
			type ReqBody struct {
				Content string
			}
			var body ReqBody
			err := json.NewDecoder(r.Body).Decode(&body)
			if err != nil {
				log.Println("error parsing request body", err)
			}
			note := notes.CreateNote(body.Content)
			content, _ := note.Read()
			fmt.Println(content)
			responseBody := utililty.MakeResponseBody(w)
			responseBody.AddKeyValue("status", http.StatusOK)
			responseBody.AddKeyValue("message", "Notes added")
			responseBody.WriteResponse()
		case "get":
			fmt.Println("Getting notes")
		default:
			fmt.Println("Path not defined")
		}
	}
}

func (svc *Service) HealthCheck() {
	fmt.Println("health is ok")
}
