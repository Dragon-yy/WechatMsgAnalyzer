package chat

func AnalyzeChat(title, file string) Report {
	/*
		- 发言次数统计：根据 TalkerId 或 NickName 字段统计每个用户的发言次数。
		- 活跃度分析：分析每个用户的活跃时间段，可以根据 CreateTime 字段中的时间信息来统计用户在一天中的活跃时间。
		- 聊天内容分析：可以使用 StrContent 字段进行聊天内容的情感分析、关键词提取等。
		- 聊天情绪分析：根据聊天内容进行情绪分析，可以使用自然语言处理技术，如情感词典或机器学习模型来分析。
		- 用户互动分析：根据 IsSender 字段来分析每个用户是发送者还是接收者，进而分析用户之间的互动情况。
		- 聊天时间分析：分析聊天记录的时间分布，可以根据 CreateTime 字段来统计每天、每周、每月的聊天活动情况。
		- 聊天内容类型分析：根据 Type 和 SubType 字段来分析聊天内容的类型，如文本消息、图片消息、语音消息等。
		- 用户间关系分析：根据用户的互动情况和聊天内容，分析用户之间的关系密度和联系频率
		- 词云展示：根据聊天内容生成词云图，展示聊天记录中的关键词。
		- 图表展示：使用图表展示聊天记录的统计数据，如发言次数、活跃时间段、聊天内容类型等。
	*/
	//msgChats := ParseChat("../../data/test.xlsx")
	msgChats := ParseChat(file)
	report := Report{}
	report.analyze(msgChats, title)
	return report
}
