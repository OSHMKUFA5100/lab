package regularmatch

import (
	"regexp"
	"testing"
)

func TestChineseRegExp(t *testing.T) {
	reg, err := regexp.Compile("[\u4e00-\u9fa5]")
	if err != nil {
		t.Errorf("Failed to compile regex: %v", err)
	}
	matchStrSlice := reg.FindAllString("hello 世界", -1)

	if len(matchStrSlice) != 2 {
		t.Errorf("Expected 2 matches, got %d", len(matchStrSlice))
	}
	if matchStrSlice[0] != "世" || matchStrSlice[1] != "界" {
		t.Errorf("Expected [世, 界], got %v", matchStrSlice)
	}
}

func TestChinaMobile(t *testing.T) {
	reg, err := regexp.Compile(`^1(3\d|4[5-9]|5[0-35-9]|6[2567]|7[0-8]|8\d|9[0-35-9])\d{8}$`)
	if err != nil {
		t.Errorf("Failed to compile regex: %v", err)
	}
	// matchStrSlice := reg.FindAllString("123-12345678 1234-12345678", -1)
	// if len(matchStrSlice) != 2 {
	// 	t.Errorf("Expected 2 matches, got %d", len(matchStrSlice))
	// }
	matchStrSlice := reg.FindAllString("17623862704", -1)
	t.Log(matchStrSlice[0])
}
