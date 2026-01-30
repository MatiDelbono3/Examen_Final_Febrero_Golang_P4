package Models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Receta struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Nombre        string             `bson:"nombre"`
	Categoria     string             `bson:"categoria"`
	IdUsuario     primitive.ObjectID `bson:"usuario_id,omitempty"`
	FechaCreacion time.Time          `bson:"fechaCreacion"`
}
