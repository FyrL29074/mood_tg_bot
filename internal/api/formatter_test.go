package api

import "testing"

func TestFormatStatistics(t *testing.T) {
	firstJoyEmotion := Emotion{Name: "Счастье", Count: 2}
	secondJoyEmotion := Emotion{Name: "Удовольствие", Count: 4}
	joyCategory := Category{Name: "Радость", Emotions: []Emotion{firstJoyEmotion, secondJoyEmotion}}
	firstSadnessEmotion := Emotion{Name: "Разочарование", Count: 4}
	secondSadnessEmotion := Emotion{Name: "Сожаление", Count: 8}
	sadnessCategory := Category{Name: "Грусть", Emotions: []Emotion{firstSadnessEmotion, secondSadnessEmotion}}
	stat := []Category{joyCategory, sadnessCategory}

	got := FormatStatistics(stat)
	want := `Ваша статистика за неделю:

Грусть - 12
	• Сожаление - 8
	• Разочарование - 4
Радость - 6
	• Удовольствие - 4
	• Счастье - 2
`

	if got != want {
		t.Errorf("got != want\n got = %s\n want = %s", got, want)
	}
}

func Test_formatCategory(t *testing.T) {
	firstEmotion := Emotion{Name: "Счастье", Count: 2}
	secondEmotion := Emotion{Name: "Удовольствие", Count: 4}
	c := Category{Name: "Радость", Emotions: []Emotion{firstEmotion, secondEmotion}}

	got := formatCategory(c)
	want := `Радость - 6
	• Удовольствие - 4
	• Счастье - 2
`

	if got != want {
		t.Errorf("got != want\n got = %s\n want = %s", got, want)
	}
}

func Test_formatEmotion(t *testing.T) {
	e := Emotion{Name: "Счастье", Count: 2}

	got := formatEmotion(e)
	want := "	• Счастье - 2\n"

	if got != want {
		t.Errorf("got != want\n got = %s\n want = %s\n", got, want)
	}
}
