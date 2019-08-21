package main

import (
	"fmt"
	log "github.com/cihub/seelog"
	"github.com/spf13/cobra"
)

var (
	msg       string
	seelogCmd = &cobra.Command{
		Use:   "seelog",
		Short: "seelog msg",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
				logger, err := log.LoggerFromConfigAsFile("conf/seelog.xml")
				if err != nil {
					return err
				}
				log.ReplaceLogger(logger)
				return nil
		},
	}
	logmsgCmd = &cobra.Command{
		Use:   "logmsg",
		Short: "set seelog log level",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("msg:%s\n",msg)
			log.Trace( msg)
			log.Debug( msg)
			log.Info( msg)
			log.Warn( msg)
			log.Error( msg)
			log.Critical( msg)
		},
	}
)

func init() {
	rootCmd.AddCommand(seelogCmd)
	seelogCmd.AddCommand(logmsgCmd)
	logmsgCmd.PersistentFlags().StringVarP(&msg, "msg", "m", "", "log message")
}
