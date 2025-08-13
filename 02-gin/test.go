package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 初始化项目 go mod init test
// 安装gin框架 go get -u github.com/gin-gonic/gin
// 更新依赖 go mod tidy
// https://www.cnblogs.com/neural-networker/p/19009923

func main() {
	// 初始化一个Gin引擎实例
	r := gin.Default()

	/* 中间件加在路由前面才有效 */
	// 日志中间件
	// r.Use(func(c *gin.Context) {
	//     fmt.Println("请求地址：" + c.Request.RequestURI)
	//     // 继续处理请求
	//     c.Next()
	// })

	// 使用身份验证中间件
	// r.Use(func(c *gin.Context) {
	//     token := c.GetHeader("Authorization")
	//     // 123:456
	//     if token == "Basic MTIzOjQ1Ng==" {
	//         c.Next()
	//     } else {
	//         c.JSON(http.StatusUnauthorized, gin.H{"message": "No authorization token provided"})
	//         // 终止后续处理
	//         c.Abort()
	//     }
	// })

	// 定义一个简单的GET路由
	// 访问 GET http://localhost:8080/hello
	// 返回 JSON {"message": "Hello, world!"}
	r.GET("/hello", func(c *gin.Context) {
		/* *gin.Context */
		// 在Gin中，每一个请求都会有一个*gin.Context对象
		// Context是用于存储请求信息和处理响应的中心
		// 可以通过它访问请求的数据、绑定参数、设置响应等
		/* 访问请求数据 */
		// c.DefaultQuery("key", "default_value") 获取查询参数
		// c.PostForm("key") 获取表单数据
		// c.ShouldBindJSON(&struct) 将请求的JSON数据绑定到结构体
		/* 设置响应 */
		// c.JSON(http.StatusOK, response) 返回JSON响应
		// c.String(http.StatusOK, "Hello World") 返回字符串响应
		c.JSON(http.StatusOK, gin.H{"message": "Hello, world!"})
	})
	// 路由参数（路径参数）
	// 访问 GET http://localhost:8080/user/123
	// 返回 JSON {"user_id":"123"}
	r.GET("/user/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{"user_id": id})
	})
	// 查询参数（Query参数）
	// 访问 GET http://localhost:8080/search?q=1234
	// 返回 JSON {"query":"1234"}
	// 访问 GET http://localhost:8080/search
	// 返回 JSON {"query":"default query"}
	r.GET("/search", func(c *gin.Context) {
		// 参数名：q 默认值：default query
		query := c.DefaultQuery("q", "default query")
		c.JSON(http.StatusOK, gin.H{"query": query})
	})

	// 路由分组
	v1 := r.Group("/v1")

	// 表单参数
	// 访问 POST http://localhost:8080/v1/login form-data/x-www-urlencoded username:123 password:456
	// 返回 JSON {"password":"456","username":"123"}
	// 访问 POST http://localhost:8080/v1/login
	// 返回 JSON {"password":"defaultPassword","username":""}
	v1.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		// 默认值
		password := c.DefaultPostForm("password", "defaultPassword")
		c.JSON(http.StatusOK, gin.H{"username": username, "password": password})
	})

	// JSON绑定和验证
	// 访问 POST http://localhost:8080/v1/user JSON {"name":"ck","email":"1416978277@qq.com"}
	// 返回 JSON {"message":"User created","user":{"name":"ck","email":"1416978277@qq.com"}}
	// 访问 POST http://localhost:8080/v1/user JSON {"name":"ck","email":"1416978277"}
	// 返回 JSON {"error":"Key: 'User.Email' Error:Field validation for 'Email' failed on the 'email' tag"}
	v1.POST("/user", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "User created", "user": user})
	})

	// 重定向到外部
	r.GET("/baidu", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com/")
	})

	// 重定向到内部
	r.GET("/hello2", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/hello")
	})

	// 重定向到内部2
	r.GET("/hello3", func(c *gin.Context) {
		c.Request.URL.Path = "/hello"
		r.HandleContext(c)
	})

	// 静态文件
	// 文件目录(映射路径,物理路径)
	r.Static("/assets", "../assets")
	// 单独的文件
	r.StaticFile("/favicon.ico", "../assets/favicon.ico")

	// 加载HTML(只能生效一种)
	// HTML文件
	// r.LoadHTMLFiles("./index.html")
	// r.GET("/", func(c *gin.Context) {
	//     c.HTML(http.StatusOK, "index.html", nil)
	// })
	// 模板目录
	// r.LoadHTMLGlob("templates/*")
	// r.GET("/a", func(c *gin.Context) {
	//     c.HTML(http.StatusOK, "a.tmpl", gin.H{
	//         "title": "a",
	//     })
	// })
	r.LoadHTMLGlob("templates2/**/*")
	r.GET("/posts/b", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/b.tmpl", gin.H{
			"title": "b",
		})
	})

	// 自定义错误处理
	r.GET("/error", func(c *gin.Context) {
		err := fmt.Errorf("something went wrong")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	})

	// 启动服务，监听在8080端口
	r.Run(":8080")

}

type User struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}
