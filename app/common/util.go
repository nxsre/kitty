package common

import (
	"strings"
	"net"
)

func GetIPFromUrl(url string) string {

	url = strings.Split(url, "//")[1]
	url = strings.Split(url, "/")[0]
	if strings.Contains(url, ":") {

		return strings.Split(url, ":")[0]
	}

	return url

}


func GetLocalAddr() string {
	conn, err := net.Dial("udp", "localhost:80")
	if err != nil {
		return ""
	}

	defer conn.Close()

	return strings.Split(conn.LocalAddr().String(), ":")[0]
}

