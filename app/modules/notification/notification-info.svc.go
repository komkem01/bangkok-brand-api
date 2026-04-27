package notification

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Info(ctx context.Context, id uuid.UUID) (*ent.Notification, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "notification.Info")
	defer span.End()
	item, err := s.db.GetNotificationByID(ctx, id)
	if err != nil {
		span.RecordError(err)
		log.Errf("notification.info.error: %v", err)
		return nil, err
	}
	return item, nil
}
