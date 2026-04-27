package bank

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

type UpdateInput struct {
	NameTh   string
	NameEn   string
	Code     string
	IsActive bool
}

func (s *Service) Update(ctx context.Context, id uuid.UUID, input UpdateInput) (*ent.Bank, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "bank.Update")
	defer span.End()

	bank, err := s.db.UpdateBankByID(ctx, id, input.NameTh, input.NameEn, input.Code, input.IsActive)
	if err != nil {
		span.RecordError(err)
		log.Errf("bank.update.error id=%s: %v", id, err)
		return nil, err
	}

	log.Infof("bank.update.ok id=%s", id)
	return bank, nil
}
