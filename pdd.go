package main

import (
	"./tool"
	"github.com/gorilla/websocket"
	"net/url"
)

type PddAuth struct {
	Request_id  string
	Response    string
	Result      string
	Server_time string
	Status      string
	Uid         string
}

type PddChat struct {
	Message    Message
	Request_id string
	Response   string
	Result     string
}

type CmdResponse struct {
	Request_id        string
	Conversations     []Message
	Message           []Message
	Need_unreply_time bool
	Page              int
	Size              int
	Result            string
}

type Message struct {
	Content string
	Ts      string
	From    User
	To      User
	Info    Info
	Type    int
}

type Info struct {
	CustomerNumber int
	GoodsID        string
	GoodsName      string
	GoodsPrice     string
	GoodsThumbUrl  string
	Url            string
	ts             string
}

type Request struct {
	Access_token string
	Pdd          interface{}
}

type SendMessage struct {
	Cmd        string
	Request_id string
	Message    Message
	Random     string
}

type User struct {
	Role string
	Uid  string
}

type Chat_Log struct {
	Id        int
	Text      string
	Time      string
	Token     string
	User      string
	To_user   string
	Form_user string
}

func pddWs(pddConn map[string]string) {
	u := url.URL{Scheme: "wss", Host: *pddaddr, Path: "/"}
	q := u.Query()
	q.Set("access_token", pddConn["access_token"])
	q.Set("role", pddConn["role"])
	q.Set("client", pddConn["client"])
	//q.Set("version", pddConn["version"])
	u.RawQuery = q.Encode()

	if _, ok := pddHub.clients[pddConn["access_token"]]; ok {
		for k, _ := range pddHub.clients[pddConn["access_token"]] {
			pddHub.unregister <- k
		}
	}

	pddClient, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		tool.Writelog(err.Error())
		return
	}

	// init a Client
	pdd := &Client{hub: pddHub, send: make(chan []byte, 256), ws: pddClient, request: pddConn}
	pdd.hub.register <- pdd

	go pdd.writePump()
	go pdd.readPump()
}

func addChat(data *Chat_Log) error {
	stmt, err := engine.Prepare(`INSERT chat_log (text,time,token,user,to_user,from_user) VALUES (?,?,?,?,?,?)`)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(data.Text, data.Time, data.Token, data.User, data.To_user, data.Form_user)
	if err != nil {
		return err
	}

	_, err = res.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}

func getChatCount(to_user string, from_user string, time string) (int, error) {
	count := 0

	rows, err := engine.Query("select count(id) from chat_log where to_user=? and from_user=? and time=?", to_user, from_user, time)
	defer rows.Close()

	if err != nil {
		return count, err
	}

	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			return count, err
		}
	}

	if err = rows.Err(); err != nil {
		return count, err
	}

	return count, nil
}
