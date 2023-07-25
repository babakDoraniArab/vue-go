package html

import (
	"github.com/babakDoraniArab/vue-go/internal/provider/view"
	"github.com/gin-gonic/gin"
)

func Render(c *gin.Context, code int, name string, data gin.H) {

	// c.HTML(http.StatusOK, "modules/home/html/home", gin.H{
	// 	"title":    "home page test",
	// 	"APP_NAME": viper.Get("app.name"),
	// })

	// it was second apparoach but we will use the third approach and use the WithGlobalData function from the internal/provider/html/html.go file so we will never change this package and we will change the internal/provider/html/html.go file

	data = view.WithGlobalData(data)
	c.HTML(code, name, data)
}
