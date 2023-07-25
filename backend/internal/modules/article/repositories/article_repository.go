package repositories

import (
	ArticleModel "github.com/babakDoraniArab/vue-go/internal/modules/article/models"
	"github.com/babakDoraniArab/vue-go/pkg/database"
	"gorm.io/gorm"
)

type ArticleRepository struct {
	DB *gorm.DB
}

func NewArticleRepository() *ArticleRepository {
	db := database.Connection()
	return &ArticleRepository{
		DB: db,
	}
}

func (articleRepository ArticleRepository) List(limit int) []ArticleModel.Article {
	var articles []ArticleModel.Article
	articleRepository.DB.Limit(limit).Joins("User").Order("rand()").Find(&articles)
	return articles
}
func (articleRepository ArticleRepository) Find(id int) ArticleModel.Article {
	var article ArticleModel.Article
	articleRepository.DB.Joins("User").First(&article, id)
	return article
}
