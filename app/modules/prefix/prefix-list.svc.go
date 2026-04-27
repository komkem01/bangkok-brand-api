package prefix

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) List(ctx context.Context) ([]*ent.Prefix, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "prefix.List")
	defer span.End()

	prefixes, err := s.db.ListPrefixes(ctx)
	if err != nil {
		span.RecordError(err)
		log.Errf("prefix.list.error: %v", err)
		return nil, err
	}

	genders, err := s.dbGender.ListGenders(ctx)
	if err != nil {
		span.RecordError(err)
		log.Errf("prefix.list.gender.error: %v", err)
		return nil, err
	}

	genderMap := make(map[uuid.UUID]string, len(genders))
	for _, g := range genders {
		genderMap[g.ID] = g.NameTh
	}

	for _, p := range prefixes {
		if p.GenderID != nil {
			if name, ok := genderMap[*p.GenderID]; ok {
				p.GenderName = &name
			}
		}
	}

	log.Infof("prefix.list.ok count=%d", len(prefixes))
	return prefixes, nil
}
