package repositories

import (
	"context"
	"log"
	"time"

	"github.com/natanaelrusli/hexagonal-arch/internals/core/ports"
	mongodb "github.com/natanaelrusli/hexagonal-arch/internals/mongodb/docstructures"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	MongoClientTimeout = 5
)

type UserRepository struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

var _ ports.UserRepository = (*UserRepository)(nil)

func NewUserRepository(conn string) (*UserRepository, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), MongoClientTimeout*time.Second)
	defer cancelFunc()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(conn))
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, readpref.Primary())
	log.Println("database connected")
	if err != nil {
		return nil, err
	}

	return &UserRepository{
		client:     client,
		database:   client.Database("goHexagonal"),
		collection: client.Database("goHexagonal").Collection("users"),
	}, nil
}

func (r *UserRepository) Login(email string, password string) error {
	return nil
}

func (r *UserRepository) Register(email string, password string) error {
	doc := new(mongodb.User)

	doc.Email = email
	doc.Password = password

	_, err := r.collection.InsertOne(context.TODO(), doc)

	if err != nil {
		return err
	}

	return nil
}
