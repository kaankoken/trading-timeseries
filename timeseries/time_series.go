package timeseries

import (
	"errors"
	"fmt"
)

var ErrUnexpectedTime = errors.New("time is earlier or equal previous")

// TimeSeries represents series of trading candles
type TimeSeries struct {
	candles []*Candle
}

// New creates TimeSeries
func New() *TimeSeries {
	ts := new(TimeSeries)
	ts.candles = make([]*Candle, 0)

	return ts
}

// AddCandle adds trading candle to series.
// Each new candle must have later time than previous
func (ts *TimeSeries) AddCandle(c *Candle) error {
	if c == nil {
		return fmt.Errorf("candle cannot be nil")
	}

	if ts.LastCandle() == nil || c.Time.After(ts.LastCandle().Time) {
		ts.candles = append(ts.candles, c)
		return nil
	}

	return ErrUnexpectedTime
}

// RemoveCandles remove the candle from series
// [startIndex] is mandatory
// [endIndex] is optional
// returns a new timeseries object
func (ts *TimeSeries) RemoveCandles(startIndex int, endIndex *int) (*TimeSeries, error) {
	if ts.Length() == 0 {
		return nil, fmt.Errorf("timeseries cannot be empty")
	}

	if startIndex < 0 {
		return nil, fmt.Errorf("startIndex cannot be negative")
	}

	if endIndex != nil && startIndex >= *endIndex {
		return nil, fmt.Errorf("endIndex should be greater than startIndex")
	}

	if endIndex != nil && len(ts.candles) < *endIndex {
		return nil, fmt.Errorf("endIndex should be less than candle size")
	}

	newTs := *ts

	if endIndex == nil {
		newTs.candles = newTs.candles[startIndex:]
	} else {
		newTs.candles = newTs.candles[startIndex:*endIndex]
	}

	return &newTs, nil
}

// LastCandle returns last candle in series
func (ts *TimeSeries) LastCandle() *Candle {
	if len(ts.candles) > 0 {
		return ts.candles[len(ts.candles)-1]
	}

	return nil
}

// Candle returns candle by index [0, len(series)-1]
func (ts *TimeSeries) Candle(index int) *Candle {
	if index >= 0 && index < len(ts.candles) {
		return ts.candles[index]
	}

	return nil
}

// Length returns length of series
func (ts *TimeSeries) Length() int {
	return len(ts.candles)
}
