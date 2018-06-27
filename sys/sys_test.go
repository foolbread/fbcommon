/*
author: foolbread
file: sys/
date: 2018/6/27
*/
package sys

import (
	"testing"
)

func Test_GetProcessByName(t *testing.T) {
	pids,err := GetProcessByName("asdasd")
	if err != nil{
		t.Fatal(err)
	}

	t.Log("pids:",pids)
}
