package main

import (
	"context"
	"fmt"
	"github-tracker/github-tracker-app/models"
	"github-tracker/github-tracker-app/repository/entity"
	"io"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func postHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received Post Request")

	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading Request")
		return
	}

	fmt.Println("Received Body: ", string(body))
}

func insertGitHubWebhook(ctx context.Context, repo repository.Commit, webhook models.GitHubWebhook, body string, createdTime time.Time) (err error) {
	commit := entity.Commit{
		RepoName:       webhook.Repository.FullName,
		CommitID:       webhook.HeadCommit.ID,
		CommitMessage:  webhook.HeadCommit.Message,
		AuthorUsername: webhook.HeadCommit.Author.Username,
		AuthorEmail:    webhook.HeadCommit.Author.Email,
		Payload:        body,
		CreatedAt:      createdTime,
		UpdatedAt:      createdTime,
	}
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/hello", postHandler).Methods("POST")

	fmt.Println("Server esta corrindo no porto http://localhost:8080")
	err := http.ListenAndServe(":8080", router)

	if err != nil {
		fmt.Println(err.Error())
	}
}
