package kyc

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) Create(ctx context.Context, item *ent.KYCVerification) (*ent.KYCVerification, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "kyc.Create")
	defer span.End()
	created, err := s.db.CreateKYCVerification(ctx, item)
	if err != nil {
		span.RecordError(err)
		log.Errf("kyc.create.error: %v", err)
		return nil, err
	}
	return created, nil
}
