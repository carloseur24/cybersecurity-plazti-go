package main

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github-tracker/github-tracker/models"
	"github-tracker/github-tracker/repository"
	"github-tracker/github-tracker/repository/entity"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestDummy(t *testing.T) {
	c := require.New(t)

	result := 22

	c.Equal(22, result)
}

func TestInsertGitHubWebhook(t *testing.T) {
	c := require.New(t)

	webhook := models.GitHubWebhook{
		Repository: models.Repository{
			FullName: "carloseur24/cybersecurity-platzi-go",
		},
		HeadCommit: models.Commit{
			ID:      "1asdf5a1sd56f1sad5fdfasfd5165as1dfafadsf51f651",
			Message: "Add sample code for testing the handling of github webhooks",
			Author: models.CommitUser{
				Email:    "carloseur24@gmail.com",
				Username: "Carlos",
			},
		},
	}
	body, err := json.Marshal(webhook)

	createdTime := time.Now()

	m := mock.Mock{}
	mockCommit := repository.MockCommit{Mock: &m}

	commit := entity.Commit{
		RepoName:       webhook.Repository.FullName,
		CommitID:       webhook.HeadCommit.ID,
		CommitMessage:  webhook.HeadCommit.Message,
		AuthorUsername: webhook.HeadCommit.Author.Username,
		AuthorEmail:    webhook.HeadCommit.Author.Email,
		Payload:        string(body),
		CreatedAt:      createdTime,
		UpdatedAt:      createdTime,
	}

	ctx := context.Background()
	mockCommit.On("Insert", ctx, &commit).Return(nil)
	err = insertGitHubWebhook(ctx, &mockCommit, webhook, string(body), createdTime)
	c.NoError(err)

	m.AssertExpectations(t)
}
