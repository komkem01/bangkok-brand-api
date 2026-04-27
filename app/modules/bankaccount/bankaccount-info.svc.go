package bankaccount

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Info(ctx context.Context, id uuid.UUID) (*ent.MemberBankAccount, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "memberBankAccount.Info")
	defer span.End()

	item, err := s.db.GetMemberBankAccountByID(ctx, id)
	if err != nil {
		span.RecordError(err)
		log.Errf("member-bank-account.info.error id=%s: %v", id, err)
		return nil, err
	}

	log.Infof("member-bank-account.info.ok id=%s", id)
	return item, nil
}
