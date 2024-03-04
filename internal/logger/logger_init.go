package logger

import "log/slog"

var Log *Logger

func init() {
    config := LoggerConfig{
        LoggerType: "text",
        Level:      slog.LevelInfo,
        AddSource:  true,
    }
	
    logger, err := New(config)
    if err != nil {
        panic(err)
    }
    Log = logger
}