package auth

import (
	"context"

	"bangkok-brand/app/utils"
)

func (s *Service) Logout(ctx context.Context) error {
	_, span, _ := utils.NewLogSpan(ctx, s.tracer, "auth.Logout")
	defer span.End()
	return nil
}
