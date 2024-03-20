package chat

import (
	"embed"
	"io/fs"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"github.com/pdcgo/common_conf/common_concept"
	"github.com/pdcgo/tokopedia_lib/app/chat/api"
	"github.com/pdcgo/tokopedia_lib/app/chat/config"
	"github.com/pdcgo/v2_gots_sdk"
)

//go:embed assets/frontend/*
var Frontend embed.FS

type Application struct {
	config *config.AppConfig
	event  *common_concept.CoreEvent
	sio    *socketio.Server

	sdk          *v2_gots_sdk.ApiSdk
	mainApi      *api.MainApi
	accountApi   *api.AccountApi
	groupApi     *api.GroupApi
	chatApi      *api.ChatApi
	productApi   *api.ProductApi
	stickerApi   *api.StickerApi
	autoreplyApi *api.AutoReplyApi
}

func NewApplication(
	config *config.AppConfig,
	sdk *v2_gots_sdk.ApiSdk,
	event *common_concept.CoreEvent,
	sio *socketio.Server,
	mainApi *api.MainApi,
	accountApi *api.AccountApi,
	groupApi *api.GroupApi,
	chatApi *api.ChatApi,
	productApi *api.ProductApi,
	stickerApi *api.StickerApi,
	autoreplyApi *api.AutoReplyApi,
) *Application {

	return &Application{
		sdk:          sdk,
		config:       config,
		event:        event,
		sio:          sio,
		mainApi:      mainApi,
		accountApi:   accountApi,
		groupApi:     groupApi,
		chatApi:      chatApi,
		productApi:   productApi,
		stickerApi:   stickerApi,
		autoreplyApi: autoreplyApi,
	}
}

func (app *Application) Run() error {

	httpfs := http.FS(Frontend)

	fs.WalkDir(Frontend, "assets/frontend", func(file string, d fs.DirEntry, err error) error {
		uri := strings.ReplaceAll(file, "assets/frontend", "")
		if uri == "" {
			return nil
		}

		app.sdk.R.StaticFileFS(uri, file, httpfs)
		return nil
	})

	app.sdk.R.GET("/", func(ctx *gin.Context) {
		data, _ := Frontend.ReadFile("assets/frontend/index.html")
		ctx.Data(http.StatusOK, "text/html", data)
	})

	app.sdk.R.NoRoute(func(ctx *gin.Context) {
		data, _ := Frontend.ReadFile("assets/frontend/index.html")
		ctx.Data(http.StatusOK, "text/html", data)
	})

	// register socketio
	defer app.sio.Close()
	app.sdk.R.GET("/socket.io/*any", gin.WrapH(app.sio))
	app.sdk.R.POST("/socket.io/*any", gin.WrapH(app.sio))

	apiGr := app.sdk.Group("api")
	app.mainApi.Register(apiGr)
	app.accountApi.Register(apiGr.Group("akuns"))
	app.groupApi.Register(apiGr.Group("groups"))
	app.chatApi.Register(apiGr.Group("chat"))
	app.productApi.Register(apiGr.Group("product"))
	app.stickerApi.Register(apiGr.Group("sticker"))
	app.autoreplyApi.Register(apiGr.Group("autoreply"))

	err := app.sdk.R.Run(app.config.Host + ":" + app.config.Port)
	return err
}
