package webhook

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) Create(ctx context.Context, item *ent.WebhookEvent) (*ent.WebhookEvent, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "webhook.Create")
	defer span.End()
	created, err := s.db.CreateWebhookEvent(ctx, item)
	if err != nil {
		span.RecordError(err)
		log.Errf("webhook.create.error: %v", err)
		return nil, err
	}
	return created, nil
}
