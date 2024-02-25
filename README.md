# Chat Analysis

WechatMsgAnalyzer 是一个用于分析聊天记录并生成报告的命令行工具。

## 功能特性
- 从 Excel 文件中读取聊天记录数据。
- 分析聊天记录，包括统计发言次数、活跃时间段、关键词频率等。
- 生成 HTML 报告，包含图表和统计数据。


## 项目结构
```
/chat-analysis
|-- /cmd
|   |-- root.go          # 根命令
|   |-- analyze.go       # 分析并生成报告命令
|
|-- /internal
|   |-- /chat
|   |   |-- analyzer.go   # 聊天分析逻辑
|   |   |-- parser.go     # 解析Excel文件内容
|   |   |-- model.go      # 数据模型定义
|   |
|   |-- /report
|       |-- generator.go  # 报告生成逻辑
|       |-- template.go   # HTML报告模板
|
|-- /data
|   |-- chat_record.xlsx  # Excel聊天记录文件
|
|-- /web
|
|-- go.mod                # Go模块文件
|-- go.sum                # Go模块依赖版本锁定文件
|-- main.go               # 程序主入口
|-- README.md             # 项目说明文档
```

## 使用
**1. 使用[WeChatMsg](https://github.com/LC044/WeChatMsg)项目导出微信聊天记录到excel文件,筛选你想要生成报告的聊天对象保存到新excel中，例如test.xlsx,excel包含以下字段(WeChatMsg导出文件即包含这些字段)：**
- localId: 通常指的是记录在本地数据库中的唯一标识符，用于区分每条消息或通讯记录。
- TalkerId: 可能指的是对话的另一方的标识符，即消息的接收者或发送者的ID。
- Type: 表示消息的类型，这可能包括文本、图片、视频等。
- SubType: 进一步细分消息的类型，例如在文本消息中可能区分是否包含链接或特殊格式。
- IsSender: 标记消息的发送者，通常用布尔值表示（例如，true表示发送者，false表示接收者）。
- CreateTime: 消息创建的时间戳，表明消息发送或接收的具体时间。
- Status: 消息的状态，可能包括已发送、已接收、已读、错误等状态。
- StrContent: 消息的具体内容，以字符串形式表示。
- StrTime: 可能是消息时间的字符串表示形式，用于显示目的。
- Remark: 可能是对消息或通讯记录的额外注释或备注。
- NickName: 通讯双方的昵称或显示名称。
- Sender: 消息发送者的标识符，可能与TalkerId相对应。

**2. 根据聊天记录excel文件生成json**
```bash
./WechatMsgAnalyzer_MacOS-1.0 analyze -t test -f test.xlsx -o ./test.json
```
**3. 运行serve命令开启后端服务**
```bash
 ./WechatMsgAnalyzer_MacOS-1.0 serve -i 127.0.0.1 -p 9000 -r test.json
```
**4. 运行前端项目,需要注意由于我在前端项目中使用了axios请求后端服务,所以需要修改请求地址为后端服务地址**
```bash
cd web
npm install
npm run serve
```
程序将会分析聊天记录并生成报告，报告将保存为 HTML 文件。
示例报告
![img.png](img/front.png)

## 技术栈
- Go 
- Cobra
- Gin
- Vue

## 贡献
欢迎提出问题和改进建议！如果你有兴趣贡献代码，请提一个 Pull Request。

## 许可证
MIT 许可证

## 鸣谢
使用了以下开源项目：
- [wechat-report](https://github.com/myth984/wechat-report)