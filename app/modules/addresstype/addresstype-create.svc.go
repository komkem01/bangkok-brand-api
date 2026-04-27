package addresstype

import (
	"context"
	"strings"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) Create(ctx context.Context, item *ent.AddressType) (*ent.AddressType, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "addresstype.Create")
	defer span.End()

	item.NameEn = strings.ToUpper(strings.TrimSpace(item.NameEn))
	created, err := s.db.CreateAddressType(ctx, item)
	if err != nil {
		span.RecordError(err)
		log.Errf("addresstype.create.error: %v", err)
		return nil, err
	}

	return created, nil
}
