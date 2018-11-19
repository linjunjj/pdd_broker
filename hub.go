// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"./tool"
	"encoding/json"
	"fmt"
	"strings"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[string]map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan *Client

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client

	//0为客户端，1为拼多多
	tag int
}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan *Client),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[string]map[*Client]bool),
		tag:        0,
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			for _, v := range strings.Split(client.request["access_token"], ",") {
				if _, ok := h.clients[v]; ok {
					h.clients[v][client] = true
				} else {
					h.clients[v] = map[*Client]bool{client: true}
				}

				if h.tag == 0 {
					pddConn := make(map[string]string)
					pddConn["role"] = client.request["role"]
					pddConn["client"] = client.request["client"]
					pddConn["version"] = client.request["version"]
					pddConn["access_token"] = v

					go pddWs(pddConn)
				}
			}
		case client := <-h.unregister:
			for _, v := range strings.Split(client.request["access_token"], ",") {
				if _, ok := h.clients[v][client]; ok {
					delete(h.clients[v], client)
					close(client.send)
				}

				if len(h.clients[v]) == 0 {
					delete(h.clients, v)
				}
			}

		case client := <-h.broadcast:
			msg := string(client.message)
			fmt.Println(msg)
			switch h.tag {
			case 0:
				req := Request{}
				err := json.Unmarshal(client.message, &req)
				if err != nil {
					tool.Writelog("解析聊天数据错误：" + err.Error())
					continue
				}

				if _, ok := pddHub.clients[req.Access_token]; !ok || req.Access_token == "" {
					continue
				}

				pddClient := pddHub.clients[req.Access_token]

				str, err := json.Marshal(req.Pdd)
				if err != nil {
					tool.Writelog("解析聊天数据错误：" + err.Error())
					continue
				}

				if strings.Contains(msg, `"cmd":"send_message"`) {
					send := SendMessage{}
					json.Unmarshal(str, &send)

					data := Chat_Log{}
					data.Text = send.Message.Content
					data.Form_user = send.Message.From.Uid
					data.To_user = send.Message.To.Uid
					data.Time = send.Message.Ts
					data.Token = req.Access_token

					err = addChat(&data)
					if err != nil {
						tool.Writelog("聊天数据入库错误：" + err.Error())
						continue
					}

					chat := PddChat{}
					chat.Request_id = send.Request_id
					chat.Response = "clientpush"
					chat.Result = "ok"
					chat.Message = send.Message

					message := Request{}
					message.Access_token = req.Access_token
					message.Pdd = chat

					msg, err := json.Marshal(message)
					if err != nil {
						tool.Writelog("解析聊天数据错误：" + err.Error())
						continue
					}

					for k, _ := range h.clients[req.Access_token] {
						select {
						case k.send <- msg:
						default:
							h.unregister <- k
						}
					}
				}

				for k, _ := range pddClient {
					select {
					case k.send <- str:
					default:
						h.unregister <- k
					}
				}

			case 1:
				if _, ok := clientHub.clients[client.request["access_token"]]; !ok {
					continue
				}

				userClient := clientHub.clients[client.request["access_token"]]

				if strings.Contains(msg, `"response":"push"`) && strings.Contains(msg, `"message":`) {
					var chat PddChat
					err := json.Unmarshal(client.message, &chat)
					if err != nil {
						tool.Writelog("解析聊天数据错误：" + err.Error())
						continue
					}

					data := Chat_Log{}
					data.Text = chat.Message.Content

					data.Form_user = chat.Message.From.Uid
					data.To_user = chat.Message.To.Uid
					data.Time = chat.Message.Ts
					data.Token = client.request["access_token"]

					err = addChat(&data)
					if err != nil {
						tool.Writelog("聊天数据入库错误：" + err.Error())
						continue
					}

				} else if strings.Contains(msg, `"response":"latest_conversations"`) {
					var response CmdResponse
					err := json.Unmarshal(client.message, &response)
					if err != nil {
						tool.Writelog("解析聊天数据错误：" + err.Error())
						continue
					}

					count, err := getChatCount(response.Conversations[0].To.Uid, response.Conversations[0].From.Uid, response.Conversations[0].Ts)
					if err != nil {
						tool.Writelog("数据查询错误：" + err.Error())
						continue
					}

					if count <= 0 {
						data := Chat_Log{}
						//if response.Conversations[0].Info.GoodsID != "" {
						//	response.Conversations[0].Info.Url = response.Conversations[0].Content
						//	info, err := json.Marshal(response.Conversations[0].Info)
						//	if err != nil {
						//		tool.Writelog("解析聊天内容错误：" + err.Error())
						//		continue
						//	}
						//	data.Text = string(info)
						//} else {
						data.Text = response.Conversations[0].Content
						//}

						data.Form_user = response.Conversations[0].From.Uid
						data.To_user = response.Conversations[0].To.Uid
						data.Time = response.Conversations[0].Ts
						data.Token = client.request["access_token"]

						err = addChat(&data)
						if err != nil {
							tool.Writelog("聊天数据入库错误：" + err.Error())
							continue
						}
					}
				}
				req := Request{}
				req.Access_token = client.request["access_token"]
				err := json.Unmarshal(client.message, &req.Pdd)
				if err != nil {
					tool.Writelog("解析聊天数据错误：" + err.Error())
					continue
				}

				str, err := json.Marshal(req)
				if err != nil {
					tool.Writelog("解析聊天数据错误：" + err.Error())
					continue
				}
				//fmt.Println(string(str));

				for k, _ := range userClient {
					select {
					case k.send <- str:
					default:
						h.unregister <- k
					}
				}
			}
		}
	}
}
