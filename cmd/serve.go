package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"net/http"
	"os"
)

// 创建gin服务，监听端口，返回data.json

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the web server",
	Long:  `Start the web server and serve the report data.`,
	Run: func(cmd *cobra.Command, args []string) {
		// 获取参数
		report := cmd.Flag("report").Value.String()
		// 获取port
		port := cmd.Flag("port").Value.String()
		// 获取IP
		ip := cmd.Flag("ip").Value.String()
		// 解析报告数据
		reportData, err := parseReportData(report)
		if err != nil {
			fmt.Println("Error parsing report data:", err)
			return
		}
		// 启动服务器
		startWebServer(ip, port, reportData)
	},
}

func parseReportData(reportPath string) (map[string]interface{}, error) {
	// 读取报告数据文件
	reportBytes, err := os.ReadFile(reportPath)
	if err != nil {
		return nil, err
	}

	// 解析 JSON 数据
	var reportData map[string]interface{}
	err = json.Unmarshal(reportBytes, &reportData)
	if err != nil {
		return nil, err
	}

	return reportData, nil
}

func startWebServer(ip, port string, reportData map[string]interface{}) {
	r := gin.Default()
	// 设置CORS头
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")
		c.Next()
	})
	// 添加路由
	r.GET("/report", func(c *gin.Context) {
		c.JSON(http.StatusOK, reportData)
	})

	// 启动服务器
	fmt.Println(fmt.Sprintf("Listening on %s:%s", ip, port))
	err := r.Run(fmt.Sprintf("%s:%s", ip, port))
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func init() {
	serveCmd.Flags().StringP("report", "r", "data/data.json", "The report data file to be served")
	// 添加port
	serveCmd.Flags().StringP("port", "p", "9000", "The port to listen on")
	// 添加IP
	serveCmd.Flags().StringP("ip", "i", "127.0.0.1", "The IP to listen on")
	RootCmd.AddCommand(serveCmd)

}
