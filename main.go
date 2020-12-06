package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/vodafon/swork"
)

var (
	flagConfig = flag.String("c", "", "path to json file with regext. If empty use default")
	flagProcs  = flag.Int("procs", 10, "concurrency")
)

type Processor struct {
	w   io.Writer
	cfg []Re
}

func (obj Processor) Process(line string) {
	for _, re := range obj.cfg {
		ss := re.Regexp.FindString(line)
		if ss == "" {
			continue
		}

		fmt.Fprintf(obj.w, "%s: %q\n", re.Name, ss)
	}
}

func main() {
	flag.Parse()
	if *flagProcs < 1 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	processor := Processor{
		w:   os.Stdout,
		cfg: MustLoadConfig(*flagConfig),
	}

	w := swork.NewWorkerGroup(*flagProcs, processor)

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		w.StringC <- sc.Text()
	}

	close(w.StringC)

	if err := sc.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to read input: %s\n", err)
	}

	w.Wait()
}
