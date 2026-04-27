package prefix

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

type UpdateInput struct {
	GenderID *uuid.UUID
	NameTh   string
	NameEn   string
	IsActive bool
}

func (s *Service) Update(ctx context.Context, id uuid.UUID, input UpdateInput) (*ent.Prefix, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "prefix.Update")
	defer span.End()

	p, err := s.db.UpdatePrefixByID(ctx, id, input.GenderID, input.NameTh, input.NameEn, input.IsActive)
	if err != nil {
		span.RecordError(err)
		log.Errf("prefix.update.error id=%s: %v", id, err)
		return nil, err
	}

	log.Infof("prefix.update.ok id=%s", id)
	return p, nil
}
