package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

var(
	timestr string
	timeCmd = &cobra.Command{
		Use:   "time",
		Short: "set seelog log level",
		Run: func(cmd *cobra.Command, args []string) {
			dateStr := "2016-07-14 14:24:51"
			timestamp1, _ := time.Parse("2006-01-02 15:04:05", dateStr)
			timestamp2, _ := time.ParseInLocation("2006-01-02 15:04:05", dateStr, time.Local)
			fmt.Println(timestamp1)
			fmt.Println(timestamp2)               //2016-07-14 14:24:51 +0000 UTC 2016-07-14 14:24:51 +0800 CST
			fmt.Println(timestamp1.Unix(), timestamp2.Unix()) //1468506291 1468477491

			now := time.Now()
			year, mon, day := now.UTC().Date()
			hour, min, sec := now.UTC().Clock()
			zone, _ := now.UTC().Zone()
			fmt.Printf("UTC 时间是 %d-%d-%d %02d:%02d:%02d %s\n",
				year, mon, day, hour, min, sec, zone) // UTC 时间是 2016-7-14 07:06:46 UTC

			year, mon, day = now.Date()
			hour, min, sec = now.Clock()
			zone, _ = now.Zone()
			fmt.Printf("本地时间是 %d-%d-%d %02d:%02d:%02d %s\n",
				year, mon, day, hour, min, sec, zone) // 本地时间是 2016-7-14 15:06:46 CST
			fmt.Println(now.Format("2006-01-02 15:04:05"))
			dateStr = fmt.Sprintf("%d-%02d-%02d 23:59:59",year, mon, day)
			fmt.Println(dateStr)
			timezero, err := time.ParseInLocation("2006-01-02 15:04:05", dateStr, time.Local)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(timezero.Unix(),now.Unix(),timezero.Unix() - now.Unix())
			dateStr = fmt.Sprintf("%d-%02d-%02d 23:59:59",year, mon, day+1)
			nexttimezero, err := time.ParseInLocation("2006-01-02 15:04:05", dateStr, time.Local)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(nexttimezero.Unix(),timezero.Unix(),nexttimezero.Unix() - timezero.Unix())

		},
	}
)

func init() {
	rootCmd.AddCommand(timeCmd)

	timeCmd.PersistentFlags().StringVarP(&timestr, "time", "t", "", "time string")
}