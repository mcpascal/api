package routers

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // 只允许特定域名
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/ping", pong)
	api(r)
}

func api(r *gin.Engine) {
	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		authWithoutAuth(v1)
		// v1 := api.Group("/v1")
		// authWithoutAuth(v1)
		// v1.Use(middlewares.Auth())
		// {
		// 	auth(v1)
		// 	// file(v1)
		// 	// user(v1)
		// 	journals(v1)
		// }
	}
}

func authWithoutAuth(r *gin.RouterGroup) {
	// c := controllers.NewAuth()
	// r.POST("/login", c.Login)
	// r.POST("/register", c.Register)
}

func auth(r *gin.RouterGroup) {

}

// func WarpH[I any, O any](fn func(*gin.Context, *I) (O, error)) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var in I
// 		if unsafe.Sizeof(in) != 0 {
// 			switch c.Request.Method {
// 			case http.MethodGet:
// 				if err := c.ShouldBindQuery(&in); err != nil {
// 					requests.Fail(c, ErrBadRequest.With(HanddleJSONErr(err).Error()))
// 					return
// 				}
// 			case http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodPatch:
// 				if c.Request.ContentLength > 0 {
// 					if err := c.ShouldBindJSON(&in); err != nil {
// 						Fail(c, ErrBadRequest.With(HanddleJSONErr(err).Error()))
// 						return
// 					}
// 				}
// 			}
// 		}
// 		out, err := fn(c, &in)
// 		if err != nil {
// 			Fail(c, err)
// 			return
// 		}
// 		Success(c, out)
// 	}
// }

// func authWithoutAuth(r *gin.RouterGroup) {
// 	c := controllers.NewAuth()
// 	r.POST("/login", c.Login)
// 	r.POST("/register", c.Register)
// }

// func auth(r *gin.RouterGroup) {
// 	c := controllers.NewAuth()
// 	r.POST("/logout", c.Logout)
// 	r.GET("/refresh", c.Refresh)
// 	r.GET("/profile", c.Profile)
// }

// func file(r *gin.RouterGroup) {
// 	c := controllers.NewFile()
// 	f := r.Group("/files")
// 	{
// 		f.POST("/upload", c.Upload)
// 		// f.GET("/", c.Index)
// 		// f.GET("/:id", c.Show)
// 		// f.POST("/", c.Store)
// 		// f.PUT("/:id", c.Update)
// 		// f.DELETE("/:id", c.Destroy)
// 	}
// }

// func user(r *gin.RouterGroup) {
// 	c := controllers.NewUser()
// 	u := r.Group("/users")
// 	{
// 		u.GET("/", c.Index)
// 		u.GET("/:id", c.Show)
// 		u.POST("/", c.Store)
// 		u.PUT("/:id", c.Update)
// 		u.DELETE("/:id", c.Destroy)
// 		// u.GET("test/:id", WarpH(c.GetUser))
// 	}
// }

// func journals(r *gin.RouterGroup) {
// 	c := controllers.NewJournal()
// 	u := r.Group("/journals")
// 	{
// 		u.GET("/", c.Index)
// 		u.GET("/:id", c.Show)
// 		u.POST("/", c.Store)
// 		u.PUT("/:id", c.Update)
// 		u.DELETE("/:id", c.Destroy)
// 	}
// }

func pong(r *gin.Context) {
	r.JSON(200, gin.H{
		"message": "pong",
	})
}
