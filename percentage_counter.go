package main

import "fmt"

type LanguageData struct {
	languages map[string]*LanguageItem
	sum       int
}

type LanguageItem struct {
	bytes      int
	percentage string
}

// Create data struct to count percentege language statistic
func NewLanguageData() *LanguageData {
	data := make(map[string]*LanguageItem)

	return &LanguageData{
		languages: data,
		sum:       0,
	}
}

// Add item to count languages statistic for all repository
func (ld *LanguageData) AddLanguageItem(name string, bytes int) {
	ld.sum += bytes
	ld.languages[name] = &LanguageItem{
		bytes: bytes,
	}
}

// Run counting of percentege language statistic
func (ld *LanguageData) CountPercentage() {
	for _, lang := range ld.languages {
		lang.percentage = count(lang.bytes, ld.sum)
	}
}

func count(bytes int, sum int) string {
	return fmt.Sprintf("%.2f", (float64(bytes) / float64(sum) * 100))
}

// Return counted statistic for specific language
func (ld LanguageData) GetPercentege(name string) string {
	for langName, lang := range ld.languages {
		if langName == name {
			return lang.percentage
		}
	}

	return "0"
}
