package utils

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"strings"
	"sync"

	p "github.com/pulumi/pulumi-go-provider"
)

func cleanupLine(s string) string {
	return strings.ReplaceAll(strings.TrimSpace(s), "\t", " ")
}

type PulumiLoggerHandler struct {
	ctx   context.Context
	lines []string
}

func (h *PulumiLoggerHandler) IngestReaders(done chan<- struct{}, stdout io.Reader, stderr io.Reader) error {
	logger := p.GetLogger(h.ctx)

	var wg sync.WaitGroup
	wg.Add(2)

	ingest := make(chan string)

	engine := func(r io.Reader) {
		s := bufio.NewScanner(r)

		for s.Scan() {
			txt := s.Text()
			logger.InfoStatus(cleanupLine(txt))
			ingest <- txt

		}
		wg.Done()
	}

	go engine(stdout)
	go engine(stderr)

	go func() {
		wg.Wait()
		close(ingest)
	}()

	go func() {
		for line := range ingest {
			h.lines = append(h.lines, line)
		}
		close(done)
	}()

	return nil
}

func (h *PulumiLoggerHandler) AugmentError(err error) error {
	return fmt.Errorf("\n%s\n%w", strings.Join(h.lines, "\n"), err)
}

func MakePulumiLogger(ctx context.Context) *PulumiLoggerHandler {
	return &PulumiLoggerHandler{ctx: ctx}
}
