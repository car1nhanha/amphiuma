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
	Description string
	Author      string
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
	parts := getParts(decodedFile)

	parts.Header.Author = user
	parts.Header.File_origin = file.HtmlURL

	return parts, nil
}

func getParts(file string) *IgetHeader {
	content := &IgetHeader{}

	firstTraces := strings.Index(file, "---")
	secondTrace := strings.Index(file[firstTraces+3:], "---")

	if secondTrace < 0 {
		content.File = file
		content.Header.Description = ""
		content.Header.Title = getTitle(content.File)
		return content
	}

	header := file[firstTraces+3 : secondTrace+3]

	content.Header = convertYamlToStruct(header)
	content.File = file[secondTrace+7:]

	return content
}

func convertYamlToStruct(header string) *IHeaderFile {
	var headerParsed *IHeaderFile

	err := yaml.Unmarshal([]byte(header), &headerParsed)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	return headerParsed
}

func getTitle(content string) string {
	lines := strings.Split(content, "\n")

	for _, line := range lines {
		if strings.HasPrefix(line, "# ") {
			return line
		}
	}
	return ""
}
