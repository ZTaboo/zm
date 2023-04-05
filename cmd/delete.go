package cmd

import (
	"fmt"
	"github.com/ZTaboo/ZM/db"
	"github.com/spf13/cobra"
	"log"
	"os"
)

// addCmd represents the add command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "删除任务",
	Long:  `删除任务.`,
	Run:   zDelete,
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

func zDelete(cmd *cobra.Command, args []string) {
	if len(args) <= 0 {
		fmt.Println("参数错误")
		return
	}
	switch args[0] {
	case "web":
		fmt.Println("web任务为内置任务，不能删除")
		break
	default:
		//停止任务
		pid, err := db.StopTask(args[0])
		if err != nil {
			fmt.Println("任务不存在或未运行：", err)
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
			db.UpdatePid(taskName, 0)
			// 删除任务
			if err := db.DeleteTask(args[0]); err != nil {
				fmt.Println("删除失败：", err)
				return
			}
			fmt.Println("删除成功")
		}
		break
	}
}
