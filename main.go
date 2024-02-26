package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/chromedp"
	"os"
	"path/filepath"
	"time"
)

func main() {
	const prefix = "out_"
	dir := fmt.Sprintf("%s%v", prefix, time.Now().Unix())

	var (
		input   = flag.String("i", "test.svg", "input")
		width   = flag.Int("w", 1920, "width")
		height  = flag.Int("h", 1080, "height")
		frames  = flag.Int("f", 10, "frames")
		seconds = flag.Float64("s", 1, "seconds")
	)
	flag.Parse()

	fmt.Printf("input   : %s\n", *input)
	fmt.Printf("width   : %d\n", *width)
	fmt.Printf("height  : %d\n", *height)
	fmt.Printf("frames  : %d\n", *frames)
	fmt.Printf("seconds : %f\n", *seconds)
	fmt.Printf("(fps)   : %f\n", float64(*frames)/float64(*seconds))
	fmt.Println("")

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	inPath := filepath.Join(wd, *input)
	url := fmt.Sprintf("file://%s", inPath)

	adjustedFrames := *frames - 1
	spf := float64(*seconds) / float64(adjustedFrames)

	if err := os.Mkdir(dir, 0755); err != nil {
		panic(err)
	}

	ctx, cancel := chromedp.NewContext(
		context.Background(),
	)
	defer cancel()

	if err := chromedp.Run(ctx,
		chromedp.Tasks{
			chromedp.EmulateViewport(int64(*width), int64(*height)),
			chromedp.ActionFunc(func(ctx context.Context) error {
				return emulation.SetDefaultBackgroundColorOverride().WithColor(&cdp.RGBA{0, 0, 0, 0}).Do(ctx)
			}),
			chromedp.Navigate(url),
			chromedp.WaitVisible("svg", chromedp.ByQuery),
			chromedp.Evaluate(`document.querySelector('svg').pauseAnimations();`, nil),
		},
	); err != nil {
		panic(err)
	}

	current := 0.0
	for i := 0; i < *frames; i++ {
		var buf []byte
		outPath := filepath.Join(dir, fmt.Sprintf("%d.png", i))
		if err := chromedp.Run(ctx,
			chromedp.Tasks{
				chromedp.Evaluate(fmt.Sprintf("document.querySelector('svg').setCurrentTime(%v);", current), nil),
				chromedp.Screenshot("svg", &buf, chromedp.ByQuery),
			},
		); err != nil {
			panic(err)
		}
		if err := os.WriteFile(outPath, buf, 0o644); err != nil {
			panic(err)
		}
		fmt.Printf("%fs %s\n", current, outPath)
		current += spf
	}
	fmt.Println("\nfinished")
}
