package search

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Update(ctx context.Context, id uuid.UUID, item *ent.SearchHistory) (*ent.SearchHistory, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "search.Update")
	defer span.End()
	updated, err := s.db.UpdateSearchHistoryByID(ctx, id, item)
	if err != nil {
		span.RecordError(err)
		log.Errf("search.update.error: %v", err)
		return nil, err
	}
	return updated, nil
}
