package audit

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) Create(ctx context.Context, item *ent.AuditLog) (*ent.AuditLog, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "audit.Create")
	defer span.End()
	created, err := s.db.CreateAuditLog(ctx, item)
	if err != nil {
		span.RecordError(err)
		log.Errf("audit.create.error: %v", err)
		return nil, err
	}
	return created, nil
}
