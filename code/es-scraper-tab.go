type ChromeTab struct {
    ID          uint                // Id of the browser tab
    Context     *context.Context    // Pointer to the context that describes this tab
    Stop        *context.CancelFunc // CancelFunc to cancel it's context
    URL         string              // The url that is loaded in this tab
    State       tabState            // The working state of this tab.
    Store       storage.Storager    // S3 Storage where PDFs etc get stored
    MercuryURL  string              // Mercury URL
    ContentSize *dom.Rect           // Size and position of the content in the frame
}
