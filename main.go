package main

import (
	"fmt"
	"github.com/jbrodriguez/mlog"
)

func main() {

	mlog.Start(mlog.LevelTrace, "log.txt")
	mlog.Trace("-------------------------------------------------")
	mlog.Trace("             Start program                       ")
	mlog.Trace("-------------------------------------------------")

	chWeb := CreateWeb()
	mlog.Trace("chWeb", chWeb)

	// временное решение. Нажатие Enter закрывает программу
	var response string

	fmt.Println("Press Enter")
	_, _ = fmt.Scanln(&response)

	// Stop program
	mlog.Trace("Stop program\n\n")
}
