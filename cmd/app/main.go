package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	"github.com/inaciogu/go-product-service/internal/infra/messaging/akafka"
	repository "github.com/inaciogu/go-product-service/internal/infra/repositories"
	"github.com/inaciogu/go-product-service/internal/infra/web/handlers"
	usecase "github.com/inaciogu/go-product-service/internal/useCases"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(host.docker.internal:3306)/products")

	if err != nil {
		panic(err)
	}
	println("Connected to database")
	defer db.Close()

	repository := repository.NewRepositoryMysql(db)

	createProductUseCase := usecase.NewCreateProductUseCase(repository)
	listAllProductsUseCase := usecase.NewListAllProductsUseCase(repository)

	productHandlers := handlers.NewProductHandlers(createProductUseCase, listAllProductsUseCase)

	router := chi.NewRouter()

	router.Post("/", productHandlers.CreateProduct)
	router.Get("/", productHandlers.ListAllProducts)

	go http.ListenAndServe(":8080", router)

	msgChan := make(chan *kafka.Message)
	go akafka.Consume([]string{"products"}, "host.docker.internal:9094", msgChan)

	for msg := range msgChan {
		dto := usecase.CreateProductInputDto{}
		err := json.Unmarshal(msg.Value, &dto)

		if err != nil {
			println(err)
		}
		_, err = createProductUseCase.Execute(dto)

		if err != nil {
			println(err)
		}
	}
}
