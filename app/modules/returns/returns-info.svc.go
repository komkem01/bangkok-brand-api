package returns

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Info(ctx context.Context, id uuid.UUID) (*ent.ReturnRequest, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "returns.Info")
	defer span.End()
	item, err := s.db.GetReturnRequestByID(ctx, id)
	if err != nil {
		span.RecordError(err)
		log.Errf("returns.info.error: %v", err)
		return nil, err
	}
	return item, nil
}
