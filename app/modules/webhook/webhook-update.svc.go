package webhook

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Update(ctx context.Context, id uuid.UUID, item *ent.WebhookEvent) (*ent.WebhookEvent, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "webhook.Update")
	defer span.End()
	updated, err := s.db.UpdateWebhookEventByID(ctx, id, item)
	if err != nil {
		span.RecordError(err)
		log.Errf("webhook.update.error: %v", err)
		return nil, err
	}
	return updated, nil
}
