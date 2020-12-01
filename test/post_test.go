package test

import (
	"context"
	"os"
	"testing"

	"github.com/NadiaSama/larkwebhook"
)

func TestPostSend(t *testing.T) {
	botID := os.Getenv("BOT_ID")
	bot := larkwebhook.NewRobot(botID)

	err := bot.SendPost(context.Background(), "测试标题",
		larkwebhook.PostTag{
			Tag:  larkwebhook.TagText,
			Text: "测试信息1、\n",
		},
		larkwebhook.PostTag{
			Tag:  larkwebhook.TagText,
			Text: "测试信息2",
		},
	)

	if err != nil {
		t.Errorf("send fail error=%s", err.Error())
	}
}
