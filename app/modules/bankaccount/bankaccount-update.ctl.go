package bankaccount

import (
	"database/sql"
	"errors"

	"bangkok-brand/app/utils"
	"bangkok-brand/app/utils/base"
	"bangkok-brand/config/i18n"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UpdateUriRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

type UpdateBodyRequest struct {
	MemberID      *string `json:"member_id"`
	BankID        *string `json:"bank_id"`
	AccountNumber *string `json:"account_number"`
	AccountName   *string `json:"account_name"`
	BranchName    *string `json:"branch_name"`
	IsDefault     *bool   `json:"is_default"`
	IsVerified    *bool   `json:"is_verified"`
}

// Update godoc
// PATCH /member-bank-accounts/:id
func (c *Controller) Update(ctx *gin.Context) {
	var uri UpdateUriRequest
	var body UpdateBodyRequest
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	if err := ctx.ShouldBindUri(&uri); err != nil {
		base.BadRequest(ctx, i18n.MemberBankAccountInvalidID, nil)
		return
	}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		base.BadRequest(ctx, i18n.BadRequest, nil)
		return
	}

	id := uuid.MustParse(uri.ID)
	current, err := c.svc.Info(ctx.Request.Context(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			base.BadRequest(ctx, i18n.MemberBankAccountNotFound, nil)
			return
		}
		log.Errf("member-bank-account.update.fetch.error: %v", err)
		base.InternalServerError(ctx, i18n.MemberBankAccountUpdateFailed, nil)
		return
	}

	input := *current
	if body.MemberID != nil {
		v, err := uuid.Parse(*body.MemberID)
		if err != nil {
			base.BadRequest(ctx, i18n.BadRequest, nil)
			return
		}
		input.MemberID = &v
	}
	if body.BankID != nil {
		v, err := uuid.Parse(*body.BankID)
		if err != nil {
			base.BadRequest(ctx, i18n.BadRequest, nil)
			return
		}
		input.BankID = &v
	}
	if body.AccountNumber != nil {
		input.AccountNumber = body.AccountNumber
	}
	if body.AccountName != nil {
		input.AccountName = body.AccountName
	}
	if body.BranchName != nil {
		input.BranchName = body.BranchName
	}
	if body.IsDefault != nil {
		input.IsDefault = *body.IsDefault
	}
	if body.IsVerified != nil {
		input.IsVerified = *body.IsVerified
	}

	updated, err := c.svc.Update(ctx.Request.Context(), id, &input)
	if err != nil {
		log.Errf("member-bank-account.update.error: %v", err)
		base.InternalServerError(ctx, i18n.MemberBankAccountUpdateFailed, nil)
		return
	}

	base.Success(ctx, toListResponse(updated))
}
