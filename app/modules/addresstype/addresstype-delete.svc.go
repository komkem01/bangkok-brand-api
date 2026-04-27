package addresstype

import (
	"context"

	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Delete(ctx context.Context, id uuid.UUID) error {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "addresstype.Delete")
	defer span.End()

	if err := s.db.DeleteAddressTypeByID(ctx, id); err != nil {
		span.RecordError(err)
		log.Errf("addresstype.delete.error id=%s: %v", id, err)
		return err
	}

	return nil
}
