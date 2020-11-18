package main

/*
#include <stdio.h>
#include <stdlib.h>
void myhello(char* s) {
  printf("Hello C: %s\n", s);
}
*/
import "C"
import (
	"github.com/spf13/cobra"
	"unsafe"
)

var(
	cgoCmd = &cobra.Command{
		Use:   "cgo",
		Short: "cgo",
	}
	cgoHelloCmd = &cobra.Command{
		Use:   "hello",
		Short: "hello",
		Run: func(cmd *cobra.Command, args []string) {
			str:= C.CString("hello")
			C.myhello(str)
			C.free(unsafe.Pointer(str))
		},
	}
)
func init() {
	rootCmd.AddCommand(cgoCmd)
	cgoCmd.AddCommand(cgoHelloCmd)
}