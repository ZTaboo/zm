package cmd

import (
	"fmt"
	"github.com/ZTaboo/ZM/db"
	"github.com/spf13/cobra"
	"log"
	"os"
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "停止任务",
	Long:  `停止任务.`,
	Run:   zStop,
}

var (
	taskName string
)

func init() {
	rootCmd.AddCommand(stopCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stopCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stopCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	stopCmd.Flags().StringVarP(&taskName, "name", "n", "", "任务名称")
	err := stopCmd.MarkFlagRequired("name")
	if err != nil {
		log.Println(err)
		return
	}
}

func zStop(cmd *cobra.Command, args []string) {
	switch taskName {
	case "web":
		fmt.Println("开发中")
		break
	default:
		pid, err := db.StopTask(taskName)
		if err != nil {
			fmt.Println("任务不存在或已停止")
			return
		}
		// 结束pid
		process, err := os.FindProcess(pid)
		if err != nil {
			log.Println(err)
			return
		}
		if err := process.Kill(); err != nil {
			log.Println("结束进程错误", err)
		} else {
			fmt.Println("结束进程成功")
		}
		db.UpdatePid(taskName, 0)
		break
	}
}
