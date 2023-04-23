/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/ZTaboo/zm/db"
	"github.com/spf13/cobra"
)

// restartCmd represents the restart command
var restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "重启任务",
	Long:  `重启正在运行的任务`,
	Run:   zRestart,
}

func init() {
	rootCmd.AddCommand(restartCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// restartCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// restartCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// 添加name的flags绑定
	restartCmd.Flags().StringVarP(&name, "name", "n", "", "任务名称")
}

func zRestart(cmd *cobra.Command, args []string) {
	if name == "" {
		fmt.Println("请输入任务名称")
		return
	}
	// 查询任务是否存在
	if res, err := db.TaskExist(name); err != nil {
		fmt.Println("任务不存在：", err)
		return
	} else {
		pid, err := db.StopTask(name)
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
		// 启动
		c := exec.Command(res.Path)
		c.Dir = res.ExecutableFiles
		c.Stdout = nil
		c.Stderr = nil
		err = c.Start()
		if err != nil {
			log.Println("run error:", err)
			return
		}
		db.UpdatePid(res.Name, c.Process.Pid)
		fmt.Println("启动成功\npid:", c.Process.Pid, "port:", res.Port)
	}
}
