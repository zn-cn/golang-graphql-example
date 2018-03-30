# golang-graphql-example

### 简介：

只是用来玩一玩graphql的一个blog

关于更多Graphql和REST的一些想法，请见[Graphql 学习与REST比较](https://juejin.im/post/5ab7c6c06fb9a028c22abbb5) 

涉及：`golang` ,  `jwt` , `graphql-go` , `mysql`  

##### graphql文档:

```
type User {
  id: ID
  nickname: String!
  email: String!
  pw: string!
  post(id: ID!): Post
  posts: [Post!]!
  follower(id: ID!): User
  followers: [User!]!
  followee(id: ID!): User
  followees: [User!]!
}

type Post {
  id: ID
  user: User!
  title: String!
  body: String!
  praise_num: Int!
  comment_num: Int!
  comment(id: ID!): Comment
  comments: [Comment!]!
}

type Comment {
  id: ID
  user: User!
  post: Post!
  title: String
  body: String!
}

type Query {
  user(user_id: ID!): User
}

type Mutation {
  login(email: String!, pw: String!): Auth
  createUser(nickname: String!, email: String!, pw: String!): User
  removeUser(token: String!, user_id: ID!): Boolean
  follow(token: String!, follower_id: ID!, followee_id: ID!): Boolean
  unfollow(token: String!, follower_id: ID!, followee_id: ID!): Boolean
  createPost(token: String!, user_id: ID!, title: String!, body: String!): Post
  removePost(token: String!, post_id: ID!): Boolean
  updatePost(token: String!, post_id: ID!, title: String, body: String)：Boolean
  createComment(token: String!, user_id: ID!, post_id: ID!, title: String!, body: String!): Comment
  removeComment(token: String!, comment_id: ID!): Boolean
  praisePost(token: String!, post_id: ID!, user_id: ID!): Boolean
  unpraisePost(token: String!, post_id: ID!, user_id: ID!): Boolean
}

```

##### 数据库文档：

```mysql
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
```

##### 测试命令：

这里有两版：

第一版是用graphiql做的，有可视化界面，但是传输的数据只能作为url参数

在网站中输入 http://localhost:1323/graphql， 可查看API文档可视化界面

第二版使用原生的，可用下列目录进行测试：

```shell
$ curl -X POST http://localhost:1323/graphql -d 'mutation{createUser(nickname:"example",email:"example@163.com",pw:"1234567879"){nickname, email}}'
$ curl -X POST http://localhost:1323/graphql -d 'mutation{removeUser(user_id:7)}'
$ curl -X POST http://localhost:1323/graphql -d 'mutation{login(email:"abc@163.com",pw:"1234567879"){token, user{nickname}}}'
$ curl -X POST http://localhost:1323/graphql -d 'mutation{follow(follower_id:8, followee_id:9)}'
$ curl -X POST http://localhost:1323/graphql -d 'mutation{unfollow(follower_id:8, followee_id:9)}'
$ curl -X POST http://localhost:1323/graphql -d 'mutation{createPost(user_id:8, title:"hello2", body:"hello world2"){id,title,body,praise_num,comment_num,create_date}}'
$ curl -X POST http://localhost:1323/graphql -d 'mutation{removePost(post_id:2)}'
$ curl -X POST http://localhost:1323/graphql -d 'mutation{updatePost(post_id:1, title:"hello", body:"hello world")}'
$ curl -X POST http://localhost:1323/graphql -d 'mutation{unpraisePost(post_id:1)}'
$ curl -X POST http://localhost:1323/graphql -d 'mutation{praisePost(post_id:1)}'
$ curl -X POST http://localhost:1323/graphql -d 'mutation{createComment(user_id:8, post_id:1, title:"I like it", body:"I like it"){id,title, body, create_date}}'
$ curl -X POST http://localhost:1323/graphql -d 'mutation{removeComment(comment_id:2)}'
$ curl -X POST http://localhost:1323/graphql -d 'query{user(user_id:13){email}}'
$ curl -X POST http://localhost:1323/graphql -d 'query{user(user_id:13){followers{id, nickname, email}}}'
$ curl -X POST http://localhost:1323/graphql -d 'query{user(user_id:13){followers{id, nickname, email}, followerees{nickname, email}}}'
...
...
```

