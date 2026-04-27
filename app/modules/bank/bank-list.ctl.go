package bank

import (
	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
	"bangkok-brand/app/utils/base"
	"bangkok-brand/config/i18n"

	"github.com/gin-gonic/gin"
)

type ListResponse struct {
	ID        string `json:"id"`
	NameTh    string `json:"name_th"`
	NameEn    string `json:"name_en"`
	Code      string `json:"code"`
	IsActive  bool   `json:"is_active"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func toListResponse(b *ent.Bank) ListResponse {
	return ListResponse{
		ID:        b.ID.String(),
		NameTh:    b.NameTh,
		NameEn:    b.NameEn,
		Code:      b.Code,
		IsActive:  b.IsActive,
		CreatedAt: b.CreatedAt.Format(utils.RFC3339Milli),
		UpdatedAt: b.UpdatedAt.Format(utils.RFC3339Milli),
	}
}

func (c *Controller) List(ctx *gin.Context) {
	span, log := utils.LogSpanFromGin(ctx)
	defer span.End()

	banks, err := c.svc.List(ctx.Request.Context())
	if err != nil {
		log.Errf("bank.list.error: %v", err)
		base.InternalServerError(ctx, i18n.BankListFailed, nil)
		return
	}

	res := make([]ListResponse, 0, len(banks))
	for _, b := range banks {
		res = append(res, toListResponse(b))
	}

	base.Success(ctx, res)
}
