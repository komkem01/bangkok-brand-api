package prefix

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) List(ctx context.Context) ([]*ent.Prefix, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "prefix.List")
	defer span.End()

	prefixes, err := s.db.ListPrefixes(ctx)
	if err != nil {
		span.RecordError(err)
		log.Errf("prefix.list.error: %v", err)
		return nil, err
	}

	log.Infof("prefix.list.ok count=%d", len(prefixes))
	return prefixes, nil
}
