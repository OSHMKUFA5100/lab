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
