package bankaccount

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) List(ctx context.Context) ([]*ent.MemberBankAccount, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "memberBankAccount.List")
	defer span.End()

	items, err := s.db.ListMemberBankAccounts(ctx)
	if err != nil {
		span.RecordError(err)
		log.Errf("member-bank-account.list.error: %v", err)
		return nil, err
	}

	log.Infof("member-bank-account.list.ok count=%d", len(items))
	return items, nil
}
