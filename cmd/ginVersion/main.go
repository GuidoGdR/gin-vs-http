package ginVersion

import (
	"log"

	"github.com/GuidoGdR/go-speed-test/internal/auth"
	"github.com/GuidoGdR/go-speed-test/internal/platform/database"
	"github.com/GuidoGdR/go-speed-test/pkg/token"
	"github.com/gin-gonic/gin"
)

var jwtSecret string

func main() {
	jwtManager := token.NewJWTManager(jwtSecret)

	db, err := database.Postgre("postgres://admin:adminadmin@localhost:5432/my_db?sslmode=disable")
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	store := auth.NewStore(db)
	service := auth.NewService(store, jwtManager)
	handler := auth.NewGinHandler(service)

	router := gin.New()

	{
		router.POST("/login", handler.Login)
		router.POST("/refresh", handler.Refresh)
		router.POST("/register", handler.Register)
	}

	router.Run(":8000")
}

func init() {

	jwtSecret = "insecure-key-123453452414465743543256354232352352323623632"
}
