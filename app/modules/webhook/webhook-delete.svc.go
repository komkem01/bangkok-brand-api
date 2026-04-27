package webhook

import (
	"context"

	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Delete(ctx context.Context, id uuid.UUID) error {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "webhook.Delete")
	defer span.End()
	err := s.db.DeleteWebhookEventByID(ctx, id)
	if err != nil {
		span.RecordError(err)
		log.Errf("webhook.delete.error: %v", err)
	}
	return err
}
