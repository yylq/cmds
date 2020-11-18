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
	}
	contextPanicCmd = &cobra.Command{
		Use:   "panic",
		Short: "context panic",
		Run: func(cmd *cobra.Command, args []string) {
			runcontextPanicCmd()
		},
	}
	contextValueCmd = &cobra.Command{
		Use:   "value",
		Short: "context test",
		Run: func(cmd *cobra.Command, args []string) {
			runcontextValueCmd()
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
	contextCmd.AddCommand(contextPanicCmd, contextValueCmd)
}
func waitForCompletion(ctx context.Context, fn func(context.Context)) {
	wg.Add(1)
	fn(ctx)
	wg.Done()
}

func runcontextPanicCmd() {
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
type Key struct {
	Name string
}
type Value struct {
	Name string
}
var(
	key = Key{}
	key1 = Key{Name:"jdcloud"}
)


func injectkey(ctx context.Context) context.Context{
	return context.WithValue(ctx,key, &Value{"key is empty"})
}

func injeckey1tcontext(ctx context.Context) context.Context {
	return context.WithValue(ctx,key1, &Value{"key is jdcloud"})
}

func runcontextValueCmd() {
	type favContextKey string

	f := func(ctx context.Context, k Key) {
		if v,ok := ctx.Value(k).(*Value); ok {
			fmt.Printf("found value:%v\n", v)
			return
		}
		fmt.Printf("key not found:%v\n", k)
	}
	ctx := context.Background()
	ctx1 := injectkey(ctx)
	ctx2 := injeckey1tcontext(ctx1)
	fmt.Printf("found in ctx :%v\n", ctx)
	f(ctx, key)
	f(ctx, key1)
	fmt.Printf("found in ctx1 :%v\n", ctx1)
	f(ctx1, key)
	f(ctx1, key1)
	fmt.Printf("found in ctx2 :%v\n", ctx2)
	f(ctx2, key)
	f(ctx2, key1)

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
	cr := &checkerResult{TimeStamp: "1111"}
	cr2 := &checkerResult{TimeStamp: "2222"}
	cr3 := &checkerResult{TimeStamp: "1111"}
	getType := reflect.TypeOf(cr)
	fmt.Println("get Type is :", getType.Name())

	getValue := reflect.ValueOf(cr)
	fmt.Println("get all Fields is:", getValue)

/*	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		fmt.Println("Fields is:", field)
	}
*/
	fmt.Println("DeepEqual:",reflect.DeepEqual(cr,cr3))
	fmt.Println("DeepEqual:",reflect.DeepEqual(cr,cr2))
}
