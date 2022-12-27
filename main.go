package main

import (
	"fmt"
	//"os"
	"ntp"
)

func main() {
	//var ntptime , err string
	ntptime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	//ntptime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err == nil {
		fmt.Println(ntptime)
		//fmt.Println(os.Stderr, ntptime)

	} else {
		fmt.Println(err)
	}
}
