package db

import (
	"model"
	"time"
)

// InsertPost db
func InsertPost(post *model.Post) error {
	stmt, prepareErr := DB.Prepare(`
		INSERT INTO post(user_id, title, body, create_date)
		VALUES (?, ?, ?, ?)
	`)
	if prepareErr != nil {
		return prepareErr
	}
	post.CreateDate = time.Now()
	res, execErr := stmt.Exec(post.UserID, post.Title, post.Body, post.CreateDate)
	if execErr != nil {
		return execErr
	}
	id, _ := res.LastInsertId()
	post.ID = int(id)
	return nil
}

// RemovePostByID db
func RemovePostByID(id int) error {
	_, err := DB.Exec(`
		DELETE 
		FROM post 
		WHERE id=?
		`, id)
	return err
}

// GetPostByID db
func GetPostByID(id int) (*model.Post, error) {
	var (
		userID, praiseNum, commentNum int
		title, body                   string
		createDate                    time.Time
	)
	err := DB.QueryRow(`
		SELECT user_id, title, body, praise_num, comment_num, create_date
		FROM post
		WHERE id=?
	`, id).Scan(&userID, &title, &body, &praiseNum, &commentNum, &createDate)
	if err != nil {
		return nil, err
	}
	return &model.Post{
		ID:         id,
		UserID:     userID,
		Title:      title,
		Body:       body,
		PraiseNum:  praiseNum,
		CommentNum: commentNum,
		CreateDate: createDate,
	}, nil
}

// GetPostByIDAndUser db
func GetPostByIDAndUser(id, userID int) (*model.Post, error) {
	var (
		praiseNum, commentNum int
		title, body           string
		createDate            time.Time
	)
	err := DB.QueryRow(`
		SELECT title, body, praise_num, comment_num, create_date
		FROM post
		WHERE id=?
		AND user_id=?
	`, id, userID).Scan(&title, &body, &praiseNum, &commentNum, &createDate)
	if err != nil {
		return nil, err
	}
	return &model.Post{
		ID:         id,
		UserID:     userID,
		Title:      title,
		Body:       body,
		PraiseNum:  praiseNum,
		CommentNum: commentNum,
		CreateDate: createDate,
	}, nil
}

// GetPostsForUser db
func GetPostsForUser(userID int) ([]*model.Post, error) {
	rows, err := DB.Query(`
		SELECT id, title, body, praise_num, comment_num, create_date
		FROM post
		WHERE user_id=?
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var (
		posts                         = []*model.Post{}
		postID, praiseNum, commentNum int
		title, body                   string
		createDate                    time.Time
	)
	for rows.Next() {
		if err = rows.Scan(&postID, &title, &body, &praiseNum, &commentNum, &createDate); err != nil {
			return nil, err
		}
		posts = append(posts, &model.Post{
			ID:         postID,
			UserID:     userID,
			Title:      title,
			Body:       body,
			PraiseNum:  praiseNum,
			CommentNum: commentNum,
			CreateDate: createDate,
		})
	}
	return posts, nil
}

// PraisePost db
func PraisePost(postID int) error {
	stmt, prepareErr := DB.Prepare(`
		UPDATE post 
		SET praise_num=praise_num+1 
		WHERE id=?
		`)
	if prepareErr != nil {
		return prepareErr
	}
	_, execErr := stmt.Exec(postID)
	return execErr
}

// UnPraisePost db
func UnPraisePost(postID int) error {
	stmt, prepareErr := DB.Prepare(`
		UPDATE post 
		SET praise_num=praise_num-1 
		WHERE id=?
		`)
	if prepareErr != nil {
		return prepareErr
	}
	_, execErr := stmt.Exec(postID)
	return execErr
}

// UpdatePost db
func UpdatePost(postID int, title, body string) error {
	stmt, prepareErr := DB.Prepare(`
		UPDATE post 
		SET title=?, body=? 
		WHERE id=?
		`)
	if prepareErr != nil {
		return prepareErr
	}
	_, execErr := stmt.Exec(title, body, postID)
	return execErr
}
