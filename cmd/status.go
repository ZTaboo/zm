package cmd

import (
	"ZM/db"
	"fmt"
	"github.com/liushuochen/gotable"
	"github.com/spf13/cobra"
	"log"
	"net"
	"strconv"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "查看状态",
	Long:  `查看状态`,
	Run:   zStatus,
}

func init() {
	rootCmd.AddCommand(statusCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// statusCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// statusCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// 输出状态信息
func zStatus(cmd *cobra.Command, args []string) {
	table, err := gotable.Create("name", "file", "port", "status")
	if err != nil {
		log.Println("Create table failed: ", err.Error())
		return
	}
	//获取自身运行状态
	if status := getSelfStatus(); status {
		table.AddRow([]string{"ZM Web", "zm.exe", "1207", "Running"})
	} else {
		table.AddRow([]string{"ZM Web", "zm.exe", "1207", "Stop"})
	}
	//获取其它任务运行状态
	tasks, err := db.GetAllTask()
	if err != nil {
		log.Println("GetAllTask failed: ", err.Error())
		return
	}
	for _, item := range tasks {
		var status string
		if item.Status {
			status = "Running"
		} else {
			status = "Stop"
		}
		table.AddRow([]string{item.Name, item.Path, strconv.Itoa(item.Port), status})
	}
	fmt.Println(table)
}

func getSelfStatus() bool {
	listen, err := net.Listen("tcp4", ":1207")
	if err != nil {
		return true
	}
	defer func(listen net.Listener) {
		err := listen.Close()
		if err != nil {
			log.Println("Listen close failed: ", err.Error())
		}
	}(listen)
	return false
}
