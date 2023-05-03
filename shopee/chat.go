package shopee

type ChatService interface {
	GetMessage(shopID uint64, token string, params GetMessageParamsRequest) (*GetMessageResponse, error)
	GetConversationList(shopID uint64, token string, params GetConversationParamsRequest) (*GetConversationResponse, error)
	SendMessage(shopID uint64, token string, request SendMessageRequest) (*GetSendMessageResponse, error)
}

type GetMessageParamsRequest struct {
	Offset         string  `url:"offset"`
	PageSize       int     `url:"page_size"`
	ConversationID int64   `url:"conversation_id"`
	MessageIdList  []int64 `url:"message_id_list"`
}

type GetMessageResponse struct {
	BaseResponse

	Response GetMessageDataResponse `json:"response"`
}

type GetMessageDataResponse struct {
	MessagesList []Messages `json:"messages"`
}

type Messages struct {
	MessageID        string         `json:"message_id"`
	MessageType      string         `json:"message_type"`
	FromID           int64          `json:"from_id"`
	FromShopID       int64          `json:"from_shop_id"`
	ToID             int64          `json:"to_id"`
	ToShopID         int64          `json:"to_shop_id"`
	ConversationID   string         `json:"conversation_id"`
	CreatedTimeStamp int64          `json:"created_timestamp"`
	Region           string         `json:"region"`
	Status           string         `json:"status"`
	Source           string         `json:"source"`
	Content          ContentMessage `json:"content"`
	MessageOption    int            `json:"message_option"`
}

type ContentMessage struct {
	Text             string `json:"text"`
	Url              string `json:"url"`
	ThumbHeight      int    `json:"thumb_height"`
	ThumbWidth       int    `json:"thumb_width"`
	FileServerID     int64  `json:"file_server_id"`
	ShopID           int64  `json:"shop_id"`
	OfferID          int    `json:"offer_id"`
	ProductID        int    `json:"product_id"`
	TaxValue         string `json:"tax_value"`
	PriceBeforeTax   string `json:"price_before_tax"`
	TaxApplicable    bool   `json:"tax_applicable"`
	StickerID        string `json:"stiker_id"`
	StickerPackageID string `json:"sticker_package_id"`
	ItemID           int64  `json:"item_id"`
	OrderID          int64  `json:"order_id"`
	SourceContent    struct {
		OrderSN string `json:"order_sn"`
		ItemID  int64  `json:"item_id"`
	} `json:"source_content"`
}

type ChatServiceOp struct {
	client *ShopeeClient
}

func (s *ChatServiceOp) GetMessage(shopID uint64, token string, params GetMessageParamsRequest) (*GetMessageResponse, error) {
	path := "/sellerchat/get_message"

	opt := GetMessageParamsRequest{
		PageSize: params.PageSize,
	}

	resp := new(GetMessageResponse)
	err := s.client.WithShop(uint64(shopID), token).Get(path, resp, opt)
	return resp, err
}

type GetConversationParamsRequest struct {
	Direction    string `url:"direction"`
	Type         string `url:"type"`
	NextTimeNano int64  `url:"next_timestamp_nano"`
	PageSize     int    `url:"page_size"`
}

type GetConversationResponse struct {
	BaseResponse

	Response GetConversationDataResponse `json:"response"`
}

type GetConversationDataResponse struct {
	PageResult        ConversationPageResult `json:"page_result"`
	ConversationsList []Conversation         `json:"conversations"`
}

type Conversation struct {
	ConversationID       string `json:"conversation_id"`
	ToID                 int    `json:"to_id"`
	ToName               string `json:"to_name"`
	ToAvatar             string `json:"to_avatar"`
	ShopID               int    `json:"shop_id"`
	UnreadCount          int    `json:"unread_count"`
	Pinned               bool   `json:"pinned"`
	Mute                 bool   `json:"mute"`
	LastReadMessageID    string `json:"last_read_message_id"`
	LatestMessageID      string `json:"latest_message_id"`
	LatestMessageType    string `json:"latest_message_type"`
	LatestMessageContent struct {
		Text string `json:"text"`
	} `json:"latest_message_content"`
	LatestMessageFromID      int    `json:"latest_message_from_id"`
	LastMessageTimestamp     int64  `json:"last_message_timestamp"`
	LastMessageOption        int    `json:"last_message_option"`
	MaxGeneralOptionHideTime string `json:"max_general_option_hide_time"`
}

type ConversationPageResult struct {
	PageSize   int `json:"page_size"`
	NextCursor struct {
		NextMessageTimeNano string `json:"next_message_time_nano"`
		ConversationID      string `json:"conversation_id"`
	} `json:"next_cursor"`
	More bool `json:"more"`
}

func (s *ChatServiceOp) GetConversationList(shopID uint64, token string, params GetConversationParamsRequest) (*GetConversationResponse, error) {
	path := "/sellerchat/get_conversation_list"

	opt := GetConversationParamsRequest{
		PageSize:     params.PageSize,
		Direction:    params.Direction,
		Type:         params.Type,
		NextTimeNano: params.NextTimeNano,
	}

	resp := new(GetConversationResponse)
	err := s.client.WithShop(uint64(shopID), token).Get(path, resp, opt)
	return resp, err
}

type GetSendMessageResponse struct {
	BaseResponse

	Response GetSendMessageDataResponse `json:"response"`
}

type GetSendMessageDataResponse struct {
	MessageID   string `json:"message_id"`
	ToID        int    `json:"to_id"`
	MessageType string `json:"message_type"`
	Content     struct {
		Text string `json:"text"`
	} `json:"content"`
	ConversationID   int64 `json:"conversation_id"`
	CreatedTimestamp int   `json:"created_timestamp"`
	MessageOption    int   `json:"message_option"`
}

type SendMessageRequest struct {
	ToID        int64              `json:"to_id"`
	MessageType string             `json:"message_type"`
	Content     ContentSendMessage `json:"content"`
}

type ContentSendMessage struct {
	ContentMessage
	ImageURL string `json:"image_url"`
}

func (s *ChatServiceOp) SendMessage(shopID uint64, token string, request SendMessageRequest) (*GetSendMessageResponse, error) {
	path := "/sellerchat/send_message"
	resp := new(GetSendMessageResponse)
	req, err := StructToMap(request)
	if err != nil {
		return nil, err
	}

	err = s.client.WithShop(uint64(shopID), token).Post(path, req, resp)
	return resp, err
}
