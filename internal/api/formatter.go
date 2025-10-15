package api

import (
	"fmt"
	"sort"
	"strings"
)

func FormatStatistics(stat []Category) string {
	var str strings.Builder
	str.WriteString("Ваша статистика за неделю:\n\n")

	totals := make([]int, len(stat))
	for i, category := range stat {
		sum := 0
		for _, emotion := range category.Emotions {
			sum += emotion.Count
		}
		totals[i] = sum
	}
	sort.Slice(
		stat,
		func(i, j int) bool { return totals[i] > totals[j] },
	)

	for _, category := range stat {
		str.WriteString(formatCategory(category))
	}

	return str.String()
}

func formatCategory(c Category) string {
	var str strings.Builder

	counter := 0
	for _, emotion := range c.Emotions {
		counter += int(emotion.Count)
	}

	str.WriteString(fmt.Sprintf("%s - %d\n", c.Name, counter))

	sort.Slice(c.Emotions, func(i, j int) bool { return c.Emotions[i].Count > c.Emotions[j].Count })

	for _, emotion := range c.Emotions {
		str.WriteString(formatEmotion(emotion))
	}

	return str.String()
}

func formatEmotion(emotion Emotion) string {
	var str strings.Builder

	str.WriteString(fmt.Sprintf("	• %s - %d\n", emotion.Name, emotion.Count))

	return str.String()
}
