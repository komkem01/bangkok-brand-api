package notification

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) Create(ctx context.Context, item *ent.Notification) (*ent.Notification, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "notification.Create")
	defer span.End()
	created, err := s.db.CreateNotification(ctx, item)
	if err != nil {
		span.RecordError(err)
		log.Errf("notification.create.error: %v", err)
		return nil, err
	}
	return created, nil
}
