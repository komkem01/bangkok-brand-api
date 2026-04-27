package bankaccount

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) Create(ctx context.Context, item *ent.MemberBankAccount) (*ent.MemberBankAccount, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "memberBankAccount.Create")
	defer span.End()

	created, err := s.db.CreateMemberBankAccount(ctx, item)
	if err != nil {
		span.RecordError(err)
		log.Errf("member-bank-account.create.error: %v", err)
		return nil, err
	}

	log.Infof("member-bank-account.create.ok id=%s", created.ID)
	return created, nil
}
