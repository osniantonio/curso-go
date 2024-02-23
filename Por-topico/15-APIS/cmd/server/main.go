package main

import (
	"log"
	"net/http"

	"github.com/osniantonio/goexpert/15-APIS/configs"
	"github.com/osniantonio/goexpert/15-APIS/internal/entity"
	"github.com/osniantonio/goexpert/15-APIS/internal/infra/database"
	"github.com/osniantonio/goexpert/15-APIS/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"

	_ "github.com/osniantonio/goexpert/15-APIS/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Go Expert API Example
// @version 1.0
// @description Product API with authentication
// @termsOfService http://swagger.io/terms/

// @contact.name Osni Antônio Küchler
// @contact.url https://www.linkedin.com/in/osni-ant%C3%B4nio-k%C3%BCchler-5487b13b
// @contact.email osni.antonio@gmail.com

// @host localhost:8000
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})

	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB)

	r := chi.NewRouter()

	// middleware

	// r.Use(LogRequest)
	r.Use(middleware.Logger)

	// mesmo se ocorrer uma falha ele não deixa o servidor web cair
	r.Use(middleware.Recoverer)

	r.Use(middleware.WithValue("jwt", configs.TokenAuth))
	r.Use(middleware.WithValue("jwtExpiresIn", configs.JwtExpiresIn))

	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(configs.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/", productHandler.CreateProduct)
		r.Get("/{id}", productHandler.GetProduct)
		r.Get("/", productHandler.GetProducts)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	r.Post("/users", userHandler.CreateUser)
	r.Post("/users/generate_token", userHandler.GetJwt)

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))

	http.ListenAndServe(":8000", r)
}

// middleware criado para aprendizado
func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// r.Context().Value("user")
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
