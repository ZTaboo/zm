/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "初始化基础信息",
	Long:  `初始化基础信息.`,
	Run:   zInit,
}
var dbFile string

func init() {
	rootCmd.AddCommand(initCmd)
	//	 获取user目录
	userDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		return
	}
	//	 在user目录下创建zm.db文件
	dbFile = userDir + "/zm.db"
}

func zInit(cmd *cobra.Command, args []string) {
	//	 获取user目录
	userDir, err := os.UserHomeDir()
	if err != nil {
		log.Println(err)
		return
	}
	//	 在user目录下创建zm.db文件
	dbFile := userDir + "/zm.db"
	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		dbFile, err := os.Create(dbFile)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer dbFile.Close()
		log.Println("创建数据文件成功")
	}
	fmt.Println("初始化完成")
}
