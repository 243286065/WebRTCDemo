package main

import (
	"log"

	socketio "github.com/googollee/go-socket.io"
)

type Msg struct {
	UserID    string   `json:"userID"`
	Text      string   `json:"text"`
	State     string   `json:"state"`
	Namespace string   `json:"namespace"`
	Rooms     []string `json:"rooms"`
}

func main() {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	defer server.Close()

	server.OnConnect("/", func(so socketio.Conn) error {
		log.Println("on connection, ID: ", so.ID())

		so.SetContext("")
		msg := Msg{so.ID(), "Connected", "notice", "", nil}
		so.Emit("res", msg)

		return nil
	})

	server.OnEvent("/", "join", func(so socketio.Conn, room string) {
		//加入房间
		so.Join(room)
		msg := Msg{so.ID(), "-----" + so.ID() + " join " + room, "state", so.Namespace(), so.Rooms()}
		log.Println(so.ID(), " join ", room, so.Rooms())

		server.BroadcastToRoom(so.Namespace(), room, "res", msg)
	})

	server.OnEvent("/", "leave", func(so socketio.Conn, room string) {
		so.Leave(room)
		msg := Msg{so.ID(), "-----" + so.ID() + " leave " + room, "state", so.Namespace(), so.Rooms()}
		log.Println(so.ID(), " leave ", room, so.Namespace(), so.Rooms())
		server.BroadcastToRoom(so.Namespace(), room, "res", msg)
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
}
