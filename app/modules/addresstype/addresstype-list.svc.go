package addresstype

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) List(ctx context.Context) ([]*ent.AddressType, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "addresstype.List")
	defer span.End()

	items, err := s.db.ListAddressTypes(ctx)
	if err != nil {
		span.RecordError(err)
		log.Errf("addresstype.list.error: %v", err)
		return nil, err
	}

	return items, nil
}
