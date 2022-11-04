package services

import (
	"github.com/Mo3g4u/blog-myapi/models"
	"github.com/Mo3g4u/blog-myapi/repositories"
)

func (s *MyAppService) GetArticleService(articleID int) (models.Article, error) {
	article, err := repositories.SelectArticleDetail(s.db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	commentList, err := repositories.SelectCommentList(s.db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	article.CommentList = append(article.CommentList, commentList...)

	return article, nil
}

func (s *MyAppService) PostArticleService(article models.Article) (models.Article, error) {
	a, err := repositories.InsertArticle(s.db, article)
	if err != nil {
		return models.Article{}, err
	}

	return a, nil
}

func (s *MyAppService) GetArticleListService(page int) ([]models.Article, error) {
	as, err := repositories.SelectArticleList(s.db, page)
	if err != nil {
		return []models.Article{}, err
	}

	return as, nil
}

func (s *MyAppService) PostNiceService(article models.Article) (models.Article, error) {
	if err := repositories.UpdateNiceNum(s.db, article.ID); err != nil {
		return models.Article{}, err
	}

	return models.Article{
		ID:        article.ID,
		Title:     article.Title,
		Contents:  article.Contents,
		UserName:  article.UserName,
		NiceNum:   article.NiceNum + 1,
		CreatedAt: article.CreatedAt,
	}, nil
}
