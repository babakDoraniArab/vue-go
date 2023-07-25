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

func (controller *Controller) Index(c *gin.Context) {
	html.Render(c, http.StatusOK, "modules/home/html/home", gin.H{
		"title":    "home page test",
		"featured": controller.articleService.GetFeaturedArticles(),
		"stories":  controller.articleService.GetStoriesArticles(),

		// "APP_NAME": viper.Get("app.name"), I removed the APP_NAME from here and I moved it to the pkg/html/render.go file so it will be in the all files
	})

	// test the Repository
	// c.JSON(http.StatusOK, gin.H{
	// 	"featured": controller.articleService.GetFeaturedArticles(),
	// 	"stories":  controller.articleService.GetStoriesArticles(),
	// })

}
