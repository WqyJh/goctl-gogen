# goctl-gogen

goctl plugin to parse api definitions and generate api files with handler comments retained.

## Installation

```bash
go install github.com/wqyjh/goctl-gogen@v0.1.0
```

## Usage

Inspired by https://github.com/zeromicro/go-zero/issues/2464.


Define swagger document for handler in `.api` files.

```go
@server(
	group: user
    prefix: /api/v1/user
    jwt: JwtAuth
)
service api {
    // UserSelf godoc
    // @Summary      查询用户信息
    // @Description  查询用户信息
    // @Tags         user
    // @Security     ApiKeyAuth
    // @Accept       json
    // @Produce      json
    // @Param        client   header      string  true  "当前设备类型: android/ios"
    // @Success      200    {object}   types.DataResponse{data=types.UserSelfReply}
    // @Router       /api/v1/user/self [get]
    @handler UserSelf
    get /self (UserSelfReq) returns (UserSelfReply)
}
```

Generate api using goctl api plugin.

```bash
goctl api plugin -plugin goctl-gogen="--home ${CWD}/../template" -api def/main.api -dir ./ -style goZero
```

The document defined would be rendered on handler.

```go
// UserSelf godoc
//	@Summary		查询用户信息
//	@Description	查询用户信息
//	@Tags			user
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			client		header		string	true	"当前设备类型: android/ios"
//	@Success		200			{object}	types.DataResponse{data=types.UserSelfReply}
//	@Router			/api/v1/user/self [get]
func UserSelfHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserSelfReq
		if err := httpz.Parse(r, &req); err != nil {
			httpz.ErrorLog(r.Context(), w, err)
			return
		}

		logc.Infow(r.Context(), "UserSelfHandler", logc.Field("req", &req))

		l := user.NewUserSelfLogic(r.Context(), svcCtx)
		resp, err := l.UserSelf(&req)
		if err != nil {
			httpz.Error(w, err)
			logc.Infow(r.Context(), "UserSelfHandler error", logc.Field("error", err))
		} else {
			httpz.OkJson(w, resp)
			logc.Infow(r.Context(), "UserSelfHandler success", logc.Field("resp", resp))
		}
	}
}
```

Then you can generate swagger document by swaggo.
