package bank

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) Create(ctx context.Context, item *ent.Bank) (*ent.Bank, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "bank.Create")
	defer span.End()

	created, err := s.db.CreateBank(ctx, item)
	if err != nil {
		span.RecordError(err)
		log.Errf("bank.create.error: %v", err)
		return nil, err
	}

	log.Infof("bank.create.ok id=%s", created.ID)
	return created, nil
}
