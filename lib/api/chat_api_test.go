package api_test

import (
	"testing"

	"github.com/pdcgo/tokopedia_lib/lib/api"
	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/stretchr/testify/assert"
)

func TestChatApi(t *testing.T) {
	tapi, saveSession := scenario.GetTokopediaApiClient()
	defer saveSession()

	t.Run("test get chat list api", func(t *testing.T) {

		hasil, err := tapi.GetChatList(api.ChatListVar{
			Page:     1,
			PerPage:  10,
			Filter:   "all",
			Platform: "dekstop",
			Order:    "desc",
			Tab:      "tab-seller",
		})
		assert.Nil(t, err)
		assert.NotEmpty(t, hasil)
		assert.Equal(t, 10, len(hasil.Data.ChatList.List))
	})

	t.Run("test get chat room api", func(t *testing.T) {

		hasil, err := tapi.GetChatRoom(api.ChatRoomVar{
			Page:       1,
			PerPage:    10,
			MessageID:  2483170385,
			IsTextOnly: true,
		})
		assert.Nil(t, err)
		assert.NotEmpty(t, hasil)
		assert.NotEmpty(t, hasil.Data.ChatReplies.List)
	})

	t.Run("test get chat attachments api", func(t *testing.T) {

		hasil, err := tapi.GetChatAttachments(api.ChatAttachmentVar{
			MessageID:     2560136298,
			AttachmentIDs: "2758317382",
		})
		assert.Nil(t, err)
		assert.NotEmpty(t, hasil)
		assert.NotEmpty(t, hasil.Data.ChatAttachments.List)
		assert.NotEmpty(t, hasil.Data.ChatAttachments.List[0].Attributes.ProductProfile)
	})

	t.Run("test chat pin api", func(t *testing.T) {

		hasil, err := tapi.ChatPin(2560136298)
		assert.Nil(t, err)
		assert.NotEmpty(t, hasil)
		assert.True(t, hasil.Data.ChatPin.Success)
	})

	t.Run("test chat unpin api", func(t *testing.T) {

		hasil, err := tapi.ChatUnpin(2560136298)
		assert.Nil(t, err)
		assert.NotEmpty(t, hasil)
		assert.True(t, hasil.Data.ChatUnpin.Success)
	})

	t.Run("test get chat search api", func(t *testing.T) {

		hasil, err := tapi.GetChatSearch(api.ChatSearchVar{
			IsSeller: 1,
			Keyword:  "made",
			Size:     10,
			Status:   1,
		})
		assert.Nil(t, err)
		assert.NotEmpty(t, hasil)
		assert.NotEmpty(t, hasil.Data.ChatSearch.Contact.Data)
	})
}
