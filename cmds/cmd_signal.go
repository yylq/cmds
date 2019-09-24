package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var(

	sigCmd = &cobra.Command{
		Use:   "sig",
		Short: "sig test",
		Run: func(cmd *cobra.Command, args []string) {
			go signalListen()
			// main loop
			for {
				time.Sleep(30 * time.Second)
				fmt.Println("main loop.")
			}
		},
	}
)
func init() {
	rootCmd.AddCommand(sigCmd)
}

func signalListen() {
	// init os.signal channel
	c := make(chan os.Signal)
	// define catch signal
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM,syscall.SIGTERM)
	for {
		// wait channel
		sig := <-c
		// when receive signal,then notify channel,and print the follow info.
		fmt.Println("receive signal:", sig)
	}
}
