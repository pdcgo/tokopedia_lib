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
	"github.com/pdcgo/tokopedia_lib/app/chat/repo"
	"github.com/pdcgo/tokopedia_lib/app/chat/service"
)

func InitApplication(cfg *config.AppConfig) (*Application, error) {

	wire.Build(
		gin.Default,
		CreateChatSdk,
		CreateSqliteDatabase,
		common_concept.NewCoreEvent,
		config.NewInitConfig,
		repo.NewAccountRepo,
		group.NewDriverGroup,
		group.NewSocketGroup,
		group.NewChatGroup,
		service.NewAccountService,
		api.NewMainApi,
		api.NewAccountApi,
		api.NewChatApi,
		api.NewGroupApi,
		CreateSocketIO,
		NewApplication,
	)

	return &Application{}, nil
}
