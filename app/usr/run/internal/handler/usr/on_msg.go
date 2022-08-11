package usr

import (
	"context"
	"github.com/gotd/td/tg"
	"github.com/iyear/searchx/app/usr/run/internal/util"
)

func OnNewMessage(ctx context.Context, e tg.Entities, update *tg.UpdateNewMessage) error {
	return indexMessage(util.GetUsrScope(ctx), e, update.Message)
}

func OnEditMessage(ctx context.Context, e tg.Entities, update *tg.UpdateEditMessage) error {
	return indexMessage(util.GetUsrScope(ctx), e, update.Message)
}

func OnNewScheduledMessage(ctx context.Context, e tg.Entities, update *tg.UpdateNewScheduledMessage) error {
	return indexMessage(util.GetUsrScope(ctx), e, update.Message)
}

func OnNewChannelMessage(ctx context.Context, e tg.Entities, update *tg.UpdateNewChannelMessage) error {
	return indexMessage(util.GetUsrScope(ctx), e, update.Message)
}

func OnEditChannelMessage(ctx context.Context, e tg.Entities, update *tg.UpdateEditChannelMessage) error {
	return indexMessage(util.GetUsrScope(ctx), e, update.Message)
}
