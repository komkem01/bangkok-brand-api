package search

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) List(ctx context.Context) ([]*ent.SearchHistory, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "search.List")
	defer span.End()
	items, err := s.db.ListSearchHistories(ctx)
	if err != nil {
		span.RecordError(err)
		log.Errf("search.list.error: %v", err)
		return nil, err
	}
	return items, nil
}
