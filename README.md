# About
It's my a little study project. Gita is a simple logger that can print logs with different text styles and format options like timestamp or filename with line of code where log was called 

# Getting started

### Create logger with default config

```go
logger := gita.NewLogger(nil)
// or
logger := gita.NewLogger(&gita.Config{})
```

### Use logger

```go
logger.Log(gita.InfoLevel, "It's Gita's default log")
logger.Warning("Unable to get sth, set default value")
```

### Parse level from string

There are only 4 levels of logs in gita: ***info, warning, error, fatal***

```go
lvl, err := gita.ParseLevel("warning")
if err != nil {
	fmt.Println(err)
}

logger := gita.NewLogger(&gita.Config{
	Level: lvl,
})

logger.Log(gita.InfoLevel, "It's Gita's default log") // Skipped
logger.Warning("Unable to get sth, set default value")
```

### Out and Err writers

You can customize out and err writers, err used for error and fatal logs. Defaults to ***os.Stdout*** and ***os.Stderr***

```go
logger := gita.NewLogger(&gita.Config{
	Out: ..., // Custom io.Writer	
	Err: ..., // Custom io.Writer	
})
```

### Log format

You can customize output of the context of the message, like time, location (filename:line), stacktrace, current message number (ID) and also change the time formatter

```go
logger := gita.NewLogger(&gita.Config{
	Format: &gita.Formatter{
		DisplayID:         true,
		DisplayTimestamp:  true,
		DisplayLocation:   true,
		DisplayStackTrace: true,
		TimestampFormat:   time.RubyDate,
	},
})

logger.Log(gita.InfoLevel, "It's Gita's default log")
// Prints: #1 Tue Jan 30 00:29:36 +0300 2024 [INFO] (main.go:20) It's Gita's default log
```

### Customize colors and font style

You can customize style of every part of the log: text can be bold, itallic, underlined red with magenta background, and so on. Location's style always the same as level style

```go
logger := gita.NewLogger(&gita.Config{
	Format: &gita.Formatter{
		LevelStyle: gita.LevelStyle{
			gita.InfoLevel: {
				Style: gita.Bold,
				Color: gita.Red,
			},
			gita.WarningLevel: {
				Color: gita.Magenta,
			},
			gita.ErrorLevel: {
				Style: gita.Italic,
				Color: gita.VividGreen,
			},
		},
		MessageStyle: &gita.TextStyle{
			Style:   gita.Bold,
			Color:   gita.VividRed,
			BGColor: gita.RedBG,
		},
	},
})
```
