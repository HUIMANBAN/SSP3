package routes

import (
	"github.com/gin-gonic/gin"
	"io/fs"
	"log"
	"net/http"
	"os"
	"ssp/utils"
)

func DataRoutes(r *gin.Engine, staticFS fs.FS) {
	dataGroup := r.Group("/data")

	dataGroup.POST("/upload", func(c *gin.Context) {
		file, formErr := c.FormFile("poster")
		if formErr != nil {
			log.Println(formErr)
			return
		}

		saveErr := c.SaveUploadedFile(file, "ssp-data/poster")
		if saveErr != nil {
			log.Println(saveErr)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "UploadSuccessful",
		})
	})

	dataGroup.GET("/poster", func(c *gin.Context) {
		if _, err := os.Stat("ssp-data/poster"); err != nil {
			c.FileFromFS("images/poster.gif", http.FS(staticFS))
			return
		}
		c.File("ssp-data/poster")

	})

	dataGroup.POST("/loading", func(c *gin.Context) {
		var instance utils.LoadingData
		if err := c.BindJSON(&instance); err != nil {
			log.Println(err)
			return
		}

		writeErr := utils.WriteJSONToFile(instance, "ssp-data/loading.json")
		if writeErr != nil {
			log.Println(writeErr)
			c.Status(500)
		}
		c.Status(200)
	})

	dataGroup.POST("/pause", func(c *gin.Context) {
		var instance utils.PauseData

		if err := c.BindJSON(&instance); err != nil {
			log.Println(err)
			return
		}

		writeErr := utils.WriteJSONToFile(instance, "ssp-data/pause.json")
		if writeErr != nil {
			log.Println(writeErr)
			c.Status(500)
		}
		c.Status(200)
	})

	dataGroup.GET("/loading", func(c *gin.Context) {
		instance, _ := utils.ReadJSONFromFile[utils.LoadingData]("ssp-data/loading.json")
		c.JSON(200, instance)
	})

	dataGroup.GET("/pause", func(c *gin.Context) {
		instance, _ := utils.ReadJSONFromFile[utils.PauseData]("ssp-data/pause.json")
		c.JSON(200, instance)
	})
}
