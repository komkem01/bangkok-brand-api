package bank

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) List(ctx context.Context) ([]*ent.Bank, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "bank.List")
	defer span.End()

	banks, err := s.db.ListBanks(ctx)
	if err != nil {
		span.RecordError(err)
		log.Errf("bank.list.error: %v", err)
		return nil, err
	}

	log.Infof("bank.list.ok count=%d", len(banks))
	return banks, nil
}
