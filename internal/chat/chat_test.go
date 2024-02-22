package chat

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

// 测试
func TestParseChat(t *testing.T) {
	chatMsgs := ParseChat("../../data/test.xlsx")
	if len(chatMsgs) != 0 {
		fmt.Println(chatMsgs[3])
	}
}

func TestAnalyzeChat(t *testing.T) {
	report := AnalyzeChat("test.xlsx", "../../data/test.xlsx")
	fmt.Println(report)
	fmt.Println(report.ActiveTime)
	fmt.Println(report.TopWords)
}

func TestActiveTime(t *testing.T) {
	tmp, _ := strconv.ParseInt("1672394267", 10, 64)
	fmt.Println(tmp)
	tm := time.Unix(tmp, 0)
	// 从字符串解析时间
	createTime := tm.Format("2006-01-02 15:04:05")
	fmt.Println(createTime)

	// 将时间转换为所需的格式
	monthKey := tm.Format("2006-01")    // YYYY-MM 格式
	weekDayKey := tm.Weekday().String() // 周几
	hourKey := tm.Format("15")          // 小时，24小时制
	fmt.Printf("monthKey: %s, weekDayKey: %s, hourKey: %s\n", monthKey, weekDayKey, hourKey)
}
