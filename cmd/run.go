package cmd

import (
	"fmt"
	"github.com/ZTaboo/ZM/db"
	"github.com/ZTaboo/ZM/router"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "运行任务(未完成)",
	Long:  `运行任务；内置默认任务：web(未完成).`,
	Run:   Run,
}
var runBackCmd = &cobra.Command{
	Use:   "backend",
	Short: "后台运行任务(未完成)",
	Long:  `后台运行运行任务；内置默认任务：web(未完成).`,
	Run:   BackRun,
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.AddCommand(runBackCmd)
}

// BackRun 后台运行
func BackRun(cmd *cobra.Command, args []string) {
	if len(args) <= 0 {
		log.Println("请输入要运行的程序名称，例如：zm run web")
	} else {
		switch args[0] {
		case "web":
			// 判断自身是否后台运行；如果没有则运行ZM Web
			listen, err := net.Listen("tcp4", ":1999")
			if err != nil {
				log.Println("程序已运行：error:", err)
				return
			}
			defer listen.Close()
			c := exec.Command(os.Args[0], "run", "web")
			c.Stdout = nil
			c.Stderr = nil
			err = c.Start()
			if err != nil {
				log.Println("error:", err)
				return
			}
			log.Println("启动成功\npid:", c.Process.Pid, "port:1999")
			break
		default:
			// 查询任务是否存在
			// 查询任务是否存在
			if task, err := db.TaskExist(args[0]); err != nil {
				fmt.Println("任务名称：", args[0], "不存在")
			} else {
				fmt.Println("任务名称：", task.Name)
				fmt.Println("任务路径：", task.Path)
				c := exec.Command(task.Path)
				c.Dir = task.ExecutableFiles
				c.Stdout = nil
				c.Stderr = nil
				err = c.Start()
				if err != nil {
					log.Println("run error:", err)
					return
				}
				db.UpdatePid(task.Name, c.Process.Pid)
				fmt.Println("启动成功\npid:", c.Process.Pid, "port:", task.Port)
			}
		}
	}

}

// Run 前台运行
func Run(cmd *cobra.Command, args []string) {
	if len(args) <= 0 {
		log.Println("请输入要运行的程序名称，例如：zm run web")
	} else {
		switch args[0] {
		case "web":
			app := fiber.New()
			router.Router(app)
			log.Fatal(app.Listen(":1999"))
			break
		default:
			//查询任务是否存在
			if task, err := db.TaskExist(args[0]); err != nil {
				fmt.Println("任务名称：", args[0], "不存在")
			} else {
				fmt.Println("任务名称：", task.Name)
				fmt.Println("任务路径：", task.Path)
				c := exec.Command(task.Path)
				c.Dir = task.ExecutableFiles
				stdoutPipe, err := c.StdoutPipe()
				if err != nil {
					log.Println("error:", err)
					return
				}
				stderr, err := c.StderrPipe()
				if err != nil {
					log.Println("error:", err)
					return
				}
				err = c.Start()
				if err != nil {
					log.Println("run error:", err)
					return
				}
				if _, err := io.Copy(os.Stdout, stdoutPipe); err != nil {
					log.Println("error:", err)
				}
				if _, err := io.Copy(os.Stderr, stderr); err != nil {
					log.Println("error:", err)
				}
			}

			break
		}
	}
}
