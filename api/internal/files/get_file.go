package files

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/car1nhanha/amphiuma/internal/interfaces"
	"github.com/car1nhanha/amphiuma/pkg/utils"
	"gopkg.in/yaml.v2"
)

type IHeaderFile struct {
	Title       string
	Date        string
	Author      string
	Description string
	File_origin string
}

type IgetHeader struct {
	File   string
	Header *IHeaderFile
}

func GetFile(user, repo, path string) (*IgetHeader, error) {
	var file *interfaces.GithubFindFile
	url := "https://api.github.com/repos/" + user + "/" + repo + "/contents/" + path

	response, err := http.Get(url)
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

	err = json.Unmarshal(bodyBytes, &file)
	if err != nil {
		return nil, err
	}

	decodedFile := utils.DecodeBase64(file.Content)
	return getParts(decodedFile), nil
}

func getParts(file string) *IgetHeader {
	content := &IgetHeader{}
	// pegar a primeira posição dos ---
	firstTraces := strings.Index(file, "---")
	if firstTraces < 0 {
		content.File = file
		return content
	}
	secondTrace := strings.Index(file[firstTraces+3:], "---")
	if secondTrace < 0 {
		content.File = file
		return content
	}

	header := file[firstTraces+3 : secondTrace+3]
	fmt.Println(header)

	content.Header = convertYamlToStruct(header)
	content.File = file[secondTrace+7:]

	return content
}

//nolint:unused // used within this file
func convertYamlToStruct(header string) *IHeaderFile {
	var headerParsed *IHeaderFile

	err := yaml.Unmarshal([]byte(header), &headerParsed)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	return headerParsed
}
