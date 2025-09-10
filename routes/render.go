package routes

import (
	"github.com/gin-gonic/gin"
	"io/fs"
	"net/http"
)

var subPaths = []string{"L-end", "L-endNoCam",
	"L-loading", "L-pause", "L-watermark",
	"P-endNoCam", "P-loading", "P-pause", "P-watermark"}

func contains(subPaths []string, subPath string) bool {
	for _, v := range subPaths {
		if v == subPath {
			return true
		}
	}
	return false
}

func renderMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		subPath := c.Param("subPath")
		if contains(subPaths, subPath) {
			c.Next()
		} else {
			c.AbortWithStatus(404)
		}
	}
}

func RenderRoutes(r *gin.Engine, staticFS fs.FS) {
	renderGroup := r.Group("/render", renderMiddleware())

	renderGroup.GET("/:subPath", func(c *gin.Context) {
		c.FileFromFS("views/render/"+c.Param("subPath")+".html", http.FS(staticFS))
	})
}
