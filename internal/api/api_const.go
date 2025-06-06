package api

const (
	getUpdatesMethod            = "getUpdates"
	sendMessageMethod           = "sendMessage"
	sendPhotoMethod             = "sendPhoto"
	suggetCheckEmotionText      = "Как ты себя сейчас чувствуешь?"
	moodWasSuccesfullyAddedText = "Ваша эмоция была успешная сохранена!"
	chooseYourEmotion           = "Выбери эмоцию"
)

var emotionCategories = map[string]struct{}{
	"Радость":     {},
	"Грусть":      {},
	"Злость":      {},
	"Страх":       {},
	"Спокойствие": {},
}

var emotionCategoryButtons = []replyKeyboardButton{
	{Text: "Радость"},
	{Text: "Грусть"},
	{Text: "Злость"},
	{Text: "Страх"},
	{Text: "Спокойствие"},
}

var emotions = map[string]struct{}{
	// Радость
	"Счастье":       struct{}{},
	"Удовольствие":  struct{}{},
	"Восторг":       struct{}{},
	"Гордость":      struct{}{},
	"Благодарность": struct{}{},
	"Надежда":       struct{}{},

	// Грусть
	"Разочарование": struct{}{},
	"Сожаление":     struct{}{},
	"Усталость":     struct{}{},
	"Одиночество":   struct{}{},
	"Вина":          struct{}{},
	"Печаль":        struct{}{},

	// Злость
	"Раздражение": struct{}{},
	"Возмущение":  struct{}{},
	"Зависть":     struct{}{},
	"Ненависть":   struct{}{},
	"Гнев":        struct{}{},

	// Страх
	"Растерянность": struct{}{},
	"Опасение":      struct{}{},
	"Стыд":          struct{}{},
	"Испуг":         struct{}{},
	"Тревога":       struct{}{},
	"Паника":        struct{}{},

	// Спокойствие
	"Умиротворение":   struct{}{},
	"Удовлетворение":  struct{}{},
	"Безмятежность":   struct{}{},
	"Расслабленность": struct{}{},
	"Равнодушие":      struct{}{},
}

var joyEmotionButtons = []replyKeyboardButton{
	{Text: "Счастье"},
	{Text: "Удовольствие"},
	{Text: "Восторг"},
	{Text: "Гордость"},
	{Text: "Благодарность"},
	{Text: "Надежда"},
}

var sadnessEmotionButtons = []replyKeyboardButton{
	{Text: "Разочарование"},
	{Text: "Сожаление"},
	{Text: "Усталость"},
	{Text: "Одиночество"},
	{Text: "Вина"},
	{Text: "Печаль"},
}

var angerEmotionButtons = []replyKeyboardButton{
	{Text: "Раздражение"},
	{Text: "Возмущение"},
	{Text: "Зависть"},
	{Text: "Ненависть"},
	{Text: "Гнев"},
}

var fearEmotionButtons = []replyKeyboardButton{
	{Text: "Растерянность"},
	{Text: "Опасение"},
	{Text: "Стыд"},
	{Text: "Испуг"},
	{Text: "Тревога"},
	{Text: "Паника"},
}

var calmnessEmotionButtons = []replyKeyboardButton{
	{Text: "Умиротворение"},
	{Text: "Удовлетворение"},
	{Text: "Безмятежность"},
	{Text: "Расслабленность"},
	{Text: "Равнодушие"},
}
