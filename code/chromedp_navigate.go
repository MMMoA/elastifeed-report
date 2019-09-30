task := chromedp.ActionFunc(func(ctx context.Context) error {
    emulation.SetUserAgentOverride(userAgent).Do(ctx)
    _, _, _, err := page.Navigate(url).WithReferrer("https://www.google.de/").Do(ctx)
    if err != nil {
        return err
    }
    wait := waitLoadedOrTimeout(ctx)
    _, _, contentSize, err := page.GetLayoutMetrics().Do(ctx)
    if err != nil {
        return err
    }
    tab.ContentSize = contentSize
    return wait
})
