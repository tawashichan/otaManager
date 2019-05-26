package common

import (
	"fmt"
	"testing"
)

func TestIsContinuous(t *testing.T) {
	s1, _ := NewDateSpan(NewDate(2018, 11, 12), NewDate(2018, 11, 13))
	s2, _ := NewDateSpan(NewDate(2018, 11, 14), NewDate(2018, 11, 15))
	if !s1.IsContinuous(s2) {
		t.Errorf("error! s1: %s %s s2: %s %s", s1.StartDate, s1.EndDate, s2.StartDate, s2.EndDate)
	}
}

func TestIsContinuous2(t *testing.T) {
	s1, _ := NewDateSpan(NewDate(2018, 11, 12), NewDate(2018, 11, 12))
	s2, _ := NewDateSpan(NewDate(2018, 11, 14), NewDate(2018, 11, 15))
	if s1.IsContinuous(s2) {
		t.Errorf("error! s1: %s %s s2: %s %s", s1.StartDate, s1.EndDate, s2.StartDate, s2.EndDate)
	}
}

func TestIsContinuous3(t *testing.T) {
	s1, _ := NewDateSpan(NewDate(2018, 11, 14), NewDate(2018, 11, 15))
	s2, _ := NewDateSpan(NewDate(2018, 11, 12), NewDate(2018, 11, 13))
	if !s1.IsContinuous(s2) {
		t.Errorf("error! s1: %s %s s2: %s %s", s1.StartDate, s1.EndDate, s2.StartDate, s2.EndDate)
	}
}

func TestIsContinuous4(t *testing.T) {
	s1, _ := NewDateSpan(NewDate(2018, 11, 10), NewDate(2018, 11, 13))
	s2, _ := NewDateSpan(NewDate(2018, 11, 9), NewDate(2018, 11, 15))
	if !s1.IsContinuous(s2) {
		t.Errorf("error! s1: %s %s s2: %s %s", s1.StartDate, s1.EndDate, s2.StartDate, s2.EndDate)
	}
}

func TestDateSpan_IsOverlapping(t *testing.T) {
	s, _ := NewDateSpan(NewDate(2018, 11, 10), NewDate(2018, 11, 13))
	if !s.IsOverlapping(s) {
		t.Error("")
	}
}

func TestMerge(t *testing.T) {
	s1, _ := NewDateSpan(NewDate(2018, 11, 10), NewDate(2018, 11, 13))
	s2, _ := NewDateSpan(NewDate(2018, 11, 14), NewDate(2018, 11, 15))
	spans := DateSpans{s1, s2}.Merge()
	if spans.Len() != 1 {
		t.Error("")
	}
	if !spans[0].EndDate.IsEqual(NewDate(2018, 11, 15)) {
		fmt.Printf("%s", spans)
		t.Errorf("%s", spans[0].EndDate)
	}
}
