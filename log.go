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

// Debug prints error and returns true if err != nil,
// otherwise it just returns false without printing anything
func Debug(err error, i ...interface{}) bool {
	if err == nil {
		return false
	}

	ErrorStatus.printf("%v %s", err, fmt.Sprint(i...))
	return true
}

// Debug prints error and returns true if err != nil,
// otherwise it just returns false without printing anything.
// Prints any other values through "%#v"
func Debugv(err error, i ...interface{}) bool {
	if err == nil {
		return false
	}

	f := "%v "
	for j := 0; j < len(i); j++ {
		f += "%#v "
	}
	items := []interface{}{err}
	ErrorStatus.printf(f, append(items, i...)...)
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

func Infov(i ...interface{}) {
	f := ""
	for j := 0; j < len(i); j++ {
		f += "%#v "
	}
	InfoStatus.printf(f, i...)
}

// Pretty converts `s` to json with <TAB> indent and prints it out
func Pretty(s interface{}) {
	sJson, _ := json.MarshalIndent(s, "", "\t")
	InfoStatus.print(string(sJson))
}
