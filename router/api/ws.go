package api

import (
	"gin-blog/news/utils"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"net/http"
)

// 客户端连接
func WsClient(context *gin.Context) {
	upGrande := websocket.Upgrader{
		//设置允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		//设置请求协议
		Subprotocols: []string{context.GetHeader("Sec-WebSocket-Protocol")},
	}
	//创建连接
	conn, err := upGrande.Upgrade(context.Writer, context.Request, nil)
	if err != nil {
		zap.L().Info("websocket connect error", zap.String("websocket connect error", context.Param("channel")))
		//log.WSLog(fmt.Sprintf("websocket connect error: %s", context.Param("channel")))
		//format.NewResponseJson(context).Error(51001)
		utils.HandleValidatorError(context, err)
		return
	}
	//生成唯一标识client_id
	var uuid = uuid.NewV4().String()
	//fmt.Println("uuid==",uuid)
	client := &utils.Client{
		Id:      uuid,
		Socket:  conn,
		Message: make(chan []byte, 1024),
	}
	//注册
	utils.WebsocketManager.RegisterClient(client)

	//起协程，实时接收和回复数据
	go client.Read()
	go client.Write()
}
