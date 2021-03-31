package example

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func maind() {
// 	router := gin.Default()

// 	// ○ routing

// 	// router.GET("/someGet", getting)
// 	// router.POST("/somePost", posting)
// 	// router.PUT("/somePut", putting)
// 	// router.DELETE("/someDelete", deleting)
// 	// router.PATCH("/somePatch", patching)
// 	// router.HEAD("/someHead", head)
// 	// router.OPTIONS("/someOptions", options)

// 	// 1. json 반환
// 	router.GET("/ping", func(c *gin.Context) {
// 		// JSON response
// 		c.JSON(200, gin.H{
// 			"message": "pong",
// 		})
// 	})

// 	// 2. :name, Param
// 	router.GET("/user/:name", func(c *gin.Context) {
// 		name := c.Param("name")
// 		// String response
// 		c.String(http.StatusOK, "Hello %s", name)
// 	})

// 	// 3. * : / 를 붙여야됨
// 	// http://127.0.0.1:8080/user/chlee/swim 시
// 	// name : chlee
// 	// action : /swim
// 	router.GET("/user/:name/*action", func(c *gin.Context) {
// 		name := c.Param("name")
// 		action := c.Param("action")

// 		message := name + " is " + action
// 		c.String(http.StatusOK, message)
// 	})

// 	// 4. FullPath info
// 	router.POST("/user/:name/*action", func(c *gin.Context) {
// 		fmt.Println(c.FullPath() == "/user/:name/*action") // true
// 	})

// 	// 5. Query, DefaultQuery
// 	// POST에서 QueryString으로 보낸 data
// 	// FullPath : /welcome?firstname=lee&lastname=cheolhee
// 	router.GET("/welcome", func(c *gin.Context) {
// 		firstname := c.DefaultQuery("firstname", "Guest")
// 		lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

// 		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
// 	})

// 	// 6. PostForm, DefaultPostForm
// 	// POST form-data 방식으로 보낸 data
// 	router.POST("/form_post", func(c *gin.Context) {
// 		message := c.PostForm("message")
// 		nick := c.DefaultPostForm("nick", "anonymous")

// 		// 7. JSON 응답
// 		c.JSON(200, gin.H{
// 			"status":  "posted",
// 			"message": message,
// 			"nick":    nick,
// 		})
// 	})
// 	router.Run()
// 	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
// 	// http://127.0.0.1:8080/ping
// }
