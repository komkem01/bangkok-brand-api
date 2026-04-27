package kyc

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Update(ctx context.Context, id uuid.UUID, item *ent.KYCVerification) (*ent.KYCVerification, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "kyc.Update")
	defer span.End()
	updated, err := s.db.UpdateKYCVerificationByID(ctx, id, item)
	if err != nil {
		span.RecordError(err)
		log.Errf("kyc.update.error: %v", err)
		return nil, err
	}
	return updated, nil
}
