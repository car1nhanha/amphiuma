package rose

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/car1nhanha/go-blog/internal/db"
	"github.com/car1nhanha/go-blog/internal/model"
	"github.com/car1nhanha/go-blog/internal/sina"
)

type IFindFileFromGithub struct {
	User string
	Repo string
	Path string
}

func FindFilesFromDB(user string) ([]model.File, error) {
	files, err := sina.FindFilesFromDB(user)
	if err != nil {
		return nil, fmt.Errorf("repositórios do usuário %s não encontrados", user)
	}
	if len(files) == 0 {
		sina.SyncUserFiles(user)
		return nil, fmt.Errorf("estou importando os dados para %s. aguarde alguns segundos e recarregue", user)
	}
	return files, nil
}

func FindRepoFromDB(user string) ([]model.Repo, error) {
	// sem serventia por enquanto
	// recebe um userId
	// retorna todos os repositórios de um usuário
	out, err := db.Client.Query(context.Background(), &dynamodb.QueryInput{
		TableName:              aws.String("repos"),
		KeyConditionExpression: aws.String("#u = :owner_login"),
		ExpressionAttributeNames: map[string]string{
			"#u": "owner_login",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":owner_login": &types.AttributeValueMemberS{Value: user},
		},
	})
	if err != nil {
		return nil, err
	}
	if out.Items == nil {
		return nil, fmt.Errorf("repositórios do usuário %s não encontrados", user)
	}
	var repos []model.Repo
	if err := attributevalue.UnmarshalListOfMaps(out.Items, &repos); err != nil {
		return nil, err
	}

	return repos, nil
}

func FindFileFromGithub(param *IFindFileFromGithub) (string, error) {
	var resp *http.Response
	var githubResponse model.IGithubGetFile
	// recebe um user, um repo e um path
	// montar a url
	url := "https://api.github.com/repos/" + param.User + "/" + param.Repo + "/contents/" + param.Path

	// acessar o arquivo
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("erro ao acessar GitHub: %v", err)
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		log.Printf("GitHub retornou erro (%d): %s", resp.StatusCode, string(body))
		return "", fmt.Errorf("GitHub retornou status %d", resp.StatusCode)
	}

	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&githubResponse)

	// decodificar o arquivo
	decodedFile := DecodeBase64(githubResponse.Content)
	// retornar o conteúdo
	return decodedFile, nil
}

func DecodeBase64(file string) string {
	clean := strings.ReplaceAll(file, "\n", "")
	decoded, err := base64.StdEncoding.DecodeString(clean)
	if err != nil {
		log.Printf("erro ao decodificar base64: %v", err)
		return ""
	}
	return string(decoded)
}
