package user

import (
	"AuthenticationService/internal/domain"
	"context"
)

func (s *Service) DeleteUser(ctx context.Context, id domain.UserID) error {
	err := s.db.DeleteUser(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
