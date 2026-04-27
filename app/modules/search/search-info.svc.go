package search

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Info(ctx context.Context, id uuid.UUID) (*ent.SearchHistory, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "search.Info")
	defer span.End()
	item, err := s.db.GetSearchHistoryByID(ctx, id)
	if err != nil {
		span.RecordError(err)
		log.Errf("search.info.error: %v", err)
		return nil, err
	}
	return item, nil
}
