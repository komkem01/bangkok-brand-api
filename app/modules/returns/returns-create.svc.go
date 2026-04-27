package returns

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) Create(ctx context.Context, item *ent.ReturnRequest) (*ent.ReturnRequest, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "returns.Create")
	defer span.End()
	created, err := s.db.CreateReturnRequest(ctx, item)
	if err != nil {
		span.RecordError(err)
		log.Errf("returns.create.error: %v", err)
		return nil, err
	}
	return created, nil
}
