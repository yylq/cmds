package main

import (
	"bytes"
	"encoding/hex"
	"github.com/spf13/cobra"
	"net"
	"time"

	"encoding/binary"
	"fmt"
)

const (
	TRACKER_PROTO_CMD_SERVER_LIST_ALL_GROUPS byte = 91
)

type ProtoHeader struct {
	PkgLen int64
	Cmd    byte
	Status byte
}
type TrackerGroupStat struct {
	Group_name               [17]byte
	Sz_total_mb              int64 //total disk storage in MB
	Sz_free_mb               int64 //free disk storage in MB
	Sz_trunk_free_mb         int64 //trunk free space in MB
	Sz_count                 int64 //server count
	Sz_storage_port          int64
	Sz_storage_http_port     int64
	Sz_active_count          int64 //active server count
	Sz_current_write_server  int64
	Sz_store_path_count      int64
	Sz_subdir_count_per_path int64
	Sz_current_trunk_file_id int64
}

type ProtoTrackerGroupStats struct {
	Header ProtoHeader
	GroupState []TrackerGroupStat
}
func (p *ProtoTrackerGroupStats)Read(data []byte)(int64,  error) {
	buff := bytes.NewBuffer(data)
	blen := buff.Len()
	fmt.Printf("blen:%d\n",blen)
	headbuff := bytes.NewBuffer(buff.Next(binary.Size(p.Header)))
	binary.Read(headbuff, binary.BigEndian, &p.Header)
	blen = buff.Len()
	fmt.Printf("blen:%d\n",blen)
	if p.Header.PkgLen >0 && (p.Header.PkgLen % (int64(binary.Size(&TrackerGroupStat{}))) != 0) {
		return 0, fmt.Errorf("invalid pkg")
	}

	sts := make([]TrackerGroupStat,0)
	for {
		var ts TrackerGroupStat
		fmt.Printf("p.Header.PkgLen:%d\n",p.Header.PkgLen)
		sz := binary.Size(ts)
		fmt.Printf("sz:%d\n",sz)
		bodybuff := bytes.NewBuffer(buff.Next(sz))
		binary.Read(bodybuff, binary.BigEndian, &ts)
		fmt.Printf("ts:%v\n",ts)
		sts=append(sts, ts)
		blen = buff.Len()
		fmt.Printf("blen:%d\n",blen)
		p.Header.PkgLen -= int64(sz)
		fmt.Printf("p.Header.PkgLen:%d\n",p.Header.PkgLen)
		if  p.Header.PkgLen<= 0 {
			break
		}
	}
	p.GroupState = sts
	return int64(0), nil
}




var (
	fastCmd = &cobra.Command{
		Use:   "fast",
		Short: "fast",
	}
	fastConfigCmd = &cobra.Command{
		Use:   "list",
		Short: "list",
		Run: func(cmd *cobra.Command, args []string) {
			var p1 ProtoHeader
			p1.Cmd = TRACKER_PROTO_CMD_SERVER_LIST_ALL_GROUPS
			p1.Status = 2
			fmt.Printf("p1:%v %v \n", p1, binary.Size(&p1))
			var ts TrackerGroupStat
			fmt.Printf("ts:%v %v \n", ts, binary.Size(&ts))
			data := encode(&p1)

			host := "116.196.117.15:22122"
			conn, err := net.DialTimeout("tcp", host, 200*time.Millisecond)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer conn.Close()
			n, err := conn.Write(data)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("sent data len:%d\n", n)
			buf := make([]byte, 1024)
			n, err = conn.Read(buf)
			if err != nil {
				fmt.Println(err)
				return
			}
			if n == 0 {
				fmt.Println("conn is colse")
				return
			}
			var p2 ProtoTrackerGroupStats

			_, err = p2.Read(buf[:n])
			if err != nil {
				fmt.Println("conn is colse")
				return
			}
			fmt.Printf("p2:%v\n", p2)
			for _, st := range p2.GroupState {
				fmt.Printf("groupname:%s\n",st.Group_name)
			}


		},
	}
)

func init() {
	rootCmd.AddCommand(fastCmd)
	fastCmd.AddCommand(fastConfigCmd)
}

func encode(p *ProtoHeader) []byte {
	var buff bytes.Buffer

	binary.Write(&buff, binary.BigEndian, p)
	data := buff.Bytes()
	dst := make([]byte, hex.EncodedLen(len(data)))
	hex.Encode(dst, data)
	fmt.Printf("%s\n", dst)
	return data
}

func decode(data []byte, p *ProtoTrackerGroupStats) {

	dst := make([]byte, hex.EncodedLen(len(data)))
	hex.Encode(dst, data)
	fmt.Printf("%s\n", dst)
	buff := bytes.NewBuffer(data)
	binary.Read(buff, binary.BigEndian, p)
}
