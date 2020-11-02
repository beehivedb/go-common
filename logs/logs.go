//Package logs  Leveled Log Library.
//Copyright (C) 2020 To All Authors. All rights reserved.
//Author: Ron.
//Date: 2020-08-08
//Version: 1.0
package logs

import (
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"strings"
	"time"
)

//Logs - define log interface.
type Logs interface {
	Trace(v ...interface{})
	Debug(v ...interface{})
	Info(v ...interface{})
	Error(v ...interface{})
	SetLevel(level string)
	SetOutput(out io.Writer)
	Panic(v ...interface{})
	Fatal(v ...interface{})
}

//Level - log level define.
type Level uint8

const (
	//TRACE level
	lt Level = iota
	ld
	li
	le
	lo
	lp
	lf
)

var (
	px            = "common-logs"
	lvl           = li
	out io.Writer = os.Stdout
)

func toLevel(level string) Level {
	up := strings.ToUpper(level)
	switch up {
	case "TRACE":
		return lt
	case "DEBUG":
		return ld
	case "INFO":
		return li
	case "ERROR":
		return le
	case "OFF":
		return lo
	case "PANIC":
		return lp
	case "FATAL":
		return lf
	default:
		return li
	}
}

//SetLevel - set global log level.
func SetLevel(level string) {
	lvl = toLevel(level)
}

//SetPrefix - set global log prefix.
func SetPrefix(prefix string) {
	px = prefix
}

//SetOutput - set global log outout.
func SetOutput(w io.Writer) {
	out = w
}

//New - Create New Log Instance.
func New(prefix string) Logs {
	return &Logger{
		lvl,
		prefix,
		out,
	}
}

//Logger - define Logger struct.
type Logger struct {
	level  Level
	prefix string
	out    io.Writer
}

//SetLevel - set log instance level.
func (l *Logger) SetLevel(level string) {
	l.level = toLevel(level)
}

//SetOutput - set log instance output.
func (l *Logger) SetOutput(writer io.Writer) {
	l.out = writer
}

//Trace - trace level message.
func (l *Logger) Trace(v ...interface{}) {
	if l.level == lt {
		log(l.out, l.prefix, "trace", v...)
	}
}

//Tracef - trace level message.
func (l *Logger) Tracef(ft string, v ...interface{}) {
	if l.level == lt {
		logf(l.out, l.prefix, "trace", ft, v...)
	}
}

//Trace - Global Trace Message.
func Trace(v ...interface{}) {
	if lvl == lt {
		log(out, px, "trace", v...)
	}
}

//Tracef - Global Trace Message.
func Tracef(ft string, v ...interface{}) {
	if lvl == lt {
		logf(out, px, "trace", ft, v...)
	}
}

//Info - info level message.
func (l *Logger) Info(v ...interface{}) {
	if l.level <= li {
		log(l.out, l.prefix, "info", v...)
	}
}

//Infof - info level message.
func (l *Logger) Infof(ft string, v ...interface{}) {
	if l.level <= li {
		logf(l.out, l.prefix, "info", ft, v...)
	}
}

//Info - global info level message.
func Info(v ...interface{}) {
	if lvl <= li {
		log(out, px, "info", v...)
	}
}

//Infof - global info level message.
func Infof(ft string, v ...interface{}) {
	if lvl <= li {
		logf(out, px, "info", ft, v...)
	}
}

//Debug - debug message.
func (l *Logger) Debug(v ...interface{}) {
	if l.level <= ld {
		log(l.out, l.prefix, "debug", v...)
	}
}

//Debugf - debug message.
func (l *Logger) Debugf(ft string, v ...interface{}) {
	if l.level <= ld {
		logf(l.out, l.prefix, "debug", ft, v...)
	}
}

//Debug - global debug message.
func Debug(v ...interface{}) {
	if lvl <= ld {
		log(out, px, "debug", v...)
	}
}

//Debugf - global debug message.
func Debugf(ft string, v ...interface{}) {
	if lvl <= ld {
		logf(out, px, "debug", ft, v...)
	}
}

//Error - error message.
func (l *Logger) Error(v ...interface{}) {
	if l.level <= le {
		log(l.out, l.prefix, "error", v...)
	}
}

//Errorf - format error message.
func (l *Logger) Errorf(ft string, v ...interface{}) {
	if l.level <= le {
		logf(l.out, l.prefix, "error", ft, v...)
	}
}

//Error - global error message.
func Error(v ...interface{}) {
	if lvl <= le {
		log(out, px, "error", v...)
	}
}

//Errorf - error format message.
func Errorf(ft string, v ...interface{}) {
	if lvl <= le {
		logf(out, px, "error", ft, v...)
	}
}

//Panic - Painc Message.
func (l *Logger) Panic(v ...interface{}) {
	if l.level <= lp {
		msg := fmt.Sprint(v...)
		panic(msg)
	}
}

//Paincf - Painc format message.
func (l *Logger) Paincf(ft string, v ...interface{}) {
	if l.level <= lp {
		msg := fmt.Sprintf(ft, v...)
		panic(msg)
	}
}

//Panic log - global panic message.
func Panic(v ...interface{}) {
	if lvl <= lp {
		msg := fmt.Sprint(v...)
		panic(msg)
	}
}

//Paincf log - format log message.
func Paincf(ft string, v ...interface{}) {
	if lvl <= lp {
		msg := fmt.Sprintf(ft, v...)
		panic(msg)
	}
}

//Fatal - print message then exit.
func (l *Logger) Fatal(v ...interface{}) {
	if l.level <= lf {
		log(l.out, l.prefix, "fatal", v...)
		os.Exit(1)
	}
}

//Fatalf - print format message and exit.
func (l *Logger) Fatalf(ft string, v ...interface{}) {
	if l.level <= lf {

		logf(l.out, l.prefix, "fatal", ft, v...)
		os.Exit(1)
	}
}

//Fatal - Global fatal message.
func Fatal(v ...interface{}) {
	if lvl <= lf {
		log(out, px, "fatal", v...)
		os.Exit(1)
	}
}

//Fatalf - format log message.
func Fatalf(fmt string, v ...interface{}) {
	if lvl <= lf {
		logf(out, px, "fatal", fmt, v...)
		os.Exit(1)
	}
}

func log(out io.Writer, prefix string, lvl string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(2)
	d := time.Now().Format(time.RFC3339)
	msg := fmt.Sprint(v...)
	_, err := fmt.Fprintf(out, "%s: %s [%s] %s@%d: %s\n", prefix, d, lvl, path.Base(file), line, msg)
	if err != nil {
		panic(err)
	}
}

func logf(out io.Writer, prefix string, lvl string, f string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(2)
	d := time.Now().Format(time.RFC3339)
	msg := fmt.Sprintf(f, v...)
	_, err := fmt.Fprintf(out, "%s: %s [%s] %s@%d: %s\n", prefix, d, lvl, path.Base(file), line, msg)
	if err != nil {
		panic(err)
	}
}
