package main

import (
	"context"
	handlers "examen_final_febrero_golang_P4/Handlers"
	Service "examen_final_febrero_golang_P4/Service"
	"examen_final_febrero_golang_P4/middlewares"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	r := gin.Default()
	r.Use(middlewares.AuthMiddleware())
	// Mongo directo
	client, err := mongo.Connect(context.TODO(),
		options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	collection := client.Database("examen").Collection("Recetas")
	service := Service.NewRecetaService(collection)

	handler := handlers.NewRecetaHandler(service)

	r.POST("/recetas", handler.Crear)
	r.GET("/recetas", handler.ListarPaginado)
	r.GET("/receta/:categoria", handler.FiltrarRecetasPorCategoria)
	r.GET("/receta/:Nombre", handler.FiltrarRecetasPorNombre)
	r.GET("/receta/:ID", handler.FiltrarRecetasPorID)
	r.Run(":8080")
}
