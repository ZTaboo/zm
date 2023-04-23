package cmd

import (
	"fmt"
	"github.com/ZTaboo/zm/db"
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
}

// 输出状态信息
func zStatus(cmd *cobra.Command, args []string) {
	table, err := gotable.Create("Name", "File", "Port", "Status")
	if err != nil {
		log.Println("Create table failed: ", err.Error())
		return
	}
	//获取自身运行状态
	if status := getSelfStatus(); status {
		table.AddRow([]string{"ZM Web", "zm.exe", "1999", "Running"})
	} else {
		table.AddRow([]string{"ZM Web", "zm.exe", "1999", "Stop"})
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
	listen, err := net.Listen("tcp4", ":1999")
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
