package webhook

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Info(ctx context.Context, id uuid.UUID) (*ent.WebhookEvent, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "webhook.Info")
	defer span.End()
	item, err := s.db.GetWebhookEventByID(ctx, id)
	if err != nil {
		span.RecordError(err)
		log.Errf("webhook.info.error: %v", err)
		return nil, err
	}
	return item, nil
}
