package report

import (
	"encoding/json"
	"github.com/Dragon-yy/WechatMsgAnalyzer/internal/chat"
	"os"
)

//func GenerateWordCloud(title string, words map[string]int) {
//	// 创建词云图
//	wordCloud := charts.NewWordCloud()
//	wordCloud.SetGlobalOptions(
//		charts.WithTitleOpts(opts.Title{
//			Title: title,
//		}),
//		charts.WithTooltipOpts(opts.Tooltip{Show: true}),
//	)
//
//	var wcData []opts.WordCloudData
//	for word, count := range words {
//		wcData = append(wcData, opts.WordCloudData{
//			Name:  word,
//			Value: count,
//		})
//	}
//
//	wordCloud.AddSeries("wordCloud", wcData).
//		SetSeriesOptions(
//			charts.WithWorldCloudChartOpts(
//				opts.WordCloudChart{
//					SizeRange: []float32{14, 80},
//				}),
//		)
//
//	// 生成HTML文件
//	//f, _ := os.Create("word_cloud.html")
//	//wordCloud.Render(f)
//}

// 把report数据格式化为json并保存在data/data.json中
func GenerateReportData(data chat.Report, output string) {
	jsonReport, _ := json.Marshal(data)
	err := os.WriteFile(output, jsonReport, 0644)
	if err != nil {
		panic(err)
	}
}
