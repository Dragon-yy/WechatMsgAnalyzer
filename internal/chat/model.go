package chat

import (
	"github.com/go-ego/gse"
	"sort"
	"strconv"
	"strings"
	"time"
)

/*
localId	TalkerId	Type	SubType	IsSender	CreateTime	Status	StrContent	StrTime	Remark	NickName	Sender
*/
// 将词频转换为切片
type WordCount struct {
	Word  string
	Count int
}

// ChatMessage represents a chat message
type ChatMessage struct {
	LocalId    string
	TalkerId   string
	Type       string
	SubType    string
	IsSender   string
	CreateTime string
	Status     string
	StrContent string
	StrTime    string
	Remark     string
	NickName   string
	Sender     string
}

// Report store the analysis result
type Report struct {
	// 报告标题
	Title string
	// 第一条记录时间
	FirstTime string
	// 发言次数统计
	SpeakCount int
	// love词语出现次数
	LoveCount map[string]int
	// 照片出现次数
	ImgCount int
	// emoji出现次数
	EmojiCount int
	// 活跃度分析
	ActiveTime map[string]int
	// 最频繁聊天词汇
	TopWords []WordCount
	// 最长的聊天记录
	LongestMsg ChatMessage
	// 最晚的聊天记录
	LatestMsg ChatMessage
	// 最常用的词
	MostUsedWord map[string]int
}

func (r *Report) analyze(chatMsgs []ChatMessage, title string) {
	// 临时保存聊天信息,以分词
	tmpChatMsgs := ""
	r.LoveCount = make(map[string]int)
	r.ActiveTime = make(map[string]int)
	for _, msg := range chatMsgs {
		// title
		r.Title = title

		// 实现第一条记录时间的逻辑
		// 1. 遍历chatMsgs，找到最早的一条记录
		// 2. 将记录的时间存储到FirstTime中

		if r.FirstTime == "" {
			r.FirstTime = msg.CreateTime
		} else {
			if msg.CreateTime < r.FirstTime {
				r.FirstTime = msg.CreateTime
			}
		}

		// 实现发言次数统计的逻辑
		r.SpeakCount++

		// 爱你 想你 喜欢你词语出现次数
		if strings.Contains(msg.StrContent, "爱你") {
			r.LoveCount["爱你"]++
		} else if strings.Contains(msg.StrContent, "想你") {
			r.LoveCount["想你"]++
		} else if strings.Contains(msg.StrContent, "喜欢你") {
			r.LoveCount["喜欢你"]++
		}

		// 照片出现次数
		if msg.Type == "3" {
			r.ImgCount++
		}
		// emoji出现次数
		if msg.Type == "47" {
			r.EmojiCount++
		}
		// 最长的聊天记录，遇到<msg>标签 xml标签跳过
		if !strings.Contains(msg.StrContent, "<msg>") {
			if r.LongestMsg.StrContent == "" {
				r.LongestMsg = msg
			} else {
				if len(msg.StrContent) > len(r.LongestMsg.StrContent) {
					r.LongestMsg = msg
				}
			}
		}
		// 最晚的聊天记录,遇到<msg>标签 xml标签跳过
		FindLatestMessageDuringNight(&msg, &r.LatestMsg)
		// 调用活跃度分析的逻辑
		analyzeActiveTime(msg, r.ActiveTime)

		// 向tmpChatMsgs中添加聊天内容
		if msg.StrContent != "" && !strings.Contains(msg.StrContent, "<msg>") {
			tmpChatMsgs += msg.StrContent + " "
		}

	}
	// 调用最频繁聊天词汇的逻辑
	r.TopWords = analyzeTopWords(&tmpChatMsgs)
	// 获取最常用的词r.TopWords 第一个
	r.MostUsedWord = map[string]int{r.TopWords[0].Word: r.TopWords[0].Count}
}

