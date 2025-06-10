package main

	
import (
	"log"
    "os"
)


type LogLevel int

const (
    LogLevelError LogLevel = iota
    LogLevelWarning
    LogLevelInfo
)

type LogExtended struct {
    *log.Logger
    logLevel LogLevel // LogLevel это enum
}

func main() {
    logger := NewLogExtended()
    logger.SetLogLevel(LogLevelWarning)
    logger.Infoln("Не должно напечататься")
    logger.Warnln("Hello")
    logger.Errorln("World")
    logger.Println("Debug")
}

func NewLogExtended() *LogExtended {
    return &LogExtended{
        Logger:   log.New(os.Stderr, "", log.LstdFlags),
        logLevel: LogLevelError,
    }
}

func (l LogLevel) IsValid() bool {
    switch l {
    case LogLevelInfo, LogLevelWarning, LogLevelError:
        return true
    default:
        return false
    }      
}

func (l *LogExtended) SetLogLevel(loglevel LogLevel) {
    if !loglevel.IsValid() {
        return
    }
    l.logLevel = loglevel
}

func (l *LogExtended) Infoln(msg string) {
    l.println(LogLevelInfo, "INFO ", msg)
}

func (l *LogExtended) Warnln(msg string) {
    l.println(LogLevelWarning, "WARN ", msg)
}

func (l *LogExtended) Errorln(msg string) {
    l.println(LogLevelError, "ERR ", msg)
}

func (l *LogExtended) println(srcLogLvl LogLevel, prefix, msg string) {
    if l.logLevel < srcLogLvl {
        return
    }

    l.Logger.Println(prefix + msg)
}