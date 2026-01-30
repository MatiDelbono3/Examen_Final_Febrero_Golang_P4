package Services

import (
	"context"
	"errors"
	"time"

	"examen_final_febrero_golang_P4/dtos"
	"examen_final_febrero_golang_P4/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RecetaServiceInterface interface {
	Crear(req dtos.RecetaRequest) (dtos.RecetaResponse, error)
	ListarPaginado(limit int, offset int) ([]models.Receta, error)
	FiltrarRecetasPorCategoria(categoria string) ([]models.Receta, error)
	FiltrarRecetasPorNombre(nombre string) ([]models.Receta, error)
	FiltrarRecetasPorID(id string) ([]models.Receta, error)
}
type RecetaService struct {
	collection     *mongo.Collection
	collectionName string
}

func NewRecetaService(collection *mongo.Collection) *RecetaService {
	return &RecetaService{
		collection: collection,
	}

}
func (service *RecetaService) Crear(req dtos.RecetaRequest) (dtos.RecetaResponse, error) {

	// Validaciones
	if req.Nombre == "" {
		return dtos.RecetaResponse{}, errors.New("el nombre de la receta  es obligatorio")
	}

	if len(req.Ingredientes) == 0 {
		return dtos.RecetaResponse{}, errors.New("Debe haber al menos un ingrediente")
	}

	// Documento a persistir
	doc := bson.M{
		"Nombre":         req.Nombre,
		"Categoria":      req.Categoria,
		"Ingredientes":   req.Ingredientes,
		"Fecha_creacion": time.Now(),
	}

	result, err := service.collection.InsertOne(context.Background(), doc)
	if err != nil {
		return dtos.RecetaResponse{}, err
	}

	id := result.InsertedID.(primitive.ObjectID)

	return dtos.RecetaResponse{
		Id:            id.Hex(),
		Nombre:        req.Nombre,
		Categoria:     req.Categoria,
		Ingredientes:  req.Ingredientes,
		FechaCreacion: time.Now(),
	}, nil
}

var total int

func (service *RecetaService) ListarPaginado(limit int, offset int) ([]models.Receta, error) {

	opts := options.Find()
	opts.SetLimit(int64(limit))
	opts.SetSkip(int64(offset))

	cursor, err := service.collection.Find(context.Background(), bson.M{}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var resultado []models.Receta

	for cursor.Next(context.Background()) {

		var doc struct {
			ID            primitive.ObjectID `bson:"_id"`
			Nombre        string             `bson:"nombre"`
			Categoria     string             `bson:"categoria"`
			Ingredientes  []dtos.Ingrediente `bson:"ingredientes"`
			FechaCreacion time.Time          `bson:"fechaCreacion"`
			IdUsuario     primitive.ObjectID `bson:"idUsuario"`
		}

		if err := cursor.Decode(&doc); err != nil {
			return nil, err
		}

		resultado = append(resultado, models.Receta{
			ID:            doc.ID,
			Nombre:        doc.Nombre,
			Categoria:     doc.Categoria,
			FechaCreacion: doc.FechaCreacion,
		})
	}

	return resultado, nil
}
func (service *RecetaService) FiltrarRecetasPorNombre(nombre string) ([]models.Receta, error) {

	if nombre == "" {
		return nil, errors.New("El estado es obligatorio")
	}

	collection := service.collection

	filter := bson.M{
		"nombre": nombre,
	}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var resultado []models.Receta
	if err := cursor.All(context.TODO(), &resultado); err != nil {
		return nil, err
	}

	return resultado, nil
}
func (service *RecetaService) FiltrarRecetasPorID(Id string) ([]models.Receta, error) {

	if Id == "" {
		return nil, errors.New("El id es obligatorio")
	}

	collection := service.collection

	filter := bson.M{
		"Id": Id,
	}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var resultado []models.Receta
	if err := cursor.All(context.TODO(), &resultado); err != nil {
		return nil, err
	}

	return resultado, nil
}
func (service *RecetaService) FiltrarRecetasPorCategoria(categoria string) ([]models.Receta, error) {

	if categoria == "" {
		return nil, errors.New("La categoria es obligatoria")
	}

	collection := service.collection

	filter := bson.M{
		"Categoria": categoria,
	}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var resultado []models.Receta
	if err := cursor.All(context.TODO(), &resultado); err != nil {
		return nil, err
	}

	return resultado, nil
}
