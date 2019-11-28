package main

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"os"
	"strings"
)

var (
	logFile string
	threadNum string
	outFile string
	logFilterCmd = &cobra.Command{
		Use:   "logfilter",
		Short: "logfilter thread log",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return runlogFilterCmd()
		},
	}
	)
func init() {
	rootCmd.AddCommand(logFilterCmd)
	logFilterCmd.PersistentFlags().StringVarP(&logFile, "log", "l", "", "log file name")
	logFilterCmd.PersistentFlags().StringVarP(&threadNum, "thread", "t", "", "thread id")
	logFilterCmd.PersistentFlags().StringVarP(&outFile, "out", "o", "", "out file name")
}

func runlogFilterCmd() error {
	wd,_ := os.Getwd()
	if outFile =="" {
		outFile = fmt.Sprintf("%s/%s.log",wd, threadNum)
	}
	tag := fmt.Sprintf("[%s]",threadNum)
	return Filter(logFile,outFile,tag)

}
func Filter(inFile ,outFile, tag string) error {
	in,err := os.Open(inFile)
	if err != nil {
		return err
	}
	defer in.Close()
	out ,err := os.OpenFile(outFile,os.O_CREATE|os.O_WRONLY,0644);
	if err != nil {
		return err
	}
	defer out.Close()
	wr := bufio.NewWriter(out)
	buf := bufio.NewReader(in)
	i:= 0
	isThreadLog := true
	for {
		i++
		line, err := buf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		if line[0] !='[' {
			if  isThreadLog {
				wr.WriteString(fmt.Sprintf("%d ",i)+line+"\n")
				continue
			}

		}

		ind := strings.Index(line, tag)
		if ind < 0 {
			isThreadLog = false
			continue
		}
		wr.WriteString(fmt.Sprintf("%d ",i)+line+"\n")
	}
	return nil
}