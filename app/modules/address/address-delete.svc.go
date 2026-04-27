package address

import (
	"context"

	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Delete(ctx context.Context, id uuid.UUID) error {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "address.Delete")
	defer span.End()

	if err := s.db.DeleteAddressByID(ctx, id); err != nil {
		span.RecordError(err)
		log.Errf("address.delete.error id=%s: %v", id, err)
		return err
	}

	log.Infof("address.delete.ok id=%s", id)
	return nil
}
