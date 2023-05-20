package report

import (
	"context"

	"github.com/chromedp/chromedp"
)

func (rep *CekReport) GetProductReport(ctx context.Context) error {

	chromedp.Run(ctx,
		chromedp.Navigate("https://seller.tokopedia.com/manage-product"),
	)

	return nil
}
