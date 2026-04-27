package prefix

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Info(ctx context.Context, id uuid.UUID) (*ent.Prefix, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "prefix.Info")
	defer span.End()

	p, err := s.db.GetPrefixByID(ctx, id)
	if err != nil {
		span.RecordError(err)
		log.Errf("prefix.info.error id=%s: %v", id, err)
		return nil, err
	}

	log.Infof("prefix.info.ok id=%s", id)
	return p, nil
}
