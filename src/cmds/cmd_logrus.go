package main

import (
	"github.com/spf13/cobra"
	log "github.com/sirupsen/logrus"
)

var(

	logrusCmd = &cobra.Command{
		Use:   "logrus",
		Short: "logrus test",
		Run: func(cmd *cobra.Command, args []string) {
			log.WithFields(log.Fields{
				"animal": "walrus",
			}).Info("A walrus appears")
		},
	}
)
func init() {
	rootCmd.AddCommand(logrusCmd)
}