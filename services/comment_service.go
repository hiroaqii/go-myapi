package services

import (
	"github.com/hiroaqii/go-myapi/models"
	"github.com/hiroaqii/go-myapi/repositories"
)

func PostCommentService(comment models.Comment) (models.Comment, error) {
	db, err := connectDB()
	if err != nil {
		return models.Comment{}, err
	}
	defer db.Close()

	newComment, err := repositories.InsertComment(db, comment)
	if err != nil {
		return models.Comment{}, err
	}

	return newComment, nil
}