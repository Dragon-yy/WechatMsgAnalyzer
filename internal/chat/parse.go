package chat

import (
	"github.com/xuri/excelize/v2"
	"sync"
)

// 读取保存在excel中的聊天记录并解析进ChatMessage列表中
// 运行效率26.5s
func ParseChat(filepath string) []ChatMessage {
	// 读取excel文件
	f, err := excelize.OpenFile(filepath)
	if err != nil {
		panic(err)
	}

	// 获取所有行数据
	rows, err := f.GetRows(f.GetSheetName(0))
	if err != nil {
		panic(err)
	}

	var chatMsgs []ChatMessage
	// 批量读取数据并解析
	batchSize := 1000 // 每次读取的行数
	for i := 0; i < len(rows); i += batchSize {
		end := i + batchSize
		if end > len(rows) {
			end = len(rows)
		}
		batchRows := rows[i:end]

		// 解析每一行数据
		for _, row := range batchRows {
			// 解析行数据并添加到chatMsgs列表中
			chatMsg := ChatMessage{}
			chatMsg.LocalId = row[0]
			chatMsg.TalkerId = row[1]
			chatMsg.Type = row[2]
			chatMsg.SubType = row[3]
			chatMsg.IsSender = row[4]
			chatMsg.CreateTime = row[5]
			chatMsg.Status = row[6]
			chatMsg.StrContent = row[7]
			chatMsg.StrTime = row[8]
			chatMsg.Remark = row[9]
			chatMsg.NickName = row[10]
			chatMsg.Sender = row[11]
			// 添加到chatMsgs中
			chatMsgs = append(chatMsgs, chatMsg)
		}
	}

	return chatMsgs[1:]
}

// 运行效率29.61 没有提升，而且由于多goroutine会打乱顺序
func ParseChatMulti(filepath string) []ChatMessage {
	// 读取excel文件
	f, err := excelize.OpenFile(filepath)
	if err != nil {
		panic(err)
	}

	// 获取所有行数据
	rows, err := f.GetRows(f.GetSheetName(0))
	if err != nil {
		panic(err)
	}

	// 定义批量处理大小
	batchSize := 10000
	totalRows := len(rows)

	// 创建等待组
	var wg sync.WaitGroup
	wg.Add((totalRows + batchSize - 1) / batchSize)

	// 创建通道
	ch := make(chan []ChatMessage, (totalRows+batchSize-1)/batchSize)

	// 并发读取数据
	for i := 0; i < totalRows; i += batchSize {
		end := i + batchSize
		if end > totalRows {
			end = totalRows
		}
		go func(start, end int) {
			defer wg.Done()
			var batchMsgs []ChatMessage
			// 解析每一行数据
			for _, row := range rows[start:end] {
				// 解析行数据并添加到chatMsgs列表中
				chatMsg := ChatMessage{}
				chatMsg.LocalId = row[0]
				chatMsg.TalkerId = row[1]
				chatMsg.Type = row[2]
				chatMsg.SubType = row[3]
				chatMsg.IsSender = row[4]
				chatMsg.CreateTime = row[5]
				chatMsg.Status = row[6]
				chatMsg.StrContent = row[7]
				chatMsg.StrTime = row[8]
				chatMsg.Remark = row[9]
				chatMsg.NickName = row[10]
				chatMsg.Sender = row[11]
				// 添加到批量消息中
				batchMsgs = append(batchMsgs, chatMsg)
			}
			// 发送批量消息到通道
			ch <- batchMsgs
		}(i, end)
	}

	// 合并结果
	go func() {
		wg.Wait()
		close(ch)
	}()

	var chatMsgs []ChatMessage
	for msgs := range ch {
		chatMsgs = append(chatMsgs, msgs...)
	}

	return chatMsgs
}
