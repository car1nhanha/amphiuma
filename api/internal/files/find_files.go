package files

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/car1nhanha/amphiuma/internal/interfaces"
)

type OFilesList struct {
	Total int
	Items []List
}

type List struct {
	Name     string
	Path     string
	Url      string
	RepoInfo RepoInfo
}

type RepoInfo struct {
	Name        string
	FullName    string
	Description string
}

type OFindOnGithub struct {
	OFilesList
	User OGetUser
}

// buscar todos os arquivos
func FindOnGithub(user, extension string) (*OFindOnGithub, error) {
	var files *interfaces.FindGithubFiles
	url := "https://api.github.com/search/code?q=extension:" + extension + "+user:" + user
	token := os.Getenv("GH_PERSONAL_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("error to get token")
	}

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Authorization", token)
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		response.Body.Close()
		return nil, fmt.Errorf("status not ok")
	}

	bodyBytes, err := io.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bodyBytes, &files)
	if err != nil {
		return nil, err
	}

	parsed := parseFiles(files)

	userFinal, err := GetUser(user)
	if err != nil {
		return nil, err
	}

	return &OFindOnGithub{
		OFilesList: OFilesList{
			Total: parsed.Total,
			Items: parsed.Items,
		},
		User: *userFinal,
	}, nil
}

// formatar uma resposta
func parseFiles(files *interfaces.FindGithubFiles) *OFilesList {
	list := &OFilesList{}
	list.Total = files.TotalCount

	arrList := make([]List, 0, len(files.Items))

	for _, v := range files.Items {
		repoInfo := &RepoInfo{}
		repoInfo.Description = v.Repository.Description
		repoInfo.Name = v.Repository.Name
		repoInfo.FullName = v.Repository.FullName

		arrList = append(arrList, List{
			Name:     v.Name,
			Path:     v.Path,
			Url:      v.URL,
			RepoInfo: *repoInfo,
		})
	}

	list.Items = arrList

	return list
}
