package report

import (
	"fmt"
	"github.com/Dragon-yy/WechatMsgAnalyzer/internal/chat"
	"testing"
)

//func TestGenerateWordCloud(t *testing.T) {
//	words := map[string]int{"Hello": 100, "World": 200, "Go": 300, "Echarts": 400, "WordCloud": 500}
//	GenerateWordCloud("Word Cloud Example", words)
//}

func TestFindLatestMessageDuringNight(t *testing.T) {
	chatMsgs := []chat.ChatMessage{
		{
			CreateTime: "1672394267",
		},
		{
			CreateTime: "1672394278",
		},
	}
	lastestMsg := chat.ChatMessage{}
	for _, msg := range chatMsgs {
		chat.FindLatestMessageDuringNight(&msg, &lastestMsg)
	}
	fmt.Println(lastestMsg)
}
