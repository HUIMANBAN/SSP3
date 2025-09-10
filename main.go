package main

import (
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/browser"
	"io/fs"
	"log"
	"net/http"
	"os"
	"ssp/routes"
	"ssp/utils"
)

//go:embed static
var staticFiles embed.FS

func init() {
	utils.EnsureDataDir("ssp-data")
	utils.EnsureDataFile[utils.LoadingData]("ssp-data/loading.json")
	utils.EnsureDataFile[utils.PauseData]("ssp-data/pause.json")
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())

	staticFS, _ := fs.Sub(staticFiles, "static")

	r.StaticFS("/static", http.FS(staticFS))

	//首页
	r.GET("/", func(c *gin.Context) {
		c.Redirect(302, "/mode")
	})

	//程序说明
	r.GET("/readme", func(c *gin.Context) {
		c.FileFromFS("views/README.html", http.FS(staticFS))
	})

	//mode路由分组
	routes.ModeRoutes(r, staticFS)

	//set路由分组
	routes.SetRoutes(r, staticFS)

	//data路由分组
	routes.DataRoutes(r, staticFS)

	//render路由分组
	routes.RenderRoutes(r, staticFS)

	var port = "59137"

	//打开链接
	openErr := browser.OpenURL("http://localhost:" + port)
	if openErr != nil {
		log.Println(openErr)
	}

	utils.PrintAuthorInfo()
	log.Printf("Server is listening on port %s\n", port)
	fmt.Println("按住Ctrl键并点击链接即可打开：http://localhost:" + port)
	fmt.Println("程序说明：http://localhost:" + port + "/readme")

	runErr := r.Run("0.0.0.0:" + port)
	if runErr != nil {
		log.Println(runErr)
		os.Exit(1)
	}
}
