package controllers

import (
	"net/http"
	"strconv"

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

func (controller *Controller) Show(c *gin.Context) {
	// html.Render(c, http.StatusOK, "modules/home/html/home", gin.H{
	// 	"title":    "home page test",

	// })

	//get the article
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// c.JSON(http.StatusNotFound, gin.H{"message": "article not found"})
		html.Render(c, http.StatusNotFound, "templates/layouts/errors/500", gin.H{
			"title":   "server error",
			"message": "there is a problem with the server id should be integer",
		})
		return
	}
	// find the article from the database

	article, err := controller.articleService.Find(id)

	//if the article is not found, return 404
	if err != nil {
		// c.JSON(http.StatusNotFound, gin.H{"message": "article not found"})
		html.Render(c, http.StatusNotFound, "templates/layouts/errors/404", gin.H{
			"title":   "Bad Request",
			"message": "article not found",
		})
		return
	}

	//if the article is found, render the article page
	html.Render(c, http.StatusOK, "modules/article/html/show", gin.H{
		"title":   "Show Article",
		"article": article,
	})
}
