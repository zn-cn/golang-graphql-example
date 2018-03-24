package db

import (
	"errors"
	"fmt"
	"model"
	"time"
	"util"
)

// Login db
func Login(user *model.User) (bool, error) {
	if !util.CheckEmail(user.Email) {
		return false, errors.New("your email is unvalid")
	}
	if user.PW == "" {
		return false, errors.New("your password is empty")
	}
	var hashPW, nickname string
	var id int
	var createDate time.Time
	err := DB.QueryRow(`
		SELECT id, hash_pw, nickname, create_date
		FROM user 
		WHERE email=?
		`, user.Email).Scan(&id, &hashPW, &nickname, &createDate)
	if err != nil {
		return false, errors.New("don't have the account")
	}
	if !util.BcryptAuth(user.PW, hashPW) {
		return false, errors.New("password is wrong")
	}

	user.ID = id
	user.NickName = nickname
	user.PW = ""
	user.CreateDate = createDate
	return true, nil
}

// InsertUser db
func InsertUser(user *model.User) error {
	if !util.CheckEmail(user.Email) {
		return errors.New("your email is unvalid")
	}
	if user.PW == "" {
		return errors.New("your password is empty")
	}
	stmt, prepareErr := DB.Prepare(`
		INSERT INTO user(
		nickname, email, hash_pw, create_date) 
		VALUES (?, ?, ?, ?)
		`)
	user.CreateDate = time.Now()
	if prepareErr != nil {
		return prepareErr
	}
	res, execErr := stmt.Exec(user.NickName, user.Email, util.GetBcryptHash(user.PW), user.CreateDate)
	if execErr != nil {
		return execErr
	}
	id, _ := res.LastInsertId()
	user.ID = int(id)
	return nil
}

// CheckUserValid db
func CheckUserValid(id int, email string) (bool, error) {
	var nickname string
	err := DB.QueryRow(`
		SELECT nickname
		FROM user
		WHERE id=?
		AND email=?
		`, id, email).Scan(&nickname)
	if err != nil || nickname == "" {
		return false, err
	}
	return true, nil
}

// GetUserByID db
func GetUserByID(id int) (*model.User, error) {
	var nickname, email string
	var createDate time.Time
	err := DB.QueryRow(`
		SELECT nickname, email, create_date
		FROM user 
		WHERE id=?
		`, id).Scan(&nickname, &email, &createDate)
	if err != nil {
		return nil, err
	}
	return &model.User{
		ID:         id,
		NickName:   nickname,
		Email:      email,
		CreateDate: createDate,
	}, nil
}

// RemoveUserByID db
func RemoveUserByID(id int) error {
	_, err := DB.Exec(`
		DELETE FROM user WHERE id=?
		`, id)
	return err
}

// Follow db
func Follow(followerID, followeeID int) error {
	_, err := DB.Exec(`
		INSERT INTO follow(follower_id, followee_id, create_date) 
		VALUES (?, ?, ?)
		`, followerID, followeeID, time.Now())
	return err
}

// Unfollow db
func Unfollow(followerID, followeeID int) error {
	_, err := DB.Exec(`
		DELETE FROM follow 
		WHERE follower_id=? 
		AND followee_id=?
		`, followerID, followeeID)
	return err
}

// GetFollowerByIDAndUser db
func GetFollowerByIDAndUser(id int, userID int) (*model.User, error) {
	var email, nickname string
	var createDate time.Time
	err := DB.QueryRow(`
		SELECT u.nickname, u.email, u.create_date 
		FROM user AS u, follow AS f 
		WHERE u.id = f.follower_id 
		AND f.follower_id=? 
		AND f.followee_id=? 
		LIMIT 1
		`, id, userID).Scan(&nickname, &email, &createDate)
	if err != nil {
		return nil, err
	}
	return &model.User{
		ID:         id,
		NickName:   nickname,
		Email:      email,
		CreateDate: createDate,
	}, nil
}

// GetFollowersForUser db: get the people who you follow
func GetFollowersForUser(id int) ([]*model.User, error) {
	rows, err := DB.Query(`
		SELECT u.id, u.nickname, u.email, u.create_date
		FROM user AS u, follow AS f
		WHERE u.id=f.follower_id
		AND f.followee_id=?
	`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var (
		users           = []*model.User{}
		userID          int
		nickname, email string
		createDate      time.Time
	)
	fmt.Println("rows")
	for rows.Next() {
		if err = rows.Scan(&userID, &nickname, &email, &createDate); err != nil {
			return nil, err
		}
		users = append(users, &model.User{
			ID:         id,
			NickName:   nickname,
			Email:      email,
			CreateDate: createDate,
		})
	}
	return users, nil
}

// GetFolloweeByIDAndUser db: get the people who follow you
func GetFolloweeByIDAndUser(id int, userID int) (*model.User, error) {
	var (
		nickname   string
		email      string
		createDate time.Time
	)
	err := DB.QueryRow(`
		SELECT u.nickname, u.email, u.create_date
		FROM users AS u, followers AS f
		WHERE u.id = f.followee_id
		AND f.followee_id=?
		AND f.follower_id=?
		LIMIT 1
	`, id, userID).Scan(&nickname, &email, &createDate)
	if err != nil {
		return nil, err
	}
	return &model.User{
		ID:         id,
		NickName:   nickname,
		Email:      email,
		CreateDate: createDate,
	}, nil
}

// GetFolloweesForUser db
func GetFolloweesForUser(id int) ([]*model.User, error) {
	rows, err := DB.Query(`
		SELECT u.id, u.nickname, u.email, u.create_date
		FROM users AS u, followers AS f
		WHERE u.id=f.follower_id
		AND f.follower_id=?
	`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var (
		users           = []*model.User{}
		userID          int
		nickname, email string
		createDate      time.Time
	)
	for rows.Next() {
		if err = rows.Scan(&userID, &nickname, &email, &createDate); err != nil {
			return nil, err
		}
		users = append(users, &model.User{
			ID:         userID,
			NickName:   nickname,
			Email:      email,
			CreateDate: createDate,
		})
	}
	return users, nil
}
