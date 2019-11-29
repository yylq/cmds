package main

import (
	"encoding/json"
	"fmt"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"

	"time"
)

type CheckerResult struct {
	TimeStamp          string `json:"TimeStamp"`
	SrcAppName         string `json:"SrcAppName"`
	DestAppName        string
	RemoteIp           string
	XForwardedFor      string
	AuthToken          string
	AuthtokenCheckOpen string
	Result             string
}

func (c *CheckerResult) String() string {
	bytes, _ := json.Marshal(c)
	return string(bytes)
}
func (c *CheckerResult) ToField() log.Fields {
	return log.Fields{"TimeStamp": c.TimeStamp,
		"SrcAppName": c.SrcAppName, "DestAppName": c.DestAppName,
		"RemoteIp": c.RemoteIp, "XForwardedFor": c.XForwardedFor,
		"AuthToken": c.AuthToken, "AuthtokenCheckOpen": c.AuthtokenCheckOpen, "Result": c.Result}

}

var (
	loglevel string
	fmtname string
	output string
	logrusCmd = &cobra.Command{
		Use:   "logrus",
		Short: "logrus test",
		PersistentPreRun :func(cmd *cobra.Command, args []string) {
			level ,err := log.ParseLevel(loglevel);
			if err == nil {
				log.SetLevel(level)
			}
		},
	}
	logrusDefaultCmd = &cobra.Command{
		Use:   "default",
		Short: "default test",
		Run: func(cmd *cobra.Command, args []string) {
			formatter := &log.TextFormatter{
				// 不需要彩色日志
				DisableColors:   true,
				// 定义时间戳格式
				TimestampFormat: "2006-01-02 15:04:05",
			}
			log.SetFormatter(formatter)
			log.Printf("hello world")
			testlog()
		},
	}
	logrusFormatCmd = &cobra.Command {
		Use:   "format",
		Short: "format test",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			setformat()
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			log.Printf("hello world")
			testlog()
		},
	}
	logrusOutputCmd = &cobra.Command{
		Use:   "output",
		Short: "output test",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			setformat()
			return setoutput()
		},
		Run: func(cmd *cobra.Command, args []string) {
			log.Printf("hello world")
			testlog()
		},
	}
)

func init() {
	rootCmd.AddCommand(logrusCmd)

	logrusCmd.AddCommand(logrusDefaultCmd, logrusFormatCmd, logrusOutputCmd)
	logrusCmd.PersistentFlags().StringVarP(&loglevel, "log_level", "l", "info", "log level")
	logrusCmd.PersistentFlags().StringVarP(&fmtname, "format", "f", "text", "log format")
	logrusOutputCmd.PersistentFlags().StringVarP(&output, "output", "o", "stdout", "output")

}
func InitLog() {
	log.SetLevel(log.InfoLevel)
	hook := newLfsHook("checkresult")
	log.AddHook(hook)
}
func newLfsHook(logName string) log.Hook {

	writer, err := rotatelogs.New(
		logName+"%Y-%m-%d",
		rotatelogs.WithLinkName(logName),          // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(7*24*time.Hour),     // 文件最大保存时间
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
	)

	if err != nil {
		log.Errorf("config local file system for logger error: %v", err)
	}
	/*
		formatter := &log.JSONFormatter{DisableTimestamp: true,
			FieldMap: log.FieldMap{
				log.FieldKeyLevel: "lev",
				log.FieldKeyMsg:   "check_result",
			},
		}*/

	lfsHook := lfshook.NewHook(lfshook.WriterMap{
		log.InfoLevel: writer,
	}, new(JdmeshFormatter))

	return lfsHook
}
func testlog() {
	log.Trace("this a Trace msg")
	log.Debug("this a debug msg")
	log.Info("this a info msg")
	log.Error("this a Error msg")
	log.Warn("this a Warn msg")
	log.Fatal("this a Fatal msg")
}
func setformat(){
	var formatter log.Formatter
	switch fmtname {
	case "json":
		fmt.Println("fname:",fmtname)
		formatter = &log.JSONFormatter{}

		break;
	default:
		formatter = &log.TextFormatter{
			// 不需要彩色日志
			DisableColors:   true,
			// 定义时间戳格式
			TimestampFormat: "2006-01-02 15:04:05",
		}
	}
	log.SetFormatter(formatter)
}
func setoutput() error{
	switch output {
	case "stdout":
		log.SetOutput(os.Stdout)
		break
	case "stderr":
		log.SetOutput(os.Stderr)
		break
	default:
		writer, err := rotatelogs.New(
			output+"%Y-%m-%d",
			rotatelogs.WithLinkName(output),      // 生成软链，指向最新日志文件
			rotatelogs.WithMaxAge(7*24*time.Hour),     // 文件最大保存时间
			rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
		)

		if err != nil {
			return nil
		}
		log.SetOutput(writer)
	}
	return nil
}