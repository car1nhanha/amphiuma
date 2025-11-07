package model

type Repo struct {
	Name       string `dynamodbav:"name"`
	FullName   string `dynamodbav:"full_name"`
	OwnerLogin string `dynamodbav:"owner_login"` // user do github
	OwnerId    string `dynamodbav:"owner_id"`    // id do user
}
