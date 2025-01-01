package services

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

func LaunchHeadlessBrowser(path string) (context.Context, context.CancelFunc) {
	// Create a new context
	ctx, _ := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(customLogger),
	)

	// Create a timeout
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)

	// Navigate to a page to ensure the browser is working
	err := chromedp.Run(ctx,
		chromedp.Navigate(path),
		chromedp.WaitVisible(`body`, chromedp.ByQuery),
	)
	if err != nil {
		log.Printf("Error launching browser: %v", err)
		cancel()
		return nil, nil
	}

	log.Println("Headless browser launched successfully")
	return ctx, cancel
}

func customLogger(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	if !strings.Contains(msg, "could not unmarshal event") {
		log.Printf(format, args...)
	}
}
