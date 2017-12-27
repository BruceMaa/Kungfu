package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

type KungfuLogger struct {
	DebugLogger *log.Logger
	TraceLogger *log.Logger
	InfoLogger  *log.Logger
	WarnLogger  *log.Logger
	ErrorLogger *log.Logger
	FatalLogger *log.Logger
	PanicLogger *log.Logger
}

var (
	DebugHandle io.Writer = os.Stdout
	TraceHandle io.Writer = os.Stdout
	InfoHandle  io.Writer = os.Stdout
	WarnHandle  io.Writer = os.Stdout
	ErrorHandle io.Writer = os.Stderr
	FatalHandle io.Writer = os.Stderr
	PanicHandle io.Writer = os.Stderr
)

var debugOnce sync.Once
var traceOnce sync.Once
var infoOnce sync.Once
var warnOnce sync.Once
var errorOnce sync.Once
var fatalOnce sync.Once
var panicOnce sync.Once

var Logger = &KungfuLogger{}

func (kLong *KungfuLogger) Debugf(format string, v ...interface{}) {
	debugOnce.Do(func() {
		fmt.Fprintln(DebugHandle, "debugOnce.Do init Debugf")
		kLong.DebugLogger = log.New(DebugHandle, "[DEBUG]\t", log.LstdFlags|log.Lshortfile)
	})
	kLong.DebugLogger.Output(2, fmt.Sprintf(format, v...))
}

func (kLong *KungfuLogger) Debug(v ...interface{}) {
	debugOnce.Do(func() {
		fmt.Fprintln(DebugHandle, "debugOnce.Do init Debug")
		kLong.DebugLogger = log.New(DebugHandle, "[DEBUG]\t", log.LstdFlags|log.Lshortfile)
	})
	kLong.DebugLogger.Output(2, fmt.Sprintln(v...))
}

func (kLong *KungfuLogger) Tracef(format string, v ...interface{}) {
	traceOnce.Do(func() {
		fmt.Fprintln(TraceHandle, "traceOnce.Do init Tracef")
		kLong.TraceLogger = log.New(TraceHandle, "[TRACE]\t", log.LstdFlags|log.Lshortfile)
	})
	kLong.TraceLogger.Output(2, fmt.Sprintf(format, v...))
}

func (kLong *KungfuLogger) Trace(v ...interface{}) {
	traceOnce.Do(func() {
		fmt.Fprintln(TraceHandle, "traceOnce.Do init Trace")
		kLong.TraceLogger = log.New(TraceHandle, "[TRACE]\t", log.LstdFlags|log.Lshortfile)
	})
	kLong.TraceLogger.Output(2, fmt.Sprintln(v...))
}

func (kLong *KungfuLogger) Infof(format string, v ...interface{}) {
	infoOnce.Do(func() {
		fmt.Fprintln(InfoHandle, "infoOnce.Do init Infof")
		kLong.InfoLogger = log.New(InfoHandle, "[INFO]\t", log.LstdFlags|log.Lshortfile)
	})
	kLong.InfoLogger.Output(2, fmt.Sprintf(format, v...))
}

func (kLong *KungfuLogger) Info(v ...interface{}) {
	infoOnce.Do(func() {
		fmt.Fprintln(InfoHandle, "infoOnce.Do init Info")
		kLong.InfoLogger = log.New(InfoHandle, "[INFO]\t", log.LstdFlags|log.Lshortfile)
	})
	kLong.InfoLogger.Output(2, fmt.Sprintln(v...))
}

func (kLong *KungfuLogger) Warningf(format string, v ...interface{}) {
	warnOnce.Do(func() {
		fmt.Fprintln(WarnHandle, "warnOnce.Do init WarningF")
		kLong.WarnLogger = log.New(WarnHandle, "[WARNING]\t", log.LstdFlags|log.Lshortfile)
	})
	kLong.WarnLogger.Output(2, fmt.Sprintf(format, v...))
}

func (kLong *KungfuLogger) Warning(v ...interface{}) {
	warnOnce.Do(func() {
		fmt.Fprintln(WarnHandle, "warnOnce.Do init Warning")
		kLong.WarnLogger = log.New(WarnHandle, "[WARNING]\t", log.LstdFlags|log.Lshortfile)
	})
	kLong.WarnLogger.Output(2, fmt.Sprintln(v...))
}

func (kLong *KungfuLogger) Errorf(format string, v ...interface{}) {
	errorOnce.Do(func() {
		fmt.Fprintln(ErrorHandle, "errorOnce.Do init Errorf")
		kLong.ErrorLogger = log.New(ErrorHandle, "[ERROR]\t", log.LstdFlags|log.Lshortfile)
	})
	kLong.ErrorLogger.Output(2, fmt.Sprintf(format, v...))
}

func (kLong *KungfuLogger) Error(v ...interface{}) {
	errorOnce.Do(func() {
		fmt.Fprintln(ErrorHandle, "errorOnce.Do init Error")
		kLong.ErrorLogger = log.New(ErrorHandle, "[ERROR]\t", log.LstdFlags|log.Lshortfile)
	})
	kLong.ErrorLogger.Output(2, fmt.Sprintln(v...))
}

func (kLong *KungfuLogger) Fatalf(format string, v ...interface{}) {
	fatalOnce.Do(func() {
		fmt.Fprintln(FatalHandle, "fatalOnce.Do init Fatalf")
		kLong.FatalLogger = log.New(FatalHandle, "[FATAL]\t", log.LstdFlags|log.Lshortfile)
	})
	kLong.FatalLogger.Output(2, fmt.Sprintf(format, v...))
}

func (kLong *KungfuLogger) Fatal(v ...interface{}) {
	fatalOnce.Do(func() {
		fmt.Fprintln(FatalHandle, "fatalOnce.Do init Fatal")
		kLong.FatalLogger = log.New(FatalHandle, "[FATAL]\t", log.LstdFlags|log.Lshortfile)
	})
	kLong.FatalLogger.Output(2, fmt.Sprintln(v...))
}

func (kLong *KungfuLogger) Panicf(format string, v ...interface{}) {
	panicOnce.Do(func() {
		fmt.Fprintln(PanicHandle, "panicOnce.Do init Panicf")
		kLong.PanicLogger = log.New(PanicHandle, "[PANIC]\t", log.LstdFlags|log.Lshortfile)
	})
	kLong.PanicLogger.Output(2, fmt.Sprintf(format, v...))
}

func (kLong *KungfuLogger) Panic(v ...interface{}) {
	panicOnce.Do(func() {
		fmt.Fprintln(PanicHandle, "panicOnce.Do init Panic")
		kLong.PanicLogger = log.New(PanicHandle, "[PANIC]\t", log.LstdFlags|log.Lshortfile)
	})
	kLong.PanicLogger.Output(2, fmt.Sprintln(v...))
}
