package routing

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"golang-endpoint-boilerplate/db" 
	"golang-endpoint-boilerplate/users" // --> Change this directory path to the name of your folder / project

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func InitialRouting() {
	fmt.Println("Starting routes with gin")
	r := gin.New()

	r.Use(cors.New(cors.Options{
		AllowedOrigins:     []string{"http://localhost:4200", "https://yourdomain.com"},
		AllowedMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowedHeaders:     []string{"Authorization", "Content-type"},
		AllowCredentials:   false,
		OptionsPassthrough: false,
		Debug:              false,
	}))

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	authMiddleware := &jwt.GinJWTMiddleware{
		Realm:         "yourdomain.com",
		Key:           []byte("something super sup3r s4per secret"),
		Timeout:       time.Hour * 24,
		MaxRefresh:    time.Hour * 24,
		Authenticator: authenticate,
		PayloadFunc:   payload,
	}

	authAdminMiddleware := &jwt.GinJWTMiddleware{
		Realm:         "admin yourdomain.com",
		Key:           []byte("something super sup3r s4per secret for 4dm1n"),
		Timeout:       time.Hour * 24,
		MaxRefresh:    time.Hour * 24,
		Authenticator: authenticateAdmin,
		PayloadFunc:   payload,
	}

	r.GET("/", home)
	r.GET("/health", health)
	r.GET("/users", usersAll)
	r.POST("/user", createPerson)
	r.GET("/user/:id", readPerson)

	r.POST("/auth/login", authMiddleware.LoginHandler)
	r.POST("/admin/login", authAdminMiddleware.LoginHandler)

	auth := r.Group("/auth")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/protected", protected)
		auth.GET("/refreshToken", authMiddleware.RefreshHandler)
	}

	admin := r.Group("/admin")
	admin.Use(authAdminMiddleware.MiddlewareFunc())
	{
		admin.GET("/protected", protectedAdmin)
	}

	fmt.Println("Listening the application on port 5000 with gin")
	log.Fatal(r.Run(":5000"))
}

func protected(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	c.String(http.StatusOK, "id: %s\nrole: %s", claims["id"], claims["role"])
}

func protectedAdmin(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	c.String(http.StatusOK, "ADMIN id: %s\nrole: %s", claims["id"], claims["role"])
}

func authenticate(email string, password string, c *gin.Context) (string, bool) {
	if users.Authenticate(email, password) {
		return email, true
	}
	return "", false
}

func authenticateAdmin(email string, password string, c *gin.Context) (string, bool) {
	if users.Authenticate(email, password) && users.IsAdmin(email) {
		return email, true
	}
	return "", false
}

func payload(email string) map[string]interface{} {
	user, err := users.ReadUserByEmail(email)
	if err != nil {
		return map[string]interface{}{
			"error": err,
		}
	}
	if user.Admin {
		return map[string]interface{}{
			"id":   user.ID,
			"role": "admin",
		}
	} else {
		return map[string]interface{}{
			"id":   user.ID,
			"role": "user",
		}
	}
}

func health(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}

func home(c *gin.Context) {
	c.JSON(200, gin.H{
		"title": "My website",
	})
}

func usersAll(c *gin.Context) {
	c.JSON(200, users.ReadAll())
}

func createPerson(c *gin.Context) {
	var user db.User
	c.BindJSON(&user)
	users.CreateUser(&user)
	c.JSON(200, user)
}

func readPerson(c *gin.Context) {
	id := c.Params.ByName("id")
	user, err := users.ReadUser(id)
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, user)
	}
}
