package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"time"

	"github.com/aybabtme/hist/Godeps/_workspace/src/github.com/aybabtme/uniplot/histogram"
	"github.com/aybabtme/hist/Godeps/_workspace/src/github.com/codegangsta/cli"
)

const (
	version = "devel"
	appName = "hist"
)

var (
	reFlag = cli.StringFlag{
		Name:  "re",
		Usage: "regexp used to match numbers in stdin, by default matches the first integer found in a line",
		Value: `^(\d+)$`,
	}

	binsFlag = cli.IntFlag{
		Name:  "bins",
		Usage: "number of bins to use",
		Value: 40,
	}

	widthFlag = cli.IntFlag{
		Name:  "width",
		Usage: "width of the histogram",
		Value: 40,
	}

	formatFlag = cli.StringFlag{
		Name:  "format",
		Usage: "format func to use when displaying the values",
		Value: "plain",
	}

	formatFunc = map[string]histogram.FormatFunc{
		"plain": func(v float64) string { return fmt.Sprintf("%f", v) },
		"unix": func(v float64) string {
			sec := math.Floor(v)
			nsec := (v - sec) * 1e9
			return time.Unix(int64(sec), int64(nsec)).Format(time.RFC3339)
		},
		"elapsed-seconds": func(v float64) string {
			return fmt.Sprintf("%.02fs", v)
		},
	}
	validFormatFunc = func() []string {
		keys := []string{}
		for k := range formatFunc {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		return keys
	}()
)

func main() {
	log.SetFlags(0)
	log.SetPrefix(appName + ": ")
	if err := newApp().Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func newApp() *cli.App {
	app := cli.NewApp()
	app.Name = appName
	app.Version = version
	app.Email = "antoinegrondin@gmail.com"
	app.Author = "Antoine Grondin"
	app.Usage = "plot histograms on stdout from input parsed on stdin"
	app.Flags = []cli.Flag{reFlag, binsFlag, widthFlag, formatFlag}
	app.Commands = []cli.Command{
		{
			Name:  "preview",
			Usage: "preview what matches in stdin",
			Action: func(ctx *cli.Context) {
				re, err := regexp.Compile(ctx.GlobalString(reFlag.Name))
				if err != nil {
					log.Fatalf("invalid regexp: %v", err)
				}
				if err := preview(os.Stdin, re); err != nil {
					log.Fatalf("can't read input: %v", err)
				}
			},
		},

		{
			Name:  "plot",
			Usage: "plots what matches in stdin",
			Action: func(ctx *cli.Context) {
				bins := ctx.GlobalInt(binsFlag.Name)
				width := ctx.GlobalInt(widthFlag.Name)
				formatter := ctx.GlobalString(formatFlag.Name)
				formatf, ok := formatFunc[formatter]
				if !ok {
					log.Fatalf("invalid format: %q\nuse of of: %v", formatter, validFormatFunc)
				}

				re, err := regexp.Compile(ctx.GlobalString(reFlag.Name))
				if err != nil {
					log.Fatalf("invalid regexp: %v", err)
				}

				numbers, err := parse(os.Stdin, re)
				if err != nil {
					log.Fatalf("can't read input: %v", err)
				}

				hist := histogram.Hist(bins, numbers)
				if err := histogram.Fprintf(
					os.Stderr,
					hist,
					histogram.Linear(width),
					formatf,
				); err != nil {
					log.Fatalf("can't plot: %v", err)
				}
			},
		},
	}
	return app
}

func preview(r io.Reader, re *regexp.Regexp) error {
	raw, err := ioutil.ReadAll(r) // todo: regexp match on a stream
	if err != nil {
		return err
	}
	subs := re.FindAllStringSubmatch(string(raw), -1)
	if len(subs) == 0 {
		return fmt.Errorf("no match for %q", re.String())
	}
	for _, sub := range subs {
		fmt.Fprintf(os.Stdout, "%q\n", sub[1])
	}
	return nil
}

func parse(r io.Reader, re *regexp.Regexp) ([]float64, error) {
	raw, err := ioutil.ReadAll(r) // todo: regexp match on a stream
	if err != nil {
		return nil, err
	}

	subs := re.FindAllStringSubmatch(string(raw), -1)
	if len(subs) == 0 {
		return nil, fmt.Errorf("no match for %q", re.String())
	}

	outs := make([]float64, 0, len(subs))
	for _, sub := range subs {
		f, err := strconv.ParseFloat(sub[1], 64)
		if err != nil {
			return outs, err
		}
		outs = append(outs, f)
	}
	return outs, nil
}
