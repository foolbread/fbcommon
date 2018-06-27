/*
author: foolbread
file: lvs/lvs.go
date: 2018/6/27
*/
package lvs

import (
	"io"
	"strings"
	"strconv"
	"bytes"
	"bufio"
	"os/exec"
	"net"
)

const  (
	net_proto_tcp = "TCP"
	net_proto_udp = "UDP"
)

type LVSInfo struct {
	Vip VIPInfo
	Rs []*RSInfo
}

type VIPInfo struct {
	Proto string
	IP string
	Port string
	Strategy string
}

type RSInfo struct {
	IP string
	Port string
	Weight int
	ActiveConn int
	InActConn int
}

func GetAllLvsInfo()([]*LVSInfo,error){
	cmd := exec.Command("ipvsadm","-ln")
	data,err := cmd.Output()
	if err != nil{
		return nil,err
	}

	rd := bytes.NewReader(data)
	brd := bufio.NewReader(rd)
	var ret []*LVSInfo
	var i *LVSInfo
	for {
		li,err := brd.ReadString('\n')
		if err != nil {
			if err == io.EOF{
				break
			}
			return nil,err
		}

		lis := strings.Fields(li)

		if len(lis) > 2 && (lis[0] == net_proto_tcp || lis[0] == net_proto_udp){
			i = new(LVSInfo)
			ret = append(ret,i)

			i.Vip.Proto = lis[0]
			i.Vip.Strategy = lis[2]
			i.Vip.IP,i.Vip.Port,_ = net.SplitHostPort(lis[1])
		}

		if len(lis)>5 && lis[0] == "->"&& i != nil{
			rs := new(RSInfo)
			rs.IP,rs.Port,_= net.SplitHostPort(lis[1])
			rs.Weight,_ = strconv.Atoi(lis[3])
			rs.ActiveConn,_ = strconv.Atoi(lis[4])
			rs.InActConn,_ = strconv.Atoi(lis[5])

			i.Rs = append(i.Rs,rs)
		}
	}

	return ret,nil
}