package bank

import (
	"context"

	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Delete(ctx context.Context, id uuid.UUID) error {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "bank.Delete")
	defer span.End()

	if err := s.db.DeleteBankByID(ctx, id); err != nil {
		span.RecordError(err)
		log.Errf("bank.delete.error id=%s: %v", id, err)
		return err
	}

	log.Infof("bank.delete.ok id=%s", id)
	return nil
}
