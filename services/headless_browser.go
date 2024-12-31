package services

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

func LaunchHeadlessBrowser() (selenium.WebDriver, error) {
	// Check if ChromeDriver is installed
	chromeDriverPath, err := exec.LookPath("/usr/bin/chromium-browser")
	if err != nil {
		return nil, fmt.Errorf("ChromeDriver not found in PATH: %v", err)
	}

	// Start Slenium WebDriver server instance
	opts := []selenium.ServiceOption{
		selenium.Output(nil),
	}
	service, err := selenium.NewChromeDriverService(chromeDriverPath, 9515, opts...)
	if err != nil {
		return nil, fmt.Errorf("error starting ChromeDriver service: %v", err)
	}
	defer service.Stop()

	// Connect to WebDriver instance
	caps := selenium.Capabilities{"browserName": "chromium"}
	chromeCaps := chrome.Capabilities{
		Path: "/usr/bin/chromium-browser",
		Args: []string{
			"--headless", // Run Chrome in headless mode
			"--no-sandbox",
            "--disable-dev-shm-usage",
		},
	}
	caps.AddChrome(chromeCaps)

	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", 9515))
	if err != nil {
        return nil, fmt.Errorf("error connecting to ChromeDriver: %v", err)
    }

	log.Println("Headless browser launched successfully")
	return wd, nil
}
