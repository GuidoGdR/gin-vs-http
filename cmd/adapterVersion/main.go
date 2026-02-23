package adapterVersion

import (
	"log"
	"net/http"

	"github.com/GuidoGdR/go-speed-test/internal/auth"
	"github.com/GuidoGdR/go-speed-test/internal/platform/adapter"
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
	handler := auth.NewAdapterHandler(service, validate)
	{
		http.HandleFunc("/login", adapter.HTTP(handler.Login))
		http.HandleFunc("/refresh", adapter.HTTP(handler.Refresh))
		http.HandleFunc("/register", adapter.HTTP(handler.Register))
	}

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func init() {

	jwtSecret = "insecure-key-123453452414465743543256354232352352323623632"
	/*
		jwtKey = []byte(os.Getenv("JWT_SECRET"))
		if len(jwtKey) == 0 {
			panic("JWT_SECRET environment variable not set")
		}
	*/
}
