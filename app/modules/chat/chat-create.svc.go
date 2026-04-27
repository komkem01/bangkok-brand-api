package chat

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"
)

func (s *Service) Create(ctx context.Context, item *ent.ChatRoom) (*ent.ChatRoom, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "chat.Create")
	defer span.End()
	created, err := s.db.CreateChatRoom(ctx, item)
	if err != nil {
		span.RecordError(err)
		log.Errf("chat.create.error: %v", err)
		return nil, err
	}
	return created, nil
}
