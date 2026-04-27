package address

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) List(ctx context.Context) ([]*ent.MemberAddress, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "address.List")
	defer span.End()

	items, err := s.db.ListAddresses(ctx)
	if err != nil {
		span.RecordError(err)
		log.Errf("address.list.error: %v", err)
		return nil, err
	}

	log.Infof("address.list.ok count=%d", len(items))
	return items, nil
}
