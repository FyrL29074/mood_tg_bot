package api

const (
	getUpdatesMethod            = "getUpdates"
	sendMessageMethod           = "sendMessage"
	sendPhotoMethod             = "sendPhoto"
	SuggetCheckEmotionText      = "Как ты себя сейчас чувствуешь?"
	moodWasSuccesfullyAddedText = "Ваша эмоция была успешная сохранена!"
	chooseYourEmotion           = "Выбери эмоцию"
	backSymbol                  = "←"
)

var emotionCategories = map[string]struct{}{
	"Радость":     {},
	"Грусть":      {},
	"Злость":      {},
	"Страх":       {},
	"Спокойствие": {},
}

var emotionCategoryButtons = [][]inlineKeyboardButton{
	{{Text: "Радость", CallbackData: "Радость"}},
	{{Text: "Грусть", CallbackData: "Грусть"}},
	{{Text: "Злость", CallbackData: "Злость"}},
	{{Text: "Страх", CallbackData: "Страх"}},
	{{Text: "Спокойствие", CallbackData: "Спокойствие"}},
}

var emotionsCategoriesMap = map[string]string{
	// Радость
	"Счастье":       "Радость",
	"Удовольствие":  "Радость",
	"Восторг":       "Радость",
	"Гордость":      "Радость",
	"Благодарность": "Радость",
	"Надежда":       "Радость",

	// Грусть
	"Разочарование": "Грусть",
	"Сожаление":     "Грусть",
	"Усталость":     "Грусть",
	"Одиночество":   "Грусть",
	"Вина":          "Грусть",
	"Печаль":        "Грусть",

	// Злость
	"Раздражение": "Злость",
	"Возмущение":  "Злость",
	"Зависть":     "Злость",
	"Ненависть":   "Злость",
	"Гнев":        "Злость",

	// Страх
	"Растерянность": "Страх",
	"Опасение":      "Страх",
	"Стыд":          "Страх",
	"Испуг":         "Страх",
	"Тревога":       "Страх",
	"Паника":        "Страх",

	// Спокойствие
	"Умиротворение":   "Спокойствие",
	"Удовлетворение":  "Спокойствие",
	"Безмятежность":   "Спокойствие",
	"Расслабленность": "Спокойствие",
	"Равнодушие":      "Спокойствие",
}

var joyEmotionButtons = [][]inlineKeyboardButton{
	{{Text: "Счастье", CallbackData: "Счастье"}},
	{{Text: "Удовольствие", CallbackData: "Удовольствие"}},
	{{Text: "Восторг", CallbackData: "Восторг"}},
	{{Text: "Гордость", CallbackData: "Гордость"}},
	{{Text: "Благодарность", CallbackData: "Благодарность"}},
	{{Text: "Надежда", CallbackData: "Надежда"}},
	{{Text: backSymbol, CallbackData: backSymbol}},
}

var sadnessEmotionButtons = [][]inlineKeyboardButton{
	{{Text: "Разочарование", CallbackData: "Разочарование"}},
	{{Text: "Сожаление", CallbackData: "Сожаление"}},
	{{Text: "Усталость", CallbackData: "Усталость"}},
	{{Text: "Одиночество", CallbackData: "Одиночество"}},
	{{Text: "Вина", CallbackData: "Вина"}},
	{{Text: "Печаль", CallbackData: "Печаль"}},
	{{Text: backSymbol, CallbackData: backSymbol}},
}

var angerEmotionButtons = [][]inlineKeyboardButton{
	{{Text: "Раздражение", CallbackData: "Раздражение"}},
	{{Text: "Возмущение", CallbackData: "Возмущение"}},
	{{Text: "Зависть", CallbackData: "Зависть"}},
	{{Text: "Ненависть", CallbackData: "Ненависть"}},
	{{Text: "Гнев", CallbackData: "Гнев"}},
	{{Text: backSymbol, CallbackData: backSymbol}},
}

var fearEmotionButtons = [][]inlineKeyboardButton{
	{{Text: "Растерянность", CallbackData: "Растерянность"}},
	{{Text: "Опасение", CallbackData: "Опасение"}},
	{{Text: "Стыд", CallbackData: "Стыд"}},
	{{Text: "Испуг", CallbackData: "Испуг"}},
	{{Text: "Тревога", CallbackData: "Тревога"}},
	{{Text: "Паника", CallbackData: "Паника"}},
	{{Text: backSymbol, CallbackData: backSymbol}},
}

var calmnessEmotionButtons = [][]inlineKeyboardButton{
	{{Text: "Умиротворение", CallbackData: "Умиротворение"}},
	{{Text: "Удовлетворение", CallbackData: "Удовлетворение"}},
	{{Text: "Безмятежность", CallbackData: "Безмятежность"}},
	{{Text: "Расслабленность", CallbackData: "Расслабленность"}},
	{{Text: "Равнодушие", CallbackData: "Равнодушие"}},
	{{Text: backSymbol, CallbackData: backSymbol}},
}
