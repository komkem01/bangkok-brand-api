package kyc

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Info(ctx context.Context, id uuid.UUID) (*ent.KYCVerification, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "kyc.Info")
	defer span.End()
	item, err := s.db.GetKYCVerificationByID(ctx, id)
	if err != nil {
		span.RecordError(err)
		log.Errf("kyc.info.error: %v", err)
		return nil, err
	}
	return item, nil
}
