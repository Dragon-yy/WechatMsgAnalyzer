package cmd

import (
	"fmt"
	"github.com/Dragon-yy/WechatMsgAnalyzer/internal/chat"
	reportPgk "github.com/Dragon-yy/WechatMsgAnalyzer/internal/report"
	"github.com/spf13/cobra"
)

// analyzeCmd represents the analyze command
var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Analyze Wechat messages from an Excel file",
	Long:  `Analyze Wechat messages from an Excel file and generate a report.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("analyze called")
		// 获取命令行参数
		excelFile := cmd.Flag("file").Value.String()
		reportTitle := cmd.Flag("title").Value.String()
		outputFile := cmd.Flag("output").Value.String()

		// 这里调用internal包的函数
		report := chat.AnalyzeChat(reportTitle, excelFile)
		//fmt.Println(len(report.TopWords))
		//
		//fmt.Println(report)
		//fmt.Println(report.ActiveTime)
		//fmt.Println(report.TopWords)

		// 把report json序列化后保存在data/data.json中
		reportPgk.GenerateReportData(report, outputFile)

		// 生成词云
		//reportPgk.GenerateWordCloud(reportTitle, report.TopWords)
	},
}

func init() {
	analyzeCmd.Flags().StringP("file", "f", "", "The Excel file to be analyzed")
	analyzeCmd.MarkFlagRequired("file")
	analyzeCmd.Flags().StringP("title", "t", "", "The title of the report")
	analyzeCmd.MarkFlagRequired("title")
	analyzeCmd.Flags().StringP("output", "o", "data/data.json", "The output file of the report")
	analyzeCmd.MarkFlagRequired("output")

	RootCmd.AddCommand(analyzeCmd)
}
