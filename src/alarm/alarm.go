package alarm

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

type errorString struct {
	s string
}

type errorInfo struct {
	Time     string `json:"time"`
	Alarm    string `json:"alarm"`
	Message  string `json:"message"`
	Filename string `json:"filename"`
	Line     int    `json:"line"`
	Funcname string `json:"funcname"`
}

func (e *errorString) Error() string {
	return e.s
}

func New(text string) error {
	alarm("INFO", text, 2)
	return &errorString{text}
}

func Email(text string) error {
	alarm("EMAIL", text, 2)
	return &errorString{text}
}

func Sms(text string) error {
	alarm("SMS", text, 2)
	return &errorString{text}
}

func Wechat(text string) error {
	alarm("WX", text, 2)
	return &errorString{text}
}

func alarm(level string, str string, skip int) {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	fileName, line, functionName := "?", 0, "?"
	pc, fileName, line, ok := runtime.Caller(skip)
	if ok {
		functionName = runtime.FuncForPC(pc).Name()
		functionName = filepath.Ext(functionName)
		functionName = strings.TrimPrefix(functionName, ".")
	}
	var msg = errorInfo{
		Time:     currentTime,
		Alarm:    level,
		Message:  str,
		Filename: fileName,
		Line:     line,
		Funcname: functionName,
	}
	jsons, errs := json.Marshal(msg)
	if errs != nil {
		fmt.Println("json marshal error:", errs)
	}
	errsJsonInfo := string(jsons)
	fmt.Println(errsJsonInfo)

	if level == "EMAIL" {

	} else if level == "SMS" {

	} else if level == "WX" {

	} else if level == "INFO" {

	} else if level == "PANIC" {

	}
}

func Panic(text string) error {
	alarm("PANIC", text, 4)
	return &errorString{text}
}
