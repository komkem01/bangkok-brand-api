package audit

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) List(ctx context.Context) ([]*ent.AuditLog, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "audit.List")
	defer span.End()
	items, err := s.db.ListAuditLogs(ctx)
	if err != nil {
		span.RecordError(err)
		log.Errf("audit.list.error: %v", err)
		return nil, err
	}
	return items, nil
}
