package routes

import (
	"github.com/gin-gonic/gin"
	"io/fs"
	"net/http"
)

func ModeRoutes(r *gin.Engine, staticFS fs.FS) {
	r.GET("/mode", func(c *gin.Context) {
		c.FileFromFS("views/set/mode.html", http.FS(staticFS))
	})
}
