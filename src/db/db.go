// Package db implements a interface for mysql
package db

import (
	"conf"
	"database/sql"
	"fmt"
	"log"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// DB db
var DB *sql.DB

// Init db
func Init() error {
	// create database
	// err := prepare()
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	var err error
	DB, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", conf.Config.DB.UserName, conf.Config.DB.PWD, conf.Config.DB.Host, conf.Config.DB.Port, conf.Config.DB.DBName))
	if err != nil {
		log.Fatal(err)
	}
	err = initTable()
	if err != nil {
		log.Fatal(err)
	}
	return err

}

func prepare() error {
	testDB, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s?parseTime=true", conf.Config.DB.UserName, conf.Config.DB.PWD, "test"))
	defer testDB.Close()
	if err != nil {
		log.Fatalln(err)
	}
	_, err = testDB.Exec("CREATE DATABASE IF NOT EXISTS `blog`")
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func initTable() error {
	tableSQL := [4]string{
		`CREATE TABLE IF NOT EXISTS user(
			id INT UNSIGNED AUTO_INCREMENT,
			nickname VARCHAR(31) NOT NULL,
			email VARCHAR(100) NOT NULL UNIQUE,
			hash_pw VARCHAR(255) NOT NULL,
			create_date DATETIME NOT NULL,
			PRIMARY KEY (id)
		)ENGINE=InnoDB DEFAULT CHARSET=utf8;`,
		`CREATE TABLE IF NOT EXISTS post(
			id INT UNSIGNED AUTO_INCREMENT,
			user_id INT UNSIGNED NOT NULL,
			title VARCHAR(31) NOT NULL,
			body VARCHAR(255) NOT NULL,
			praise_num INT NOT NULL DEFAULT 0,
			comment_num INT NOT NULL DEFAULT 0,
			create_date DATETIME NOT NULL,
			PRIMARY KEY (id),
			FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE,
			INDEX user (user_id)
		)ENGINE=InnoDB DEFAULT CHARSET=utf8;`,
		`CREATE TABLE IF NOT EXISTS comment(
			id INT UNSIGNED AUTO_INCREMENT,
			post_id INT UNSIGNED NOT NULL,
			user_id INT UNSIGNED NOT NULL,
			title VARCHAR(31) NOT NULL,
			body VARCHAR(255) NOT NULL,
			create_date DATETIME NOT NULL,
			PRIMARY KEY (id),
			FOREIGN KEY (post_id) REFERENCES post(id) ON DELETE CASCADE,
			INDEX post (post_id)
		)ENGINE=InnoDB DEFAULT CHARSET=utf8;`,
		`CREATE TABLE IF NOT EXISTS follow(
			id INT UNSIGNED AUTO_INCREMENT,
			follower_id INT UNSIGNED NOT NULL,
			followee_id INT UNSIGNED NOT NULL,
			create_date DATETIME NOT NULL,
			PRIMARY KEY (id),
			FOREIGN KEY (follower_id) REFERENCES user(id) ON DELETE CASCADE,
			FOREIGN KEY (followee_id) REFERENCES user(id) ON DELETE CASCADE,
			INDEX follower (follower_id),
			INDEX followee (followee_id)
		)ENGINE=InnoDB DEFAULT CHARSET=utf8;`,
	}

	// sql Events
	tx, err := DB.Begin()
	for _, v := range tableSQL {
		_, err = tx.Exec(v)
		// _, err = DB.Exec(tableSQL[i])
		if err != nil {
			log.Fatal(err)
		}
	}
	tx.Commit()
	return nil
}
