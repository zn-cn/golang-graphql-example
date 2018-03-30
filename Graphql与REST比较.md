学习链接：

+ [官方英文文档](https://graphql.org/)
+ [官方中文文档](http://graphql.cn/learn/queries/) 

## 一、graphql 初体验

##### 官方定义：

**A query language for your API** (API 查询语言)

GraphQL is a query language for APIs and a runtime for **fulfilling those queries with your existing data**. GraphQL provides a complete and understandable description of the data in your API, gives clients the power to **ask for exactly what they need** and nothing more, makes it easier to evolve APIs over time, and enables powerful developer tools.

举个例子：

graphql 如下：

如果请求如下（真实开发可能有点差别）：

```
query{ 
	comment(id: "1") { 
	      title 
	      user { 
	           nickname
               email
	      } 
	}
}
```

后台一部分的类型定义如下：

```
type Comment {
  id: ID
  user: User!
  post: Post!
  title: String
  body: String!
}
```

那么将可能会返回：

```
{
  "title": "hello world",
  "user": {
    "nickname": "molscar",
    "email": "example@gmail.com",
  }
}
```

而如果使用`REST` 来发请求的话，会是这样：

```
GET /comment/1
```

## 二、graphql 和 REST的比较

RESTful 大概是这样：

![img](https://user-gold-cdn.xitu.io/2017/6/19/78ad4112dcd66f01524eca4c02f2ff9f?imageView2/0/w/1280/h/960/format/webp/ignore-error/1)



那么 graphql 大概就是这样：

![img](https://user-gold-cdn.xitu.io/2017/6/19/217cfad3d404089c1446f18778eab810?imageView2/0/w/1280/h/960/format/webp/ignore-error/1)

再给一个例子感受一下：

需求：有关电影和其出演人员的信息

+ REST

  那么可能要先GET到有关电影的信息，再根据有关电影中的出演人员的信息来GET一次演员的信息

  当然如果后台愿意额外构建一个类似`/moviesAndActors` 的接口的话，也是可以的，但是经常改需求的话，后台接口会越来越多，会很臃肿

+ GraphQL

  graphql直接使用如下查询语句即可

  ```
  query MoviesAndActors {
    movies {
      title
      image
      actors {
        image
        name
      }
    }
  }
  ```

  ​

#### RESTful的一些不足：

+ 单个RESTful接口返回数据越来越臃肿（无法控制后台返回的数据）

  比如获取用户信息`/users/:id`，最初可能只有id、昵称，但随着需求的变化，用户所包含的字段可能会越来越多，年龄、性别、头像、经验、等级，等等等等。

  而具体到某个前端页面，可能只需要其中一小部分数据，这样就会增加网络传输量，前端获取了大量不必要的数据。

+ 有时候可能需要多个请求才能获取到足够的数据

  比如获取一个帖子，刚开始开发的时候可能只需要将帖子的内容返回，但是后期可能还要返回发帖人的各种信息，如头像、昵称等，或者还需要获取帖子的评论，或者需求又改了。。。。

  那么这时候你要么在后台再加一个独立的接口，或者让前端使用多个请求来获取足够的数据

#### GraphQL的一些优点：

+ 可以通过请求控制返回的数据

  如果：请求如下，那么只会返回评论的title，以及发评论用户的nickname和email，而不会返回用户的id等其他信息

  ```
  query{ 
  	comment(id: "1") { 
  	      title 
  	      user { 
  	           nickname
                 email
  	      } 
  	}
  }
  ```

+ 请求数量大减

  可以通过一个资源入口访问到关联的其他资源，只要事先在schema中定义好资源之间的关系即可，传输不一样的数据，而REST则提供了多个URL端点来获取相关的资源。

+ 参数类型检验

  graphql提供自动的类型检验以及转换机制

+ 文档清晰

  graphql可以根据代码直接生成可视化的文档界面，界面如下：

  ![img](https://user-gold-cdn.xitu.io/2018/3/26/16260164f49bbd23?imageView2/2/w/480/h/480/q/85/interlace/1)

+ 扩展性好

  可以轻松应对需求变更



#### 路由处理器（Route Handlers）和 解析器（Resolvers）

Route Handlers：

1. 服务器收到HTTP请求，提取出HTTP方法名与URL路径
2. API框架找到提前注册好的、请求路径与请求方法都匹配的代码
3. 该段代码被执行，并得到相应结果
4. API框架对数据进行序列化，添加上适当的状态码与响应头后，返回给客户端

Resolvers：

1. 服务器收到HTTP请求，提取其中的GraphQL查询
2. 遍历查询语句，调用里面每个字段所对应的Resolver。
3. Resolver函数被执行并返回相应结果
4. GraphQL框架把结果根据查询语句的要求进行组装（匹配）



## 三、个人感受

GraphQL使用解析器来构建API，让返回的数据不会那么臃肿，而且也非常酷啊，可以自动解析需要返回的数据，以及校验类型，graphql还有很多东西，这里就不涉及了，我也不是很了解，不过graphql和REST还是有很多相通的地方。

虽然GraphQL确实有很多优点，但是个人感觉不是很好构建一个完整的API，尤其是特别复杂的项目的时候，而且现在graphql没啥轮子，最近用golang玩一玩graphql的时候，想加个装饰器都不好加，graphql没有暴露很多接口，封装很紧，而且因为graphql自己有一套语法，如果想自己解析的话，得自己写一套解析器，找了半天都没找打怎么获取request的header部分的内容等。。。想给graphql加一层装饰器，然后传一个参数给下一层都不好传。。。肯定是自己太菜了

个人实战示例：https://github.com/tofar/golang-graphql-example
