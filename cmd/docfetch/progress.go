package main

import (
	"fmt"
	"net/url"
	"os"
	"sync"
	"time"

	"github.com/AlphaTechini/doc-fetch/v2/pkg/fetcher"
)

type progressDisplay struct {
	host        string
	interactive bool
	startedAt   time.Time
	updates     chan fetcher.Progress
	done        chan struct{}
	wg          sync.WaitGroup
}

func newProgressDisplay(rawURL string, enabled bool) *progressDisplay {
	host := rawURL
	if parsed, err := url.Parse(rawURL); err == nil && parsed.Host != "" {
		host = parsed.Host
	}

	interactive := false
	if info, err := os.Stderr.Stat(); err == nil {
		interactive = info.Mode()&os.ModeCharDevice != 0
	}
	return &progressDisplay{
		host:        host,
		interactive: enabled && interactive,
		updates:     make(chan fetcher.Progress, 1),
		done:        make(chan struct{}),
	}
}

func (display *progressDisplay) Start() {
	display.startedAt = time.Now()
	if !display.interactive {
		return
	}

	display.wg.Add(1)
	go func() {
		defer display.wg.Done()
		ticker := time.NewTicker(100 * time.Millisecond)
		defer ticker.Stop()

		frames := []byte{'|', '/', '-', '\\'}
		frame := 0
		latest := fetcher.Progress{}
		width := 0
		for {
			select {
			case update := <-display.updates:
				latest = update
			case <-ticker.C:
				line := fmt.Sprintf("%c Crawling %s | %d processed | %d discovered | %d failed | %ds",
					frames[frame%len(frames)], display.host, latest.Processed, latest.Discovered,
					latest.Failed, int(time.Since(display.startedAt).Seconds()))
				frame++
				if len(line) > width {
					width = len(line)
				}
				fmt.Fprintf(os.Stderr, "\r%-*s", width, line)
			case <-display.done:
				fmt.Fprintf(os.Stderr, "\r%-*s\r", width, "")
				return
			}
		}
	}()
}

func (display *progressDisplay) Update(progress fetcher.Progress) {
	if !display.interactive {
		return
	}
	select {
	case display.updates <- progress:
	default:
		select {
		case <-display.updates:
		default:
		}
		select {
		case display.updates <- progress:
		default:
		}
	}
}

func (display *progressDisplay) Stop() {
	if !display.interactive {
		return
	}
	close(display.done)
	display.wg.Wait()
}
