package con_data

import (
	// "encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

type TryUser struct {
	Email    string
	Password string
}

type AUser struct {
	FullName string `json:"fullnama"`
	Alias    string `json:"alias"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type Msg struct {
	Jwt  string `json:"jwt"`
	Type string `json:"type"`
	Data string `json:"data"`
}

func SendMsg(j string, t string, d string, c websocket.Conn) {
	m := Msg{j, t, d}
	if err := c.WriteJSON(m); err != nil {
		fmt.Println(err)
	}

	// mm := json.Marshal(m)
	// fmt.Println(string(mm))
}
