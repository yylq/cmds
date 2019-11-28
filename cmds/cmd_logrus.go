package main

import (
	"encoding/json"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
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
func (c *CheckerResult) ToField() logrus.Fields {
		return logrus.Fields{ "TimeStamp": c.TimeStamp,
		"SrcAppName":c.SrcAppName, "DestAppName":c.DestAppName,
		"RemoteIp":c.RemoteIp, "XForwardedFor":c.XForwardedFor,
		"AuthToken":c.AuthToken, "AuthtokenCheckOpen":c.AuthtokenCheckOpen, "Result":c.Result}

}

var (
	logrusCmd = &cobra.Command{
		Use:   "logrus",
		Short: "logrus test",
	}
	logrusDefaultCmd = &cobra.Command{
		Use:   "default",
		Short: "default test",
		Run: func(cmd *cobra.Command, args []string) {
			logrus.Trace("this a Trace msg")
			logrus.Debug("this a debug msg")
			logrus.Info("this a info msg")
			logrus.Error("this a Error msg")
			logrus.Warn("this a Warn msg")
			logrus.Fatal("this a Fatal msg")
		},
	}
	logrusFormatCmd = &cobra.Command{
		Use:   "format",
		Short: "format test",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	logrusFileCmd = &cobra.Command{
		Use:   "file",
		Short: "file test",
		Run: func(cmd *cobra.Command, args []string) {
			//InitLog()
			//ch := &CheckerResult{TimeStamp: time.Now().UTC().Format("2006-01-02T15:04:05")}
			//log.WithFields(ch.ToField()).Infoln()
			//logger.WithFields(ch.ToField()).Infoln("")
			//entry := logger.WithFields(log.Fields{"request_id": "aaaa", "user_ip": "192.168.11.1"})

			//ch := &CheckerResult{}

			//entry.Infoln(ch)
			/*
				logger.WithFields(log.Fields{
					"animal": "walrus",
				}).Log(log.InfoLevel,"")

				logger.Log(log.InfoLevel,"")
			*/
			//entry.Infoln(ch)
			//logger.Info(ch)
		},
	}
)

func init() {
	rootCmd.AddCommand(logrusCmd)
	logrusCmd.AddCommand(logrusDefaultCmd,logrusFormatCmd,logrusFileCmd)
}
func InitLog() {
	logrus.SetLevel(logrus.InfoLevel)
	hook := newLfsHook("checkresult")
	logrus.AddHook(hook)
}
func newLfsHook(logName string) logrus.Hook {

	writer, err := rotatelogs.New(
		logName+"%Y-%m-%d",
		rotatelogs.WithLinkName(logName),      // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(7*24*time.Hour),     // 文件最大保存时间
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
	)

	if err != nil {
		logrus.Errorf("config local file system for logger error: %v", err)
	}
    /*
	formatter := &log.JSONFormatter{DisableTimestamp: true,
		FieldMap: log.FieldMap{
			log.FieldKeyLevel: "lev",
			log.FieldKeyMsg:   "check_result",
		},
	}*/

	lfsHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.InfoLevel:  writer,
	}, new(JdmeshFormatter))

	return lfsHook
}
