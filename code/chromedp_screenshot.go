tasks := chromedp.Tasks{
    tab.setDeviceMetricsAction(),
    chromedp.ActionFunc(func(ctx context.Context) error {
        viewport := page.Viewport{
            X:      tab.ContentSize.X,
            Y:      tab.ContentSize.Y,
            Width:  tab.ContentSize.Width,
            Height: tab.ContentSize.Height,
            Scale:  1,
        }
        var caperr error
        result, caperr = page.CaptureScreenshot().WithClip(&viewport).Do(ctx)
        return caperr
    }),
}
