package bankaccount

import (
	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
	"bangkok-brand/app/utils/base"
	"bangkok-brand/config/i18n"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ListResponse struct {
	ID            string  `json:"id"`
	MemberID      *string `json:"member_id,omitempty"`
	BankID        *string `json:"bank_id,omitempty"`
	AccountNumber *string `json:"account_number,omitempty"`
	AccountName   *string `json:"account_name,omitempty"`
	BranchName    *string `json:"branch_name,omitempty"`
	IsDefault     bool    `json:"is_default"`
	IsVerified    bool    `json:"is_verified"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
}

func toListResponse(a *ent.MemberBankAccount) ListResponse {
	return ListResponse{
		ID:            a.ID.String(),
		MemberID:      uuidToStringPtr(a.MemberID),
		BankID:        uuidToStringPtr(a.BankID),
		AccountNumber: a.AccountNumber,
		AccountName:   a.AccountName,
		BranchName:    a.BranchName,
		IsDefault:     a.IsDefault,
		IsVerified:    a.IsVerified,
		CreatedAt:     a.CreatedAt.Format(utils.RFC3339Milli),
		UpdatedAt:     a.UpdatedAt.Format(utils.RFC3339Milli),
	}
}

func uuidToStringPtr(id *uuid.UUID) *string {
	if id == nil {
		return nil
	}
	s := id.String()
	return &s
}

func (c *Controller) List(ctx *gin.Context) {
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	items, err := c.svc.List(ctx.Request.Context())
	if err != nil {
		log.Errf("member-bank-account.list.error: %v", err)
		base.InternalServerError(ctx, i18n.MemberBankAccountListFailed, nil)
		return
	}

	res := make([]ListResponse, 0, len(items))
	for _, item := range items {
		res = append(res, toListResponse(item))
	}

	base.Success(ctx, res)
}
