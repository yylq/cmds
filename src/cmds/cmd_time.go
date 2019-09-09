package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"strings"
	"time"
)

var (
	timestr string
	timeCmd = &cobra.Command{
		Use:   "time",
		Short: "set seelog log level",
		Run: func(cmd *cobra.Command, args []string) {
			dateStr := "2016-07-14 14:24:51"
			timestamp1, _ := time.Parse("2006-01-02 15:04:05", dateStr)
			timestamp2, _ := time.ParseInLocation("2006-01-02 15:04:05", dateStr, time.Local)
			fmt.Println(timestamp1)
			fmt.Println(timestamp2)                           //2016-07-14 14:24:51 +0000 UTC 2016-07-14 14:24:51 +0800 CST
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
			dateStr = fmt.Sprintf("%d-%02d-%02d 23:59:59", year, mon, day)
			fmt.Println(dateStr)
			timezero, err := time.ParseInLocation("2006-01-02 15:04:05", dateStr, time.Local)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(timezero.Unix(), now.Unix(), timezero.Unix()-now.Unix())
			dateStr = fmt.Sprintf("%d-%02d-%02d 23:59:59", year, mon, day+1)
			nexttimezero, err := time.ParseInLocation("2006-01-02 15:04:05", dateStr, time.Local)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(nexttimezero.Unix(), timezero.Unix(), nexttimezero.Unix()-timezero.Unix())
			Url := "http://127.0.0.1:8000"
			fmt.Printf("%s\n", strings.Replace(Url, "http://", "", -1))
			fmt.Printf("%s\n",formatDate(now))
			fmt.Printf("%s\n",now.UTC().Format("2006-01-02T15:04:05"))
			var sa []string
			fmt.Printf("%d\n",len(sa))
		},
	}
	logCmd = &cobra.Command{
		Use:   "log",
		Short: "set seelog log level",
		Run: func(cmd *cobra.Command, args []string) {
			log.SetFlags(log.Ldate)

			log.Println("sssssssss",10)
		},
	}
	chanCmd = &cobra.Command{
		Use:   "log",
		Short: "set seelog log level",
		Run: func(cmd *cobra.Command, args []string) {
			log.SetFlags(log.Ldate)

			log.Println("sssssssss",10)
		},
	}
)

func init() {
	rootCmd.AddCommand(timeCmd,logCmd, chanCmd)

	timeCmd.PersistentFlags().StringVarP(&timestr, "time", "t", "", "time string")
}

func formatDate(t time.Time) string {
	t = t.UTC()
	year, month, day := t.Date()
	hour, minute, second := t.Clock()

	buf := make([]byte, 19)

	buf[0] = byte((year/1000)%10) + '0'
	buf[1] = byte((year/100)%10) + '0'
	buf[2] = byte((year/10)%10) + '0'
	buf[3] = byte(year%10) + '0'
	buf[4] = '-'
	buf[5] = byte((month)/10) + '0'
	buf[6] = byte((month)%10) + '0'
	buf[7] = '-'
	buf[8] = byte((day)/10) + '0'
	buf[9] = byte((day)%10) + '0'
	buf[10] = 'T'
	buf[11] = byte((hour)/10) + '0'
	buf[12] = byte((hour)%10) + '0'
	buf[13] = ':'
	buf[14] = byte((minute)/10) + '0'
	buf[15] = byte((minute)%10) + '0'
	buf[16] = ':'
	buf[17] = byte((second)/10) + '0'
	buf[18] = byte((second)%10) + '0'

	return string(buf)
}