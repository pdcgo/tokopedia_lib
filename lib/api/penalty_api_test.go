package api_test

import (
	"strconv"
	"testing"
	"time"

	"github.com/pdcgo/tokopedia_lib/lib/api"
	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/stretchr/testify/assert"
)

func TestGettingPenaltyApi(t *testing.T) {
	client, saveSession := scenario.GetTokopediaApiClient()
	defer saveSession()

	t.Run("getting penalty api", func(t *testing.T) {
		stdate := time.Now()
		stdate = stdate.AddDate(0, -3, 0)

		t.Log(stdate)

		shopID := strconv.FormatInt(client.AuthenticatedData.UserShopInfo.Info.ShopID, 10)
		hasil, err := client.ShopScorePenaltyDetail(&api.ShopScorePenaltyDetailVar{
			ShopID:    shopID,
			Page:      1,
			Total:     50,
			StartDate: stdate,
			EndDate:   time.Now(),
			Sort:      0,
			Source:    "icarus",
		})

		assert.NotEmpty(t, hasil)
		assert.Nil(t, err)
	})

	t.Run("getting summary", func(t *testing.T) {
		enddate := time.Now()
		stdate := enddate.AddDate(0, -2, 0)

		shopID := strconv.FormatInt(client.AuthenticatedData.UserShopInfo.Info.ShopID, 10)

		hasil, err := client.ShopScorePenaltySummary(&api.ShopScorePenaltySummaryVar{
			StartDate: stdate,
			EndDate:   enddate,
			ShopID:    shopID,
			Source:    "icarus",
		})

		assert.NotEmpty(t, hasil)
		assert.Nil(t, err)

		t.Log(hasil.Data.ShopScorePenaltySummary.Error.Message)

		assert.NotEqual(t, hasil.Data.ShopScorePenaltySummary.Result.Penalty, 0)
	})
}
