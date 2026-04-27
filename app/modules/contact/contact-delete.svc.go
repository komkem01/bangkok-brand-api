package contact

import (
	"context"

	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Delete(ctx context.Context, id uuid.UUID) error {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "contact.Delete")
	defer span.End()

	if err := s.db.DeleteContactByID(ctx, id); err != nil {
		span.RecordError(err)
		log.Errf("contact.delete.error id=%s: %v", id, err)
		return err
	}

	log.Infof("contact.delete.ok id=%s", id)
	return nil
}
