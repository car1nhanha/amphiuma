package files

import (
	"encoding/json"
	"io"
	"net/http"
)

type OGetUser struct {
	Login        string
	Avatar_url   string
	Url          string
	Name         string
	Public_repos int
	Followers    int
	Following    int
}

func GetUser(user string) (*OGetUser, error) {
	var url = "https://api.github.com/users/" + user
	var finalUser *OGetUser

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		response.Body.Close()
		return nil, err
	}

	bodyBytes, err := io.ReadAll(response.Body)
	response.Body.Close()

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bodyBytes, &finalUser)
	if err != nil {
		return nil, err
	}

	return finalUser, nil
}
