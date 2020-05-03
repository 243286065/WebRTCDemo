package handler

import (
	"log"

	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
)

var (
	server *socketio.Server
	err    error
)

const (
	MaxUserCnt = 2
)

type Msg struct {
	UserID    string   `json:"userID"`
	Text      string   `json:"text"`
	State     string   `json:"state"`
	Namespace string   `json:"namespace"`
	Rooms     []string `json:"rooms"`
}

func init() {
	server, err = socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	server.OnConnect("/", func(so socketio.Conn) error {
		log.Println("on connection, ID: ", so.ID())

		so.SetContext("")
		msg := Msg{so.ID(), "Connected", "notice", "", nil}
		so.Emit("res", msg)

		return nil
	})

	server.OnEvent("/", "join", func(so socketio.Conn, room string) {
		if server.RoomLen(so.Namespace(), room) >= MaxUserCnt {
			//房间已满
			so.Emit("full", room)
			return
		}

		//加入房间
		so.Join(room)
		log.Println(so.ID(), " join ", room, so.Rooms())
		//全员发送joined消息，客户端自己判断是否是有新用户加入
		server.BroadcastToRoom(so.Namespace(), room, "joined", room, so.ID())
	})

	//处理用户离开消息
	server.OnEvent("/", "leave", func(so socketio.Conn, room string) {
		log.Println(so.ID(), " leave ", room, so.Namespace(), so.Rooms())
		server.BroadcastToRoom(so.Namespace(), room, "leaved", room, so.ID())

		so.Leave(room)
	})

	server.OnEvent("/", "message", func(so socketio.Conn, room string, msg interface{}) {
		//原封不动地转发
		server.BroadcastToRoom(so.Namespace(), room, "message", room, so.ID(), msg)
	})

	server.OnEvent("/", "ready", func(so socketio.Conn, room string) {
		//原封不动地转发
		server.BroadcastToRoom(so.Namespace(), room, "ready", room, so.ID())
	})

	server.OnEvent("/", "chat", func(so socketio.Conn, msg string) {
		res := Msg{so.ID(), "----" + msg, "normal", so.Namespace(), so.Rooms()}
		so.SetContext(res)
		log.Println("chat receive", msg, so.Namespace(), so.Rooms(), server.Rooms(so.Namespace()))
		rooms := so.Rooms()

		for _, room := range rooms {
			server.BroadcastToRoom(so.Namespace(), room, "res", res)
		}

	})

	go server.Serve()
}

func SocketIOServerHandler(c *gin.Context) {

	//server.OnEvent("/", "notice")
	if server != nil {
		log.Println("WebSocket server start...")
		server.ServeHTTP(c.Writer, c.Request)
	}
}
