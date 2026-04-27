package audit

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Update(ctx context.Context, id uuid.UUID, item *ent.AuditLog) (*ent.AuditLog, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "audit.Update")
	defer span.End()
	updated, err := s.db.UpdateAuditLogByID(ctx, id, item)
	if err != nil {
		span.RecordError(err)
		log.Errf("audit.update.error: %v", err)
		return nil, err
	}
	return updated, nil
}
