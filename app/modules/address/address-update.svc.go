package address

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Update(ctx context.Context, id uuid.UUID, item *ent.MemberAddress) (*ent.MemberAddress, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "address.Update")
	defer span.End()

	updated, err := s.db.UpdateAddressByID(ctx, id, item)
	if err != nil {
		span.RecordError(err)
		log.Errf("address.update.error id=%s: %v", id, err)
		return nil, err
	}

	log.Infof("address.update.ok id=%s", id)
	return updated, nil
}
