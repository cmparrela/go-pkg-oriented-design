package book

import (
	"context"
	"go-pkg-oriented-design/internal/platform/identifier"
	"go-pkg-oriented-design/internal/platform/validator"
)

type Service interface {
	List(ctx context.Context) ([]*Book, error)
	Find(ctx context.Context, id string) (*Book, error)
	Create(ctx context.Context, input *CreateDto) (*Book, error)
	Update(ctx context.Context, id string, input *UpdateDto) (*Book, error)
	Delete(ctx context.Context, id string) error
}

type service struct {
	repository Repository
	validator  validator.Validator
	identifier identifier.Identifier
}

func NewService(
	repository Repository,
	validator validator.Validator,
	identifier identifier.Identifier,
) Service {
	return &service{
		repository: repository,
		validator:  validator,
		identifier: identifier,
	}
}

func (s *service) List(ctx context.Context) ([]*Book, error) {
	return s.repository.List(ctx)
}

func (s *service) Find(ctx context.Context, id string) (*Book, error) {
	return s.repository.Find(ctx, id)
}

func (s *service) Create(ctx context.Context, input *CreateDto) (*Book, error) {
	book := Book{
		Title:  input.Title,
		Author: input.Author,
	}

	if err := s.validator.Validate(book); err != nil {
		return &book, err
	}

	return s.repository.Create(ctx, &book)
}

func (s *service) Update(ctx context.Context, id string, input *UpdateDto) (*Book, error) {
	book, err := s.repository.Find(ctx, id)
	if err != nil {
		return nil, err
	}

	book.Title = input.Title
	book.Author = input.Author

	return s.repository.Update(ctx, book)
}

func (s *service) Delete(ctx context.Context, id string) error {
	book, err := s.repository.Find(ctx, id)
	if err != nil {
		return err
	}

	return s.repository.Delete(ctx, book)
}
