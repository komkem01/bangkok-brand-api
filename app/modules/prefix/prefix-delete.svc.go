package prefix

import (
	"context"

	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Delete(ctx context.Context, id uuid.UUID) error {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "prefix.Delete")
	defer span.End()

	if err := s.db.DeletePrefixByID(ctx, id); err != nil {
		span.RecordError(err)
		log.Errf("prefix.delete.error id=%s: %v", id, err)
		return err
	}

	log.Infof("prefix.delete.ok id=%s", id)
	return nil
}
