package chat

import (
	"context"

	"bangkok-brand/app/modules/entities/ent"
	"bangkok-brand/app/utils"

	"github.com/google/uuid"
)

func (s *Service) Info(ctx context.Context, id uuid.UUID) (*ent.ChatRoom, error) {
	ctx, span, log := utils.NewLogSpan(ctx, s.tracer, "chat.Info")
	defer span.End()
	item, err := s.db.GetChatRoomByID(ctx, id)
	if err != nil {
		span.RecordError(err)
		log.Errf("chat.info.error: %v", err)
		return nil, err
	}
	return item, nil
}