func analyzeActiveTime(msg ChatMessage, ActiveTime map[string]int) {
	// 实现活跃度分析的逻辑
	// 1. 创建一个map用于存储用户活跃时间段
	// 2. 遍历chatMsgs，统计每个用户的活跃时间段，时间段包括一年中每月聊天次数、一周中周一到周日的聊天次数、一天中每个时间段的聊天次数
	// 3. 将统计结果存储到map中
	// 4. 返回map
	tmp, _ := strconv.ParseInt(msg.CreateTime, 10, 64)
	tm := time.Unix(tmp, 0)
	//createTime := tm.Format("2006-01-02 15:04:05")

	// 将时间转换为所需的格式
	monthKey := tm.Format("2006-01")    // YYYY-MM 格式
	weekDayKey := tm.Weekday().String() // 周几
	hourKey := tm.Format("15")          // 小时，24小时制
	//fmt.Printf("monthKey: %s, weekDayKey: %s, hourKey: %s\n", monthKey, weekDayKey, hourKey)
	// 检查是否已经为该用户创建了内层的 map，如果没有则创建
	if ActiveTime == nil {
		ActiveTime = make(map[string]int)
	}
	// 更新月份计数
	ActiveTime["month_"+monthKey]++
	// 更新星期计数
	ActiveTime["weekday_"+weekDayKey]++
	// 更新小时计数
	ActiveTime["hour_"+hourKey]++

}

func analyzeTopWords(msg *string) []WordCount {
	// 实现最频繁聊天词汇的逻辑
	// 1. 创建一个map用于存储词汇和出现次数
	// 2. 遍历chatMsgs，统计每条记录的词汇出现次数
	// 3. 将统计结果存储到map中
	// 4. 返回map
	// 创建分词器
	var seg gse.Segmenter
	// 加载词典
	err := seg.LoadDict("zh_s")
	if err != nil {
		panic(err)
	}
	// 分词

	// 过滤<msg>标签
	if strings.Contains(*msg, "<msg>") {
		return nil // 跳过
	}
	// 分词
	segments := seg.Cut(*msg, true)
	// 统计词频，取前100个
	punc := "。，、；：？！“”‘’（）《》【】[]{} < > ？ @ # $ % ^ & * ( ) _ + - = | \\ / ； ： ' ' ； 、 ？ ！ ￥ … （ ） 《 》 【 】"

	// 临时保存词频，以便后续排序
	tmpWordFreqs := make(map[string]int)

	for _, segment := range segments {
		// 忽略标点符号和单字词,因为是中文，所以占用3个字节
		if len(segment) > 3 && !strings.ContainsAny(segment, punc) && !strings.Contains(segment, "<msg>") {
			tmpWordFreqs[segment]++
		}

	}

	var wfSlice []WordCount
	for k, v := range tmpWordFreqs {
		wfSlice = append(wfSlice, WordCount{k, v})
	}

	// 对切片按照词频降序排序
	sort.Slice(wfSlice, func(i, j int) bool {
		return wfSlice[i].Count > wfSlice[j].Count
	})
	// 只保留前100个高频词
	if len(wfSlice) > 100 {
		wfSlice = wfSlice[:100]
	}

	return wfSlice

}

func FindLatestMessageDuringNight(msg, latestMsg *ChatMessage) {
	// 实现最晚的聊天记录的逻辑
	// 1.比较时间即可，6点之后的时间越大越晚，6点之前的时间越小越晚
	tmp, _ := strconv.ParseInt(msg.CreateTime, 10, 64)
	tm := time.Unix(tmp, 0)
	//createTime := tm.Format("2006-01-02 15:04:05")
	if latestMsg.CreateTime == "" {
		*latestMsg = *msg
	} else {
		tmp, _ := strconv.ParseInt(latestMsg.CreateTime, 10, 64)
		latestTime := time.Unix(tmp, 0)
		// 判断当前消息是否在晚间时间范围内
		if tm.Hour() >= 6 && latestTime.Hour() >= 6 {
			if tm.After(latestTime) {
				*latestMsg = *msg
			}
		}
		if tm.Hour() < 6 && latestTime.Hour() < 6 {
			if tm.After(latestTime) {
				*latestMsg = *msg
			}
		}
		if tm.Hour() < 6 && latestTime.Hour() >= 6 {
			*latestMsg = *msg
		}
	}
}

func (r *Report) analyzeEmotion(chatMsgs []ChatMessage) {
	// TODO: 实现聊天情绪分析的逻辑
}

func (r *Report) analyzeTimeLine(data []string) {
	// TODO: 实现时间线分析的逻辑
}
