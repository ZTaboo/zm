/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/ZTaboo/zm/db"
	"log"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "添加任务",
	Long:  `添加任务.`,
	Run:   zAdd,
}

var (
	name string
	port int
)

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addCmd.Flags().StringVarP(&name, "name", "n", "", "为任务命名")
	addCmd.Flags().IntVarP(&port, "port", "p", 0, "任务程序端口")
	// 保证flags必填
	if err := addCmd.MarkFlagRequired("name"); err != nil {
		log.Println(err)
		return
	}
	if err := addCmd.MarkFlagRequired("port"); err != nil {
		log.Println(err)
		return
	}

}

func zAdd(cmd *cobra.Command, args []string) {
	if len(args) <= 0 && args[0] != "web" {
		fmt.Println("参数错误")
	} else {
		// 获取程序绝对路径
		path, err := exec.LookPath(args[0])
		if err != nil {
			log.Println("程序名称：", args[0], "不存在")
			return
		}
		fmt.Println("path：", path)
		abs, err := filepath.Abs(path)
		if err != nil {
			log.Println("error：", err)
			return
		}
		dir := filepath.Dir(abs)
		fmt.Printf("程序路径：%s\n程序名称：%s\n任务端口：%d\ndir：%s\n", abs, name, port, dir)
		if err := db.Add(abs, name, port, dir); err != nil {
			fmt.Println("添加失败")
			return
		} else {
			fmt.Println("添加成功")
		}
	}
}
