package book

import (
	"context"
	"errors"
	"fmt"
	"go-pkg-oriented-design/internal/platform/customerror"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	List(ctx context.Context) ([]*Book, error)
	Find(ctx context.Context, id string) (*Book, error)
	Create(ctx context.Context, book *Book) (*Book, error)
	Update(ctx context.Context, book *Book) (*Book, error)
	Delete(ctx context.Context, book *Book) error
}

type repository struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewRepository(db *mongo.Database) Repository {
	return &repository{
		db,
		db.Collection("books"),
	}
}

func (r *repository) find(ctx context.Context, in interface{}) ([]*Book, error) {
	var bookmarks []*Book

	cursor, err := r.collection.Find(ctx, in)
	if err != nil {
		fmt.Println("repository.find", err.Error())
		return nil, customerror.ErrInternalServerError
	}

	err = cursor.All(ctx, &bookmarks)
	if err != nil {
		return nil, err
	}

	if bookmarks == nil {
		return nil, customerror.ErrNotFound
	}

	return bookmarks, nil
}
func (r *repository) findOne(ctx context.Context, in bson.M) (*Book, error) {
	var data Book

	err := r.collection.FindOne(ctx, in).Decode(&data)
	if err != nil && errors.Is(err, mongo.ErrNoDocuments) {
		return nil, customerror.ErrNotFound
	}
	if err != nil {
		fmt.Println("repository.findOne", err.Error(), err)
		return nil, customerror.ErrInternalServerError
	}

	return &data, nil
}

func (r *repository) List(ctx context.Context) ([]*Book, error) {
	return r.find(ctx, nil)
}

func (r *repository) Find(ctx context.Context, id string) (*Book, error) {
	return r.findOne(ctx, bson.M{"id": id})
}

func (r *repository) Create(ctx context.Context, book *Book) (*Book, error) {
	result, err := r.collection.InsertOne(ctx, book)
	if err != nil {
		fmt.Println("repository.create", err.Error())
		return nil, customerror.ErrInternalServerError
	}

	return r.findOne(ctx, bson.M{"_id": result.InsertedID})
}

func (r *repository) Update(ctx context.Context, book *Book) (*Book, error) {
	filter := bson.M{"id": book.ID}
	update := bson.M{"$set": book}

	_, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		fmt.Println("repository.update", err.Error(), err)
		return nil, customerror.ErrInternalServerError
	}

	return book, nil
}

func (r *repository) Delete(ctx context.Context, book *Book) error {
	id, _ := primitive.ObjectIDFromHex(book.ID)
	_, err := r.find(ctx, bson.M{"_id": id})
	if err != nil {
		return customerror.ErrNotFound
	}

	if _, err := r.collection.DeleteOne(ctx, bson.M{"_id": id}); err != nil {
		fmt.Println("repository.delete", err.Error())
		return customerror.ErrInternalServerError
	}

	return nil
}
