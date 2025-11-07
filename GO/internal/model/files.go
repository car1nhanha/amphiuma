package model

import "time"

type File struct {
	Name       string    `dynamodbav:"name"`
	Path       string    `dynamodbav:"path"`
	URL        string    `dynamodbav:"download_url"`
	RepoId     string    `dynamodbav:"repo_id"`
	OwnerLogin string    `dynamodbav:"owner_login"` // user do github
	OwnerId    string    `dynamodbav:"owner_id"`    // id do user
	CreatedAt  time.Time `dynamodbav:"date"`
}
