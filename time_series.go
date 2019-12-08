package timeseries

import "fmt"

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

	if c.Close == 0 {
		return fmt.Errorf("close cannot be 0")
	}

	if c.High == 0 {
		return fmt.Errorf("high cannot be 0")
	}

	if c.Low == 0 {
		return fmt.Errorf("low cannot be 0")
	}

	if c.Open == 0 {
		return fmt.Errorf("open cannot be 0")
	}

	if ts.LastCandle() == nil || c.Time.After(ts.LastCandle().Time) {
		ts.candles = append(ts.candles, c)
		return nil
	}

	return fmt.Errorf("time is earlier or equal previous")
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
