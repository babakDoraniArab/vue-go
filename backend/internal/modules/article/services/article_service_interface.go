package service

import (
	ArticleResponse "github.com/babakDoraniArab/vue-go/internal/modules/article/responses"
)

type ArticleServiceInterface interface {
	GetFeaturedArticles() ArticleResponse.Articles
	GetStoriesArticles() ArticleResponse.Articles
	Find(id int) (ArticleResponse.Article, error)
}
