package api

type Category struct {
	Name     string
	Emotions []Emotion
}

type Emotion struct {
	Name  string
	Count int
}
