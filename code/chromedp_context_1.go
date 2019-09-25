// New browser with default BrowserContext and default tab.
base, stop := chromedp.NewContext(context.Background())
defer stop
// Create a second tab in the same BrowserContext as the default tab
tab, _ := chromedp.NewContext(base)
