package main

import "command/internal"

func main() {
	tv := &internal.Tv{}
	onCommand := &internal.OnCommand{
		tv,
	}
	offCommand := &internal.OffCommand{
		tv,
	}
	onButton := &internal.Button{
		onCommand,
	}
	onButton.Press()
	offButton := &internal.Button{
		offCommand,
	}
	offButton.Press()
}
