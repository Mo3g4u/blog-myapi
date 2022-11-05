package services

import (
	"github.com/Mo3g4u/blog-myapi/apperrors"
	"github.com/Mo3g4u/blog-myapi/models"
	"github.com/Mo3g4u/blog-myapi/repositories"
)

func (s *MyAppService) PostCommentService(comment models.Comment) (models.Comment, error) {
	c, err := repositories.InsertComment(s.db, comment)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "fail to record data")
		return models.Comment{}, err
	}

	return c, nil
}
