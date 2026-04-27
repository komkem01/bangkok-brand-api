package address

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Info(ctx context.Context, id uuid.UUID) (*ent.MemberAddress, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "address.Info")
	defer span.End()

	item, err := s.db.GetAddressByID(ctx, id)
	if err != nil {
		span.RecordError(err)
		log.Errf("address.info.error id=%s: %v", id, err)
		return nil, err
	}

	log.Infof("address.info.ok id=%s", id)
	return item, nil
}
