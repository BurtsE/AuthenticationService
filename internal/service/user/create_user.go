package user

import (
	"AuthenticationService/internal/model"
	"context"
)

func (s *Service) CreateUser(ctx context.Context, User *model.User) error {
	err := s.db.CreateUser(ctx, User)
	if err != nil {
		return err
	}
	return nil
}
