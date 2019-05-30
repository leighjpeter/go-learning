package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"net/http"
	"time"
)

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d/%02d/%02d", year, month, day)
}

var html = template.Must(template.New("https").Parse(`
<html>
<head>
  <title>Https Test</title>
  <script src="/assets/app.js"></script>
</head>
<body>
  <h1 style="color:red;">Welcome, Ginner!</h1>
</body>
</html>
`))

type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required" `
}

type LoginJson struct {
	User     string `json:"user" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Person struct {
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func main() {
	router := gin.Default()

	router.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	router.LoadHTMLGlob("templates/**/*")
	router.GET("/posts/:action", func(c *gin.Context) {
		action := c.Param("action")
		if action == "index" {
			c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
				"title": "Posts",
			})
		} else {
			c.String(http.StatusOK, "Action= %s", action)
		}

	})
	router.GET("/users/:name/:action", func(c *gin.Context) {
		// firstname:=c.DefaultQuery("firstname","Guest")
		// lastname := c.Query("lastname") 是c.Request.URL.Query().GET("lastname") 的简写
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	router.GET("/raw", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.tmpl", map[string]interface{}{
			"now": time.Date(2017, 07, 01, 0, 0, 0, 0, time.UTC),
		})
	})

	//eg. curl -v --form user=user --form password=password http://localhost:9090/login
	router.POST("/loginFrom", func(c *gin.Context) {
		var form LoginForm
		// c.ShouldBindWith(&form, binding.Form)
		// message := c.PostForm("message")
		// nick := c.DefaultPostForm("nick", "anonymous") // 此方法可以设置默认值
		if err := c.ShouldBind(&form); err == nil {
			if form.User == "user" && form.Password == "password" {
				c.JSON(200, gin.H{"status": "you are logged in"})
			} else {
				c.JSON(401, gin.H{"status": "unauthorized"})
			}
		}
	})

	router.POST("loginJSON", func(c *gin.Context) {
		var login_json LoginJson
		if err := c.ShouldBindJSON(&login_json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if login_json.User != "user" || login_json.Password != "pwd" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	router.GET("startPage", startPage)

	//使用 SecureJSON 防止 json 劫持。如果给定的结构是数组值，则默认预置 "while(1)," 到响应体。
	// 你也可以使用自己的 SecureJSON 前缀
	// router.SecureJsonPrefix(")]}',\n")
	router.GET("/someJSON", func(c *gin.Context) {
		names := []string{"lena", "austin", "foo"}

		// 将输出：while(1);["lena","austin","foo"]
		c.SecureJSON(http.StatusOK, names)
	})

	router.GET("/JSONP?callback=x", func(c *gin.Context) {
		data := map[string]interface{}{
			"foo": "bar",
		}
		// callback 是 x
		// 将输出：x({\"foo\":\"bar\"})
		c.JSONP(http.StatusOK, data)
	})
	// 文件上传
	router.POST("/upload", func(c *gin.Context) {
		// 单文件
		file, _ := c.FormFile("file")
		// log.Println(file.Filename)

		// 上传文件至指定目录
		// c.SaveUploadedFile(file, dst)
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))

		// 多文件 Multipart form
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]

		for _, file := range files {
			fmt.Println(file.Filename)
			// 上传文件至指定目录
			// c.SaveUploadedFile(file, dst)
		}
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	})

	//Redirect
	router.GET("/test", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.google.com/")
	})
	//路由重定向
	router.GET("/test", func(c *gin.Context) {
		c.Request.URL.Path = "/test2"
		router.HandleContext(c)
	})
	router.GET("/test2", func(c *gin.Context) {
		c.JSON(200, gin.H{"hello": "world"})
	})

	//静态文件服务
	router.Static("/assets", "./assets")
	router.StaticFS("/more_static", http.Dir("my_file_system"))
	router.StaticFile("/favicon.ico", "./resources/favicon.ico")

	router.Run(":9090")
}

func startPage(c *gin.Context) {
	var person Person
	// 如果是 `GET` 请求，只使用 `Form` 绑定引擎（`query`）。
	// 如果是 `POST` 请求，首先检查 `content-type` 是否为 `JSON` 或 `XML`，然后再使用 `Form`（`form-data`）。
	// 查看更多：https://github.com/gin-gonic/gin/blob/master/binding/binding.go#L48
	if c.ShouldBind(&person) == nil {
		log.Println(person.Name)
		log.Println(person.Address)
		log.Println(person.Birthday)
	}

	c.String(200, "Success")
}
