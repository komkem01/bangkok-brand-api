package bankaccount

import (
	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
	"bangkok-brand/app/utils/base"
	"bangkok-brand/config/i18n"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateBodyRequest struct {
	MemberID      *string `json:"member_id"`
	BankID        *string `json:"bank_id"`
	AccountNumber *string `json:"account_number"`
	AccountName   *string `json:"account_name"`
	BranchName    *string `json:"branch_name"`
	IsDefault     *bool   `json:"is_default"`
	IsVerified    *bool   `json:"is_verified"`
}

// Create godoc
// POST /member-bank-accounts
func (c *Controller) Create(ctx *gin.Context) {
	var body CreateBodyRequest
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	if err := ctx.ShouldBindJSON(&body); err != nil {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}

	item := &ent.MemberBankAccount{}
	if body.MemberID != nil {
		v, err := uuid.Parse(*body.MemberID)
		if err != nil {
			base.BadRequest(ctx, i18n.BadRequest, nil)
			return
		}
		item.MemberID = &v
	}
	if body.BankID != nil {
		v, err := uuid.Parse(*body.BankID)
		if err != nil {
			base.BadRequest(ctx, i18n.BadRequest, nil)
			return
		}
		item.BankID = &v
	}
	item.AccountNumber = body.AccountNumber
	item.AccountName = body.AccountName
	item.BranchName = body.BranchName
	if body.IsDefault != nil {
		item.IsDefault = *body.IsDefault
	}
	if body.IsVerified != nil {
		item.IsVerified = *body.IsVerified
	}

	created, err := c.svc.Create(ctx.Request.Context(), item)
	if err != nil {
		log.Errf("member-bank-account.create.error: %v", err)
		base.InternalServerError(ctx, i18n.MemberBankAccountCreateFailed, nil)
		return
	}

	base.Success(ctx, toListResponse(created))
}
