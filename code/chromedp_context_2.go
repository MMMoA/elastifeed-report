// Set launch parameters for the browser
options := []chromedp.ExecAllocatorOption{
    // Skip checks
    chromedp.NoFirstRun,
    chromedp.NoDefaultBrowserCheck,
    // Set global user agent
    chromedp.UserAgent("Googlebot/2.1 (+http://www.google.com/bot.html)"),
    chromedp.Headless // Run headlessly
}
// Create only the allocator
browser, cancel := chromedp.NewExecAllocator(context.Background(), options...)
// Create two seperate BrowserContext (two seperate "windows")
window1, stop1 := chromedp.NewContext(browser)
window2, stop2 := chromedp.NewContext(browser)
// Create tabs for "window" 1
tab1_1, _ := chromedp.NewContext(window1)
tab1_2, _ := chromedp.NewContext(window1)
// Create tab for "window" 2
tab2_1, _ := chromedp.NewContext(window2)
