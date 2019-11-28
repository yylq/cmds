package main

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"reflect"
	"sync"
	"time"
)

var (
	wg      sync.WaitGroup
	baseCmd = &cobra.Command{
		Use:   "base",
		Short: "base test",
	}
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
	reflectCmd = &cobra.Command{
		Use:   "reflect",
		Short: "reflect test",
		Run: func(cmd *cobra.Command, args []string) {
			runreflectCmd()
		},
	}
)

func init() {
	rootCmd.AddCommand(baseCmd)
	baseCmd.AddCommand(chanCmd, contextCmd, reflectCmd)
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
		case <-time.After(10 * time.Second):

			fmt.Printf("10s timer\n")
		}
	}

}
func TestPanicContext(ctx context.Context) {
	str := []string{"111", "2222", "333", "4444", "5555"}
	i := 0
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(5 * time.Second):
			s := str[i]
			fmt.Printf("%s\n", s)
			i += 1
		}
	}
}

type checkerResult struct {
	TimeStamp          string
	SrcAppName         string
	DestAppName        string
	RemoteIp           string
	XForwardedFor      string
	AuthToken          string
	AuthtokenCheckOpen string
	Result             string
}

func (r *checkerResult) String() string {
	return fmt.Sprintf("%s|%s|%s|%s|%s|%s|%s|%s\n",
		r.TimeStamp, r.SrcAppName, r.DestAppName,
		r.RemoteIp, r.XForwardedFor, r.AuthToken,
		r.AuthtokenCheckOpen, r.Result)
}

func runreflectCmd() {
	cr := &CheckerResult{TimeStamp: "1111"}
	getType := reflect.TypeOf(cr)
	fmt.Println("get Type is :", getType.Name())

	getValue := reflect.ValueOf(cr)
	fmt.Println("get all Fields is:", getValue)

	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		fmt.Println("Fields is:", field)
	}
}
