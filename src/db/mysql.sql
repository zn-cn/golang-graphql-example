-- create database
CREATE DATABASE IF NOT EXISTS `blog_test`;

-- create table
CREATE TABLE IF NOT EXISTS `user`(
    `id` INT UNSIGNED AUTO_INCREMENT,
    `nickname` VARCHAR(31) NOT NULL,
    `email` VARCHAR(100) NOT NULL UNIQUE,
    `hash_pw` VARCHAR(255) NOT NULL,
    `create_date` DATE NOT NULL,
    PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `post`(
    `id` INT UNSIGNED AUTO_INCREMENT,
    `user_id` INT UNSIGNED NOT NULL,
    `title` VARCHAR(31) NOT NULL,
    `body` VARCHAR(255) NOT NULL,
    `praise_num` INT NOT NULL DEFAULT 0,
    `comment_num` INT NOT NULL DEFAULT 0,
    `create_date` DATE NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`user_id`) REFERENCES user(`id`) ON DELETE CASCADE,
    INDEX user (`user_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `comment`(
    `id` INT UNSIGNED AUTO_INCREMENT,
    `post_id` INT UNSIGNED NOT NULL,
    `user_id` INT UNSIGNED NOT NULL,
    `title` VARCHAR(31) NOT NULL,
    `body` VARCHAR(255) NOT NULL,
    `create_date` DATE NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`post_id`) REFERENCES post(`id`) ON DELETE CASCADE,
    INDEX post (`post_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `follow`(
    `id` INT UNSIGNED AUTO_INCREMENT,
    `follower_id` INT UNSIGNED NOT NULL,
    `followee_id` INT UNSIGNED NOT NULL,
    `create_date` DATE NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`follower_id`) REFERENCES user(`id`) ON DELETE CASCADE,
    FOREIGN KEY (`followee_id`) REFERENCES user(`id`) ON DELETE CASCADE,
    INDEX follower (`follower_id`),
    INDEX followee (`followee_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- insert query
INSERT INTO `user` (nickname, email, hash_pw, create_date) VALUES (?, ?, ?, ?)

INSERT INTO `post` (user_id, title, body, praise_num, comment_num, create_date) VALUES (?, ?, ?, ?, ?, ?);

INSERT INTO `comment` (post_id, user_id, title, body, create_date ) VALUES (?, ?, ?, ?, ?);

INSERT INTO `follow` (follower_id, followee_id, create_date ) VALUES (?, ?, ?);

-- select
SELECT column_name FROM `user` [WHERE Clause] [LIMIT N][ OFFSET M];
SELECT column_name FROM `post` [WHERE Clause] [LIMIT N][ OFFSET M];
SELECT column_name FROM `comment` [WHERE Clause] [LIMIT N][ OFFSET M];
SELECT column_name FROM `folllow` [WHERE Clause] [LIMIT N][ OFFSET M];

-- update 
UPDATE `user` SET field1=new-value1, field2=new-value2 [WHERE Clause]
UPDATE `post` SET field1=new-value1, field2=new-value2 [WHERE Clause]
UPDATE `comment` SET field1=new-value1, field2=new-value2 [WHERE Clause]
UPDATE `follow` SET field1=new-value1, field2=new-value2 [WHERE Clause]

-- delete 
DELETE FROM `user` [WHERE Clause]
DELETE FROM `post` [WHERE Clause]
DELETE FROM `comment` [WHERE Clause]
DELETE FROM `folllow` [WHERE Clause]