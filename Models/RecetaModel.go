package models

import (
	"examen_final_febrero_golang_P4/dtos"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Receta struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Nombre        string             `bson:"nombre"`
	Categoria     string             `bson:"categoria"`
	Ingredientes  []dtos.Ingrediente `bson:"ingredientes"`
	FechaCreacion time.Time          `bson:"fechaCreacion"`
}
