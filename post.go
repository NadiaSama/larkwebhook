package larkwebhook

import (
	"context"

	"github.com/pkg/errors"
)

type (
	//MsgPost message in lark post format
	MsgPost struct {
		Title string
		Tags  []PostTag
	}

	//PostTag a message tag
	PostTag struct {
		Tag  string `json:"tag"`
		Text string `json:"text"`
	}

	postContent struct {
		Post postPost `json:"post"`
	}

	postPost struct {
		ZhCN zhCnPost `json:"zh_cn"`
	}
	zhCnPost struct {
		Title   string      `json:"title"`
		Content [][]PostTag `json:"content"`
	}

	postResponse struct {
		Code int      `json:"code"`
		Msg  string   `json:"msg"`
		Data postData `json:"data"`
	}

	postData struct {
		MessageID string `json:"message_id"`
	}
)

const (
	//TagText tag for test message
	TagText = "text"
)

func (mp *MsgPost) Type() string {
	return "post"
}

func (mp *MsgPost) Content() interface{} {
	return &postContent{
		postPost{
			zhCnPost{
				Title: mp.Title,
				Content: [][]PostTag{
					mp.Tags,
				},
			},
		},
	}
}

//SendPost build and send post message via robot
func (r *Robot) SendPost(ctx context.Context, title string, tags ...PostTag) error {
	mg := MsgPost{
		Title: title,
		Tags:  tags,
	}

	var resp postResponse
	if err := r.Send(ctx, &mg, &resp); err != nil {
		return err
	}

	if resp.Code != 0 {
		return errors.Errorf("response error code=%d msg=%s", resp.Code, resp.Msg)
	}
	return nil
}
