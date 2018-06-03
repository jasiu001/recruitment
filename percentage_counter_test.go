package main

import "testing"

func TestLanguageData_GetPercentege(t *testing.T) {
	languageData := NewLanguageData()

	languageData.AddLanguageItem("langA", 1000)
	languageData.AddLanguageItem("langB", 2000)
	languageData.AddLanguageItem("langC", 5000)

	languageData.CountPercentage()

	if languageData.GetPercentege("langA") != "12.50" {
		t.Errorf(
			"Language percantage (langA) should be equal to 12.50 but %s is",
			languageData.GetPercentege("langA"))
	}

	if languageData.GetPercentege("langB") != "25.00" {
		t.Errorf(
			"Language percantage (langB) should be equal to 25.00 but %s is",
			languageData.GetPercentege("langB"))
	}

	if languageData.GetPercentege("langC") != "62.50" {
		t.Errorf(
			"Language percantage (langC) should be equal to 62.50 but %s is",
			languageData.GetPercentege("langC"))
	}
}
