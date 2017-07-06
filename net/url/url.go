/*
author: foolbread
file: url/url.go
date: 2017/7/6
*/
package url

import (
	"net/url"
)

func RawUrlEncode(s string)string{
	ur := url.URL{Path:s}

	return ur.String()
}