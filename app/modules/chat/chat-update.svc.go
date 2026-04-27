package chat

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Update(ctx context.Context, id uuid.UUID, item *ent.ChatRoom) (*ent.ChatRoom, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "chat.Update")
	defer span.End()
	updated, err := s.db.UpdateChatRoomByID(ctx, id, item)
	if err != nil {
		span.RecordError(err)
		log.Errf("chat.update.error: %v", err)
		return nil, err
	}
	return updated, nil
}
