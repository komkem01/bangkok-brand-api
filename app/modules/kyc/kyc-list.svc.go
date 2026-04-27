package kyc

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) List(ctx context.Context) ([]*ent.KYCVerification, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "kyc.List")
	defer span.End()
	items, err := s.db.ListKYCVerifications(ctx)
	if err != nil {
		span.RecordError(err)
		log.Errf("kyc.list.error: %v", err)
		return nil, err
	}
	return items, nil
}
