/*
author: foolbread
file: sys/sys.go
date: 2018/6/27
*/
package sys

import (
	"bytes"
	"bufio"
	"os/exec"
	"strconv"
	"strings"
)

func GetProcessByName(name string)([]int,error){
	cmd := exec.Command("pgrep",name)
	data,err := cmd.Output()
	if err != nil{
		return nil,err
	}

	rd := bytes.NewReader(data)
	brd := bufio.NewReader(rd)
	var pids []int
	var pid int
	for{
		pidstr,err := brd.ReadString('\n')
		if err != nil{
			break
		}
		pid,err = strconv.Atoi(strings.TrimRight(pidstr,"\n"))
		if err != nil{
			return pids,err
		}

		pids = append(pids,pid)
	}

	return pids,nil
}
