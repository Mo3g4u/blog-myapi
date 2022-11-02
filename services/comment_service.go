package services

import (
	"github.com/Mo3g4u/blog-myapi/models"
	"github.com/Mo3g4u/blog-myapi/repositories"
)

func PostCommentService(comment models.Comment) (models.Comment, error) {
	db, err := connectDB()
	if err != nil {
		return models.Comment{}, err
	}

	c, err := repositories.InsertComment(db, comment)
	if err != nil {
		return models.Comment{}, err
	}

	return c, nil
}
