package returns

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Update(ctx context.Context, id uuid.UUID, item *ent.ReturnRequest) (*ent.ReturnRequest, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "returns.Update")
	defer span.End()
	updated, err := s.db.UpdateReturnRequestByID(ctx, id, item)
	if err != nil {
		span.RecordError(err)
		log.Errf("returns.update.error: %v", err)
		return nil, err
	}
	return updated, nil
}
