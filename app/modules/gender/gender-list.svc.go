package gender

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) List(ctx context.Context) ([]*ent.Gender, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "gender.List")
	defer span.End()

	genders, err := s.db.ListGenders(ctx)
	if err != nil {
		span.RecordError(err)
		log.Errf("gender.list.error: %v", err)
		return nil, err
	}

	log.Infof("gender.list.ok count=%d", len(genders))
	return genders, nil
}
