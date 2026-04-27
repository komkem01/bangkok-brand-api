package search

import (
	"context"

	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Delete(ctx context.Context, id uuid.UUID) error {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "search.Delete")
	defer span.End()
	err := s.db.DeleteSearchHistoryByID(ctx, id)
	if err != nil {
		span.RecordError(err)
		log.Errf("search.delete.error: %v", err)
	}
	return err
}
