package maria

import (
	"log"
	"net/http"

	"github.com/car1nhanha/go-blog/internal/rose"
	"github.com/car1nhanha/go-blog/internal/sina"
	"github.com/gin-gonic/gin"
)

func GetFiles(context *gin.Context) {
	user := context.Param("user")
	log.Printf("\033[32mmaria - GetFiles %s\033[0m\n", user)
	files, err := rose.FindFilesFromDB(user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, files)
}

func GetFileFromGithub(context *gin.Context) {
	params := &rose.IFindFileFromGithub{}
	// extrair o IFindFileFromGithub // User, Repo, Path
	params.User = context.Param("user")
	params.Repo = context.Param("repo")
	params.Path = context.Query("path")

	sina.SyncUserFiles(params.User)

	file, err := rose.FindFileFromGithub(params)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, file)
}
