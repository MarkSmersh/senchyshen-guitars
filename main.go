package main

import (
	"log"
	"log/slog"
	"sync"

	"github.com/MarkSmersh/senchyshen-guitars/api"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	slog.SetLogLoggerLevel(slog.LevelDebug)

	api.Init()

	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
