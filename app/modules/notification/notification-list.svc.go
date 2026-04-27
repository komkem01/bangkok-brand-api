package notification

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) List(ctx context.Context) ([]*ent.Notification, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "notification.List")
	defer span.End()
	items, err := s.db.ListNotifications(ctx)
	if err != nil {
		span.RecordError(err)
		log.Errf("notification.list.error: %v", err)
		return nil, err
	}
	return items, nil
}
