package user

import (
	"AuthenticationService/internal/domain"
	"context"
	"fmt"
)

func (s *Service) DeleteUser(ctx context.Context, id domain.UserID) error {
	err := s.userStorage.DeleteUser(ctx, id)
	if err != nil {
		return fmt.Errorf("%w: %s", domain.ErrDatabaseConflict, err)
	}
	return nil
}
