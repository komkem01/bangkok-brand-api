package notification

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Update(ctx context.Context, id uuid.UUID, item *ent.Notification) (*ent.Notification, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "notification.Update")
	defer span.End()
	updated, err := s.db.UpdateNotificationByID(ctx, id, item)
	if err != nil {
		span.RecordError(err)
		log.Errf("notification.update.error: %v", err)
		return nil, err
	}
	return updated, nil
}
