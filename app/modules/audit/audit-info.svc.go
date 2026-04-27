package audit

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Info(ctx context.Context, id uuid.UUID) (*ent.AuditLog, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "audit.Info")
	defer span.End()
	item, err := s.db.GetAuditLogByID(ctx, id)
	if err != nil {
		span.RecordError(err)
		log.Errf("audit.info.error: %v", err)
		return nil, err
	}
	return item, nil
}
