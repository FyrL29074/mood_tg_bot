package storage

type Statistics struct {
	Categories map[string]Category
}

type Category struct {
	Emotions map[string]int
}
