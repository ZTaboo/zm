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
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
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
	} else {
		log.Println("数据文件已存在")
	}
}
