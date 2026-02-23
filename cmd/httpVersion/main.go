package httpVersion

import (
	"log"
	"net/http"

	"github.com/GuidoGdR/go-speed-test/internal/auth"
	"github.com/GuidoGdR/go-speed-test/internal/platform/database"
	"github.com/GuidoGdR/go-speed-test/pkg/token"
	"github.com/go-playground/validator/v10"
)

var jwtSecret string

func main() {
	jwtManager := token.NewJWTManager(jwtSecret)
	validate := validator.New()

	db, err := database.Postgre("postgres://admin:adminadmin@localhost:5432/my_db?sslmode=disable")
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	store := auth.NewStore(db)
	service := auth.NewService(store, jwtManager)
	handler := auth.NewHTTPHandler(service, validate)
	{

		http.HandleFunc("/login", handler.Login)
		http.HandleFunc("/refresh", handler.Refresh)
		http.HandleFunc("/register", handler.Register)
	}

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func init() {

	jwtSecret = "insecure-key-123453452414465743543256354232352352323623632"
}
