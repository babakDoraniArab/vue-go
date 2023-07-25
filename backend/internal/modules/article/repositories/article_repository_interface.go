package repositories

import ArticleModel "github.com/babakDoraniArab/vue-go/internal/modules/article/models"

type ArticleRepositoryInterface interface {
	// Get all Articles
	List(limit int) []ArticleModel.Article
	Find(id int) ArticleModel.Article
}
