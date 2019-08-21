package main
import(
	"fmt"
	"github.com/cihub/seelog"
	"github.com/spf13/cobra"
)
var(
	msg string
	seelogCmd = &cobra.Command{
		Use:   "seelog",
		Short: "seelog msg",
		Run: func(cmd *cobra.Command, args []string) {
			seelog.
			seelog.Error("msg:%v", msg)
		},
	}
)
func init() {
	fmt.Printf("seelog init")
	rootCmd.AddCommand(seelogCmd)
	seelogCmd.PersistentFlags().StringVarP(&msg, "msg", "m","", "log message")
}