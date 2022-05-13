package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
	"time"
)

func getCurrentTime() (time.Time, error) {
	t, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

func main() {
	time, err := getCurrentTime()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println("Current Time is: ", time)
}
