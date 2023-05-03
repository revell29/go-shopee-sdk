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

func Test_GetConversationList(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/sellerchat/get_conversation_list", app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("get_conversation_resp.json")))

	res, err := client.Chat.GetConversationList(123456, accessToken, shopee.GetConversationParamsRequest{
		PageSize:  50,
		Direction: "latest",
		Type:      "all",
	})

	if err != nil {
		t.Errorf("Chat.GetMessages error: %s", err)
	}

	t.Logf("return tok: %#v", res)

	var expectedConversationID string = "38732689394223980"
	if res.Response.ConversationsList[0].ConversationID != expectedConversationID {
		t.Errorf("ConversationList.MessageID returned %+v, expected %+v", res.Response.ConversationsList[0].ConversationID, expectedConversationID)
	}
}

func Test_SendMessage(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/sellerchat/send_message", app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("send_message_resp.json")))

	var req shopee.SendMessageRequest
	loadMockData("send_message_req.json", &req)

	res, err := client.Chat.SendMessage(shopID, accessToken, req)

	if err != nil {
		t.Errorf("Product.AddItem error: %s", err)
	}

	t.Logf("Chat.SendMessage: %#v", res)
	var expectedID int = 9018093
	if res.Response.ToID != expectedID {
		t.Errorf("ToID returned %+v, expected %+v", res.Response.ToID, expectedID)
	}
}
