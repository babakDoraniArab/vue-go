package controllers

import (
	"net/http"

	ArticleService "github.com/babakDoraniArab/vue-go/internal/modules/article/services"
	"github.com/babakDoraniArab/vue-go/pkg/html"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	articleService ArticleService.ArticleServiceInterface
}

func New() *Controller {
	return &Controller{articleService: ArticleService.NewArticleService()}
}

func (controller *Controller) Dashboard(c *gin.Context) {
	html.Render(c, http.StatusOK, "modules/dashboard/html/dashboard", gin.H{
		"title": "dashboard page test",
	})

}
func (controller *Controller) Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "test2"})

}
