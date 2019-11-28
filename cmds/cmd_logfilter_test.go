package main

import (
	"os"
	"testing"
)

func TestFilter(t *testing.T)  {
	inFile:="../testdata/t.log"
	outFile:="../testdata/293.log"
	tag:= "[293]"
	t.Log(os.Getwd())
	err := Filter(inFile,outFile,tag)
	if err != nil {
		t.Fatal(err)
	}
}