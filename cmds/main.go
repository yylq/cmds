package main

import (
	"flag"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)


var(
	rootCmd = &cobra.Command{
		Use:   "cmds",
		Short: "cmds test",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println(args)
			return nil
		},
	}

)


func main() {

	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}


/*
	formatter := &log.TextFormatter{
		// 不需要彩色日志
		DisableColors:   true,
		// 定义时间戳格式
		TimestampFormat: "2006-01-02 15:04:05",
	}
	os.Stderr
	log.SetOutput()
	log.SetFormatter(formatter)
	log.Printf("hello world")

 */
}

func init() {
	rootCmd.SetArgs(os.Args[1:])
	rootCmd.PersistentFlags().AddGoFlagSet(flag.CommandLine)
}