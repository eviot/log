package log

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

var NeedColor = true

var (
	ErrorStatus   = NewStatus(lightRed, normal, "ERR")
	InfoStatus    = NewStatus(lightBlue, normal, "INF")
	WarningStatus = NewStatus(lightYellow, normal, "WAR")
)

type Status struct {
	textPrefix    string
	coloredPrefix string
}

func NewStatus(clr int, fontSize int, text string) Status {
	return Status{
		textPrefix:    fmt.Sprintf("[%s]", text),
		coloredPrefix: fmt.Sprintf("[%s]", render(clr, fontSize, text)),
	}
}

func (s Status) prefix() string {
	if NeedColor {
		return s.coloredPrefix
	}
	return s.textPrefix
}

func (s Status) printString(str string) {
	internalFormat := "%s %s %s(%d): %s\n%s\n\n"
	dateTime := time.Now().Format("02.01.06 15:04:05.000")
	pc, file, line, _ := runtime.Caller(3)
	dir, _ := os.Getwd()
	filename := strings.Replace(file, dir, "", 1)
	function := runtime.FuncForPC(pc)
	fmt.Printf(internalFormat, s.prefix(), dateTime, filename, line, function.Name(), str)
}

func (s Status) print(i ...interface{}) {
	f := ""
	for _ = range i {
		f += "%v "
	}
	s.printString(fmt.Sprintf(f, i...))
}

func (s Status) println(i ...interface{}) {
	f := ""
	for _ = range i {
		f += "%v "
	}
	f += "\n"
	s.printString(fmt.Sprintf(f, i...))
}

func (s Status) printf(f string, i ...interface{}) {
	s.printString(fmt.Sprintf(f, i...))
}

func Debug(err interface{}, i ...interface{}) bool {
	if err == nil {
		return false
	}
	ErrorStatus.printf("%#v %s", err, fmt.Sprint(i...))
	return true
}

func Fatalln(i ...interface{}) {
	ErrorStatus.println(i...)
	panic("")
}

func Info(i ...interface{}) {
	InfoStatus.print(i...)
}

func Infof(f string, i ...interface{}) {
	InfoStatus.printf(f, i...)
}

func Pretty(s interface{}) {
	sJson, _ := json.MarshalIndent(s, "", "\t")
	InfoStatus.print(string(sJson))
}
