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
	var test1 string
	test1 = "aaa"
	fmt.Println("Press Enter", test1)
	_, _ = fmt.Scanln(&response)

	// Stop program
	mlog.Trace("Stop program\n\n")
}
