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
	ListarPaginado(limit int, offset int) (dtos.ListarPaginadoResponse, error)
	FiltrarRecetasPorCategoria(categoria string) ([]models.Receta, error)
	FiltrarRecetasPorNombre(nombre string) ([]models.Receta, error)
	FiltrarRecetasPorID(id string) (dtos.RecetaResponse, error)
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
	receta := models.Receta{

		Nombre:        req.Nombre,
		Categoria:     req.Categoria,
		Ingredientes:  req.Ingredientes,
		FechaCreacion: time.Now(),
	}

	result, err := service.collection.InsertOne(context.Background(), receta)
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

func (service *RecetaService) ListarPaginado(limit int, offset int) (dtos.ListarPaginadoResponse, error) {

	ctx := context.Background()

	total, err := service.collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return dtos.ListarPaginadoResponse{}, err
	}
	opts := options.Find()
	opts.SetLimit(int64(limit))
	opts.SetSkip(int64(offset))

	cursor, err := service.collection.Find(context.Background(), bson.M{}, opts)
	if err != nil {
		return dtos.ListarPaginadoResponse{}, err
	}
	defer cursor.Close(context.Background())

	var items []dtos.RecetaResponse

	for cursor.Next(context.Background()) {

		var r models.Receta
		if err := cursor.Decode(&r); err != nil {
			return dtos.ListarPaginadoResponse{}, err
		}

		items = append(items, dtos.RecetaResponse{
			Id:            r.ID.Hex(),
			Nombre:        r.Nombre,
			Categoria:     r.Categoria,
			Ingredientes:  r.Ingredientes,
			FechaCreacion: r.FechaCreacion,
		})
	}
	return dtos.ListarPaginadoResponse{
		Total: int(total),
		Items: items,
	}, nil
}

func (service *RecetaService) FiltrarRecetasPorNombre(nombre string) ([]models.Receta, error) {

	if nombre == "" {
		return nil, errors.New("El estado es obligatorio")
	}

	// Búsqueda parcial e insensible a mayúsculas
	filter := bson.M{
		"nombre": bson.M{
			"$regex":   nombre,
			"$options": "i",
		},
	}

	cursor, err := service.collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var recetas []models.Receta
	if err := cursor.All(context.Background(), &recetas); err != nil {
		return nil, err
	}
	return recetas, nil
}

func (s *RecetaService) FiltrarRecetasPorID(id string) (dtos.RecetaResponse, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return dtos.RecetaResponse{}, errors.New("id inválido")
	}

	var r models.Receta
	err = s.collection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&r)
	if err != nil {
		return dtos.RecetaResponse{}, err
	}

	return dtos.RecetaResponse{
		Id:            r.ID.Hex(),
		Nombre:        r.Nombre,
		Categoria:     r.Categoria,
		Ingredientes:  r.Ingredientes,
		FechaCreacion: r.FechaCreacion,
	}, nil
}
func (service *RecetaService) FiltrarRecetasPorCategoria(categoria string) ([]models.Receta, error) {

	if categoria == "" {
		return nil, errors.New("La categoria es obligatoria")
	}

	collection := service.collection

	filter := bson.M{
		"categoria": categoria,
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
