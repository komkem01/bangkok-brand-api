package search

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) Create(ctx context.Context, item *ent.SearchHistory) (*ent.SearchHistory, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "search.Create")
	defer span.End()
	created, err := s.db.CreateSearchHistory(ctx, item)
	if err != nil {
		span.RecordError(err)
		log.Errf("search.create.error: %v", err)
		return nil, err
	}
	return created, nil
}
