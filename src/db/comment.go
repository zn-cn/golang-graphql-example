package db

import (
	"model"
	"time"
)

// InsertComment db
func InsertComment(comment *model.Comment) error {
	var commentID int
	comment.CreateDate = time.Now()
	DB.QueryRow(`
		INSERT INTO comment(user_id, post_id, title, body, create_date)
		VALUES (?, ?, ?, ?, ?)
	`, comment.UserID, comment.PostID, comment.Title, comment.Body, comment.CreateDate).Scan(&commentID)

	comment.ID = commentID
	return nil
}

// RemoveCommentByID db
func RemoveCommentByID(id int) error {
	_, err := DB.Exec(`
		DELETE 
		FROM comment WHERE id=?
		`, id)
	return err
}

// GetCommentByIDAndPost db
func GetCommentByIDAndPost(id int, postID int) (*model.Comment, error) {
	var (
		userID      int
		title, body string
		createDate  time.Time
	)
	err := DB.QueryRow(`
		SELECT user_id, title, body, create_date
		FROM posts
		WHERE id=?
		AND post_id=?
	`, id, postID).Scan(&userID, &title, &body, &createDate)
	if err != nil {
		return nil, err
	}
	return &model.Comment{
		ID:         id,
		UserID:     userID,
		PostID:     postID,
		Title:      title,
		Body:       body,
		CreateDate: createDate,
	}, nil
}

// GetCommentsForPost db
func GetCommentsForPost(id int) ([]*model.Comment, error) {
	rows, err := DB.Query(`
		SELECT id, user_id, title, body, create_date
		FROM comments
		WHERE post_id=?
	`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var (
		comments          = []*model.Comment{}
		commentID, userID int
		title, body       string
		createDate        time.Time
	)
	for rows.Next() {
		if err = rows.Scan(&commentID, &userID, &title, &body, &createDate); err != nil {
			return nil, err
		}
		comments = append(comments, &model.Comment{
			ID:         commentID,
			UserID:     userID,
			PostID:     id,
			Title:      title,
			Body:       body,
			CreateDate: createDate,
		})
	}
	return comments, nil
}
