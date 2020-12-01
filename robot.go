package larkwebhook

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type (
	//Robot is a webhook instance which used to send message
	Robot struct {
		botID string
	}

	//Message specific the message send by bot
	//https://www.larksuite.com/hc/en-US/articles/360048487736#1.1.2%20Plain%20text%20messages
	Message interface {
		Type() string //speicit msg_type
		Content() interface{}
	}

	botMessage struct {
		MsgType string      `json:"msg_type"`
		Content interface{} `json:"content"`
	}
)

//NewRobot create a robot with specific botID
func NewRobot(botID string) *Robot {
	return &Robot{
		botID: botID,
	}
}

//Send msg and unmarshal json result into dst
func (r *Robot) Send(ctx context.Context, msg Message, dst interface{}) error {
	bm := botMessage{
		MsgType: msg.Type(),
		Content: msg.Content(),
	}

	raw, err := json.Marshal(bm)
	if err != nil {
		return err
	}
	buf := bytes.NewBuffer(raw)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, r.endPoint(), buf)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if err := json.Unmarshal(body, dst); err != nil {
		return err
	}
	return nil
}

func (r *Robot) endPoint() string {
	u := url.URL{
		Scheme: "https",
		Host:   "open.larksuite.com",
		Path:   fmt.Sprintf("/open-apis/bot/v2/hook/%s", r.botID),
	}
	return u.String()
}
