package routes

import (
	"github.com/gin-gonic/gin"
	"io/fs"
	"net/http"
)

func SetRoutes(r *gin.Engine, staticFS fs.FS) {
	setGroup := r.Group("/set")

	setGroup.GET("/landscape", func(c *gin.Context) {
		c.FileFromFS("views/set/L-set.html", http.FS(staticFS))
	})

	setGroup.GET("/portrait", func(c *gin.Context) {
		c.FileFromFS("views/set/P-set.html", http.FS(staticFS))
	})
}
