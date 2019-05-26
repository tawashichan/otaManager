package common

import (
	"math/rand"
	"sort"
	"time"
)

// 開始と終了に前後関係が定まっている日付の組を扱います
type DateSpan struct {
	StartDate Date `json:"startDate"  sql:"not null;type:date;default:'2000-1-1'"`
	EndDate   Date `json:"endDate"  sql:"not null;type:date;default:'2000-1-1'"`
}

type DateSpans []DateSpan

type DateSpansSlice []DateSpans

const (
	ErrorInvalidDateSpan  ErrorType = "invalid_date_span"
	ErrorOutdatedDateSpan ErrorType = "outdated_date_span"
)

func NewDateSpan(StartDate Date, EndDate Date) (DateSpan, error) {
	if StartDate.IsLater(EndDate) {
		return DateSpan{}, InvalidDateSpan()
	} else {
		span := DateSpan{
			StartDate: StartDate,
			EndDate:   EndDate,
		}
		return span, nil
	}
}

func InvalidDateSpan() CommonError {
	return ErrorBadRequest("invalid date span", ErrorInvalidDateSpan)
}

func OutdatedDateSpan() CommonError {
	return ErrorBadRequest("outdated date span", ErrorOutdatedDateSpan)
}

func (s DateSpan) GetDateList() Dates {
	var dateList = Dates{}
	ey, em, ed := s.EndDate.Date()
	var currentDate = s.StartDate
	for {
		y, m, d := currentDate.Date()
		dateList = append(dateList, currentDate)
		currentDate = currentDate.PlusNDay(1)
		if ey == y && em == m && ed == d {
			break
		}
	}
	return dateList
}

func (s DateSpan) IsContinuous(other DateSpan) bool {
	if s.StartDate.IsEarlierEq(other.StartDate) {
		return s.EndDate.PlusNDay(1).IsLaterEq(other.StartDate)
	} else {
		return other.EndDate.PlusNDay(1).IsLaterEq(s.StartDate)
	}
}

func (s DateSpan) IsOverlapping(other DateSpan) bool {
	return s.StartDate.IsEarlierEq(other.EndDate) && s.EndDate.IsLaterEq(other.StartDate)
}

func (s DateSpans) Len() int {
	return len(s)
}

func (s DateSpans) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// StartDate順に並びかえ
func (s DateSpans) Less(i, j int) bool {
	return s[i].StartDate.IsEarlier(s[j].StartDate)
}

//重複除去してマージ
func (s DateSpans) Merge() DateSpans {
	result := DateSpans{}
	//StartDate順にソート
	sort.Sort(s)
	for i, span := range s {
		if i == 0 {
			result = append(result, span)
		} else {
			if result[len(result)-1].IsContinuous(span) {
				result[len(result)-1].EndDate = span.EndDate
			} else {
				result = append(result, span)
			}
		}
	}
	return result
}

func (s DateSpans) GetDateList() Dates {
	dates := Dates{}
	for _, ds := range s {
		dateList := ds.GetDateList()
		dates = append(dates, dateList...)
	}
	return dates
}

func (s DateSpansSlice) Merge() DateSpans {
	result := DateSpans{}
	for _, spans := range s {
		result = append(result, spans...)
	}
	return result.Merge()
}

//テスト用
func GenRandSpan() DateSpan {
	rand.Seed(time.Now().UnixNano())
	var startYear int
	if rand.Intn(10) > 5 {
		startYear = 2018
	} else {
		startYear = 2019
	}
	startDate := NewDate(startYear, time.Month(rand.Intn(11)+1), rand.Intn(29)+1)
	endDate := startDate.PlusNDay(rand.Intn(10) + 3)
	span, _ := NewDateSpan(startDate, endDate)
	return span
}

func GenRandSpans(num int) DateSpans {
	spans := DateSpans{}
	for i := 0; i < num; i++ {
		spans = append(spans, GenRandSpan())
	}
	return spans.Merge()
}
