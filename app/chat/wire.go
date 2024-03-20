//go:build wireinject
// +build wireinject

package chat

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/pdcgo/common_conf/common_concept"
	"github.com/pdcgo/tokopedia_lib/app/chat/api"
	"github.com/pdcgo/tokopedia_lib/app/chat/config"
	"github.com/pdcgo/tokopedia_lib/app/chat/group"
	"github.com/pdcgo/tokopedia_lib/app/chat/helper"
	"github.com/pdcgo/tokopedia_lib/app/chat/repo"
	"github.com/pdcgo/tokopedia_lib/app/chat/service"
	"github.com/pdcgo/tokopedia_lib/lib/api_public"
)

func InitApplication(cfg *config.AppConfig) (*Application, error) {

	wire.Build(
		gin.Default,
		CreateChatSdk,
		CreateSqliteDatabase,
		common_concept.NewCoreEvent,
		api_public.NewTokopediaApiPublic,
		config.NewInitConfig,
		config.NewAutoReplyConfig,
		repo.NewAccountRepo,
		repo.NewGroupRepo,
		group.NewDriverGroup,
		group.NewSocketGroup,
		group.NewChatGroup,
		helper.NewSoundPlayer,
		service.NewAccountService,
		service.NewChatService,
		service.NewNotificationService,
		api.NewBaseDriverApi,
		api.NewMainApi,
		api.NewAccountApi,
		api.NewChatApi,
		api.NewGroupApi,
		api.NewProductApi,
		api.NewStickerApi,
		api.NewAutoReplyApi,
		CreateSocketIO,
		NewApplication,
	)

	return &Application{}, nil
}
