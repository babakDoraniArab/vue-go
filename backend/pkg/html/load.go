package html

import "github.com/gin-gonic/gin"

func LoadHTML(router *gin.Engine) {
	// pattern is /internal/modules/home/templates/*.tmpl but we are using variable for modules, home , and templates folder and tmp name
	router.LoadHTMLGlob("internal/**/**/**/*tmpl")
}
