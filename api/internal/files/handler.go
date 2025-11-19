package files

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetFileHandler(context *gin.Context) {
	user := context.Param("user")
	repo := context.Param("repo")
	path := context.Param("path")

	content, err := GetFile(user, repo, path)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
	}

	context.JSON(http.StatusOK, content)
}

func ListFiles(context *gin.Context) {
	user := context.Param("user")

	content, err := FindOnGithub(user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
	}

	context.JSON(http.StatusOK, content)
}
