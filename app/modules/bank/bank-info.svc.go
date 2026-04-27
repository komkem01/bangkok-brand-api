package bank

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Info(ctx context.Context, id uuid.UUID) (*ent.Bank, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "bank.Info")
	defer span.End()

	bank, err := s.db.GetBankByID(ctx, id)
	if err != nil {
		span.RecordError(err)
		log.Errf("bank.info.error id=%s: %v", id, err)
		return nil, err
	}

	log.Infof("bank.info.ok id=%s", id)
	return bank, nil
}
