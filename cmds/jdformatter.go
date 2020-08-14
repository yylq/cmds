package main

import (
	"bytes"
	"fmt"
	log "github.com/sirupsen/logrus"
)
var(
	atrrs =[]string{"TimeStamp","SrcAppName","DestAppName","RemoteIp","XForwardedFor","AuthToken","AuthtokenCheckOpen","Result"}
)
type JdmeshFormatter struct {

}

func (f *JdmeshFormatter)Format(entry *log.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	for i,attr := range atrrs {
		value,ok := entry.Data[attr]
		if !ok {
			f.appendKeyValue(b, attr, "")
			b.WriteByte('|')
		}
		f.appendKeyValue(b, attr, value)
		if i==0 {
			continue
		}
		b.WriteByte('|')
	}

	b.WriteByte('\n')
	return b.Bytes(), nil
}
func (f *JdmeshFormatter) appendKeyValue(b *bytes.Buffer, key string, value interface{}) {

	f.appendValue(b, value)
}
func (f *JdmeshFormatter) appendValue(b *bytes.Buffer, value interface{}) {
	stringVal, ok := value.(string)
	if !ok {
		stringVal = fmt.Sprint(value)
	}
	b.WriteString(stringVal)
}
