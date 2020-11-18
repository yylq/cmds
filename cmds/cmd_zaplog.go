package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	)

var(
	level int8
	zapCmd = &cobra.Command{
		Use:   "zap",
		Short: "zap test",
		Run:runZapCmd,
	}
)
func init() {
	rootCmd.AddCommand(zapCmd)
	zapCmd.PersistentFlags().Int8VarP(&level, "level","l", 0, "level" )
}
func runZapCmd(cmd *cobra.Command, args []string){
	fmt.Println("aaa")
	productionConfig := zap.NewProductionConfig()
	logger, err := productionConfig.Build();
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("zap.DebugLevel:%v\n",zap.DebugLevel)
	logger.Debug("this is debug")
	logger.Info("this is Info")
	logger.Warn("this is Warn")
	logger.Error("this is Error")
	logger.Panic("this is Panic")
	logger.DPanic("this is DPanic")
	logger.Fatal("this is Fatal")
	fmt.Println("---------------------------------------")
	logger= logger.With(zap.String("prod", "aaa"), zap.String("servcie", "tids"))
	logger.Debug("this is debug")
	logger.Info("this is Info")
	logger.Warn("this is Warn")
	logger.Error("this is Error")
	logger.Fatal("this is Fatal")
}