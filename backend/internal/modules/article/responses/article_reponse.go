package responses

import (
	"fmt"

	ArticleModel "github.com/babakDoraniArab/vue-go/internal/modules/article/models"
	UserResponse "github.com/babakDoraniArab/vue-go/internal/modules/user/responses"
)

type Article struct {
	ID        uint
	Title     string
	Content   string
	Image     string
	CreatedAt string
	User      UserResponse.User
}

type Articles struct {
	Data []Article
}

func ToArticle(article ArticleModel.Article) Article {
	return Article{
		ID:        article.ID,
		Title:     article.Title,
		Content:   article.Content,
		Image:     fmt.Sprintf("/assets/img/demopic/%d.jpg", article.ID),
		CreatedAt: fmt.Sprintf("%d,%02d,%02d", article.CreatedAt.Year(), article.CreatedAt.Month(), article.CreatedAt.Day()),
		User:      UserResponse.ToUser(article.User),
	}
}

func ToArticles(articles []ArticleModel.Article) Articles {
	var response Articles
	for _, article := range articles {

		response.Data = append(response.Data, ToArticle(article))
	}
	return response
}
