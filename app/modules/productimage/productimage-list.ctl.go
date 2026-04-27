package productimage

import (
	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
	"bangkok-brand/app/utils/base"
	"bangkok-brand/config/i18n"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ListResponse struct {
	ID        string  `json:"id"`
	ProductID *string `json:"product_id,omitempty"`
	StorageID *string `json:"storage_id,omitempty"`
	IsMain    bool    `json:"is_main"`
	SortOrder int     `json:"sort_order"`
	CreatedAt string  `json:"created_at"`
}

func toListResponse(p *ent.ProductImage) ListResponse {
	return ListResponse{
		ID:        p.ID.String(),
		ProductID: uuidToStringPtr(p.ProductID),
		StorageID: uuidToStringPtr(p.StorageID),
		IsMain:    p.IsMain,
		SortOrder: p.SortOrder,
		CreatedAt: p.CreatedAt.Format(utils.RFC3339Milli),
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
		log.Errf("product-image.list.error: %v", err)
		base.InternalServerError(ctx, i18n.ProductImageListFailed, nil)
		return
	}

	res := make([]ListResponse, 0, len(items))
	for _, item := range items {
		res = append(res, toListResponse(item))
	}

	base.Success(ctx, res)
}
