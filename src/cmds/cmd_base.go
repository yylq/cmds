package main

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"sync"
	"time"
)

var (
	wg      sync.WaitGroup
	chanCmd = &cobra.Command{
		Use:   "chan",
		Short: "chan test",
		Run: func(cmd *cobra.Command, args []string) {
			log.SetFlags(log.Ldate)

			log.Println("sssssssss", 10)
		},
	}
	contextCmd = &cobra.Command{
		Use:   "context",
		Short: "context test",
		Run: func(cmd *cobra.Command, args []string) {
			runcontextCmd()
		},
	}
)

func init() {
	rootCmd.AddCommand(chanCmd,contextCmd)
}
func waitForCompletion(ctx context.Context, fn func(context.Context)) {
	wg.Add(1)
	fn(ctx)
	wg.Done()
}
func runcontextCmd() {
	ctx, _ := context.WithCancel(context.Background())
	go waitForCompletion(ctx, TestPanicContext)
	for {
		select {
		case <-time.After(10*time.Second):

			fmt.Printf("10s timer\n",)
		}
	}

}
func TestPanicContext(ctx context.Context) {
	str:=[]string{"111","2222","333","4444","5555"}
	i :=0;
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(5*time.Second):
			s := str[i]
			fmt.Printf("%s\n",s)
			i+=1
		}
	}
}


