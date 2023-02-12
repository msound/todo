package db

import (
	"context"
	"errors"

	"github.com/msound/todo/pkg/todo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
)

type Client struct {
	mongoClient *mongo.Client
	db          *mongo.Database
}

const COLLECTION_LIST string = "list"

func NewClient(connectionString string) (*Client, error) {
	connStr, err := connstring.ParseAndValidate(connectionString)
	if err != nil {
		return nil, err
	}
	options := options.Client().ApplyURI(connectionString)
	c, err := mongo.Connect(context.TODO(), options)
	if err != nil {
		return nil, err
	}
	client := Client{}
	client.SetClient(c)
	db := c.Database(connStr.Database)
	client.SetDatabase(db)

	return &client, nil
}

func (c *Client) SetClient(mongoClient *mongo.Client) {
	c.mongoClient = mongoClient
}

func (c *Client) SetDatabase(db *mongo.Database) {
	c.db = db
}

func (c *Client) SaveList(list *todo.List) error {
	result, err := c.db.Collection(COLLECTION_LIST).InsertOne(context.TODO(), list)
	if err != nil {
		return err
	}

	oid, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return errors.New("invalid ObjectID for newly created list")
	}
	list.ID = oid.Hex()

	return nil
}

func (c *Client) GetList(id string) (*todo.List, error) {
	var list todo.List
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	result := c.db.Collection(COLLECTION_LIST).FindOne(context.TODO(), bson.D{{"_id", oid}})
	if result.Err() != nil {
		return nil, result.Err()
	}
	err = result.Decode(&list)
	if err != nil {
		return nil, err
	}

	return &list, nil
}