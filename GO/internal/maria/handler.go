package maria

import "github.com/gin-gonic/gin"

// acessa
func RegisterRoutes(r *gin.Engine) {
	group := r.Group("/maria")
	{
		group.GET("/:user", GetFiles)
		group.GET("/:user/:repo", GetFileFromGithub)
	}
}
