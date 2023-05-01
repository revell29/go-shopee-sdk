package shopee

type ChatService interface {
	GetMessage(shopID uint64, token string, params GetMessageParamsRequest) (*GetMessageResponse, error)
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
	StickerID        int    `json:"stiker_id"`
	StickerPackageID int    `json:"sticker_package_id"`
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
