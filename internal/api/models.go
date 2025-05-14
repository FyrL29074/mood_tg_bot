package api

// getting models
type getUpdatesResponse struct {
	Ok          bool     `json:"ok"`
	UserActions []update `json:"result"`
}

type update struct {
	Id            int            `json:"update_id"`
	MsgInfo       *msgInfo       `json:"message,omitempty"`
	CallbackQuery *callbackQuery `json:"callback_query,omitempty"`
}

type msgInfo struct {
	Chat chatInfo `json:"chat"`
	Text string   `json:"text"`
}

type chatInfo struct {
	Id int `json:"id"`
}

type callbackQuery struct {
	MsgInfo *msgInfo `json:"message"`
	Data    string   `json:"data"`
}

// sending models
type sentMessage struct {
	ChatId      int          `json:"chat_id"`
	Text        string       `json:"text"`
	ReplyMarkup *replyMarkup `json:"reply_markup,omitempty"`
}

type replyMarkup struct {
	// InlineKeyboard  [][]inlineKeyboardButton `json:"inline_keyboard,omitempty"`
	ReplyKeyboardButton [][]replyKeyboardButton `json:"keyboard,omitempty"`
	ResizeKeyboard      bool                    `json:"resize_keyboard,omitempty"`
	OneTimeKeyboard     bool                    `json:"one_time_keyboard,omitempty"`
}

type replyKeyboardButton struct {
	Text string `json:"text"`
}

// type inlineKeyboardButton struct {
// 	Text         string `json:"text"`
// 	CallbackData string `json:"callback_data"`
// }
