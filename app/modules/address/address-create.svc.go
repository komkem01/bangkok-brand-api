package address

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) Create(ctx context.Context, item *ent.MemberAddress) (*ent.MemberAddress, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "address.Create")
	defer span.End()

	created, err := s.db.CreateAddress(ctx, item)
	if err != nil {
		span.RecordError(err)
		log.Errf("address.create.error: %v", err)
		return nil, err
	}

	log.Infof("address.create.ok id=%s", created.ID)
	return created, nil
}
