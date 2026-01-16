package user

import (
	"AuthenticationService/internal/model"
	"context"
)

func (s *Service) DeleteUser(ctx context.Context, id model.UserID) error {
	err := s.db.DeleteUser(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
