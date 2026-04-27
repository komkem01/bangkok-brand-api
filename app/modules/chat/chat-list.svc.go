package chat

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) List(ctx context.Context) ([]*ent.ChatRoom, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "chat.List")
	defer span.End()
	items, err := s.db.ListChatRooms(ctx)
	if err != nil {
		span.RecordError(err)
		log.Errf("chat.list.error: %v", err)
		return nil, err
	}
	return items, nil
}
