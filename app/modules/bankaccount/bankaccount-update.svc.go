package bankaccount

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Update(ctx context.Context, id uuid.UUID, input *ent.MemberBankAccount) (*ent.MemberBankAccount, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "memberBankAccount.Update")
	defer span.End()

	updated, err := s.db.UpdateMemberBankAccountByID(ctx, id, input)
	if err != nil {
		span.RecordError(err)
		log.Errf("member-bank-account.update.error id=%s: %v", id, err)
		return nil, err
	}

	log.Infof("member-bank-account.update.ok id=%s", id)
	return updated, nil
}
