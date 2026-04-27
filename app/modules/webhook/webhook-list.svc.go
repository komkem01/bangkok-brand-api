package webhook

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) List(ctx context.Context) ([]*ent.WebhookEvent, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "webhook.List")
	defer span.End()
	items, err := s.db.ListWebhookEvents(ctx)
	if err != nil {
		span.RecordError(err)
		log.Errf("webhook.list.error: %v", err)
		return nil, err
	}
	return items, nil
}
