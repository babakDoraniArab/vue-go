package service

import (
	"errors"

	ArticleRepository "github.com/babakDoraniArab/vue-go/internal/modules/article/repositories"
	ArticleResponse "github.com/babakDoraniArab/vue-go/internal/modules/article/responses"
)

type ArticleService struct {
	articleRepository ArticleRepository.ArticleRepositoryInterface
}

func NewArticleService() *ArticleService {
	return &ArticleService{
		articleRepository: ArticleRepository.NewArticleRepository(),
	}
}

func (articleService ArticleService) GetFeaturedArticles() ArticleResponse.Articles {
	articles := articleService.articleRepository.List(4)
	return ArticleResponse.ToArticles(articles)
}

func (articleService ArticleService) GetStoriesArticles() ArticleResponse.Articles {
	articles := articleService.articleRepository.List(6)
	return ArticleResponse.ToArticles(articles)
}

func (articleService ArticleService) Find(id int) (ArticleResponse.Article, error) {
	var response ArticleResponse.Article
	article := articleService.articleRepository.Find(id)
	if article.ID == 0 {
		return response, errors.New("Article not found")
	}
	return ArticleResponse.ToArticle(article), nil
}
