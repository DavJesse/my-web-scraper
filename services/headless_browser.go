package services

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func LaunchHeadlessBrowser() (context.Context, context.CancelFunc) {
	// Create a new context
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)

	// Create a timeout
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)

	// Navigate to a page to ensure the browser is working
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://example.com`),
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
