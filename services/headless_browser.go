package services

import (
	"log"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

func LaunchHeadlessBrowser() (selenium.WebDriver, error) {
	// Start Slenium WebDriver server instance
	opts := []selenium.ServiceOption{}
	service, err := selenium.NewChromeDriverService("/usr/bin/chromium-browser", 9515, opts...)
	if err != nil {
		log.Fatal("Error starting WebDriver service: ", err)
	}
	defer service.Stop()

	// Connect to WebDriver instance
	caps := selenium.Capabilities{"browserName": "chromium"}
	chromeCaps := chrome.Capabilities{
		Args: []string{
			"--headless", // Run Chrome in headless mode
		},
	}
	caps.AddChrome(chromeCaps)

	wd, err := selenium.NewRemote(caps, "")

	return wd, err
}
