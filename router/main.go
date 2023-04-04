package router

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
)

// Run 前台运行
func Run(cmd *cobra.Command, args []string) {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello world")
	})
	log.Fatal(app.Listen(":8080"))
}

// BackRun 后台运行
func BackRun(cmd *cobra.Command, args []string) {
	c := exec.Command(os.Args[0], "run")
	c.Stdout = nil
	c.Stderr = nil
	err := c.Start()
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("pid:", c.Process.Pid)

}
