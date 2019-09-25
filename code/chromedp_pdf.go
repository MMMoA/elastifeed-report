tasks := chromedp.Tasks{
    tab.setDeviceMetricsAction(),
    chromedp.ActionFunc(func(ctx context.Context) error {
        var err error
        result, _, err = page.PrintToPDF().WithPrintBackground(true).
            WithTransferMode(page.PrintToPDFTransferModeReturnAsBase64).Do(ctx)
        return err
    }),
}
