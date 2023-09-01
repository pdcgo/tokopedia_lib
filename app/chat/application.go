package chat

import (
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"github.com/pdcgo/common_conf/common_concept"
	"github.com/pdcgo/tokopedia_lib/app/chat/api"
	"github.com/pdcgo/tokopedia_lib/app/chat/config"
	"github.com/pdcgo/v2_gots_sdk"
)

type Application struct {
	config *config.AppConfig
	event  *common_concept.CoreEvent
	socket *socketio.Server

	sdk        *v2_gots_sdk.ApiSdk
	mainApi    *api.MainApi
	accountApi *api.AccountApi
	groupApi   *api.GroupApi
	chatApi    *api.ChatApi
}

func NewApplication(
	config *config.AppConfig,
	sdk *v2_gots_sdk.ApiSdk,
	event *common_concept.CoreEvent,
	socket *socketio.Server,
	mainApi *api.MainApi,
	accountApi *api.AccountApi,
	groupApi *api.GroupApi,
	chatApi *api.ChatApi,
) *Application {

	return &Application{
		sdk:        sdk,
		config:     config,
		event:      event,
		socket:     socket,
		mainApi:    mainApi,
		accountApi: accountApi,
		groupApi:   groupApi,
		chatApi:    chatApi,
	}
}

func (app *Application) Run() error {

	// register socketio
	defer app.socket.Close()
	app.sdk.R.GET("/socket.io/*any", gin.WrapH(app.socket))
	app.sdk.R.POST("/socket.io/*any", gin.WrapH(app.socket))

	apiGr := app.sdk.Group("api")
	accountGr := apiGr.Group("akuns")
	groupGr := apiGr.Group("groups")
	chatGr := apiGr.Group("chat")

	app.mainApi.Register(apiGr)
	app.accountApi.Register(accountGr)
	app.groupApi.Register(groupGr)
	app.chatApi.Register(chatGr)

	err := app.sdk.R.Run(app.config.Host + ":" + app.config.Port)
	return err
}
