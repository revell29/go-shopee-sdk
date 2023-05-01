package tests

import (
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/revell29/go-shopee-sdk/shopee"
)

func Test_GetMessages(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/sellerchat/get_message", app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("get_messages_resp.json")))

	res, err := client.Chat.GetMessage(123456, accessToken, shopee.GetMessageParamsRequest{
		PageSize: 50,
	})

	if err != nil {
		t.Errorf("Chat.GetMessages error: %s", err)
	}

	t.Logf("return tok: %#v", res)

	var expectedMsgID string = "msg12345"
	if res.Response.MessagesList[0].MessageID != expectedMsgID {
		t.Errorf("MessageList.MessageID returned %+v, expected %+v", res.Response.MessagesList[0].MessageID, expectedMsgID)
	}
}
