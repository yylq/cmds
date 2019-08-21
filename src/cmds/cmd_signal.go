package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
)

var(

	sigCmd = &cobra.Command{
		Use:   "sig",
		Short: "sig test",
		Run: func(cmd *cobra.Command, args []string) {
			c := make(chan os.Signal) //监听所有信号
			signal.Notify(c)          //阻塞直到有信号传入
			fmt.Println("启动")
			s := <-c
			fmt.Println("退出信号", s)
		},
	}
)
func init() {
	rootCmd.AddCommand(sigCmd)
}