package user

import (
	"context"
	"go-pkg-oriented-design/internal/platform/identifier"
	"go-pkg-oriented-design/internal/platform/validator"
)

type Service interface {
	List(ctx context.Context) ([]*User, error)
	Find(ctx context.Context, id string) (*User, error)
	Create(ctx context.Context, input *CreateDto) (*User, error)
	Update(ctx context.Context, id string, input *UpdateDto) (*User, error)
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

func (s *service) List(ctx context.Context) ([]*User, error) {
	return s.repository.List(ctx)
}

func (s *service) Find(ctx context.Context, id string) (*User, error) {
	return s.repository.Find(ctx, id)
}

func (s *service) Create(ctx context.Context, input *CreateDto) (*User, error) {
	user := User{
		ID:    s.identifier.NewUuid(),
		Name:  input.Name,
		Email: input.Email,
	}

	if err := s.validator.Validate(user); err != nil {
		return &user, err
	}

	return s.repository.Create(ctx, &user)
}

func (s *service) Update(ctx context.Context, id string, input *UpdateDto) (*User, error) {
	user, err := s.repository.Find(ctx, id)
	if err != nil {
		return nil, err
	}

	user.Name = input.Name
	user.Email = input.Email

	return s.repository.Update(ctx, user)
}

func (s *service) Delete(ctx context.Context, id string) error {
	user, err := s.repository.Find(ctx, id)
	if err != nil {
		return err
	}

	return s.repository.Delete(ctx, user)
}
