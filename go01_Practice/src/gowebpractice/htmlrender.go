package main
import "github.com/gin-gonic/gin"
import "net/http"

func main() {
	router := gin.Default()


	//使用 LoadHTMLGlob() 或者 LoadHTMLFiles()来渲染模版
	router.LoadHTMLGlob("templates/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})
	router.Run(":8080")
}