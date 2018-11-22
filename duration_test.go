package duration

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestFromString(t *testing.T) {
	t.Parallel()

	// test with bad format
	_, err := FromString("asdf")
	assert.Equal(t, ErrBadFormat, err)

	// test with month
	_, err = FromString("P1M")
	assert.Equal(t, ErrNoMonth, err)

	// test with good full string
	dur, err := FromString("P1Y2DT3H4M5S")
	assert.Nil(t, err)
	assert.Equal(t, 1, dur.Years)
	assert.Equal(t, 2, dur.Days)
	assert.Equal(t, 3, dur.Hours)
	assert.Equal(t, 4, dur.Minutes)
	assert.Equal(t, float32(5), dur.Seconds)

	// test with good week string
	dur, err = FromString("P1W")
	assert.Nil(t, err)
	assert.Equal(t, 1, dur.Weeks)
}

func TestString(t *testing.T) {
	t.Parallel()

	// test empty
	d := Duration{}
	assert.Equal(t, "P", d.String())

	// test only larger-than-day
	d = Duration{Years: 1, Days: 2}
	assert.Equal(t, "P1Y2D", d.String())

	// test only smaller-than-day
	d = Duration{Hours: 1, Minutes: 2, Seconds: 3}
	assert.Equal(t, "PT1H2M3S", d.String())

	// test full format
	d = Duration{Years: 1, Days: 2, Hours: 3, Minutes: 4, Seconds: 5}
	assert.Equal(t, "P1Y2DT3H4M5S", d.String())

	// test week format
	d = Duration{Weeks: 1}
	assert.Equal(t, "P1W", d.String())
}

func TestToDuration(t *testing.T) {
	t.Parallel()

	d := Duration{Years: 1}
	assert.Equal(t, time.Hour*24*365, d.ToDuration())

	d = Duration{Weeks: 1}
	assert.Equal(t, time.Hour*24*7, d.ToDuration())

	d = Duration{Days: 1}
	assert.Equal(t, time.Hour*24, d.ToDuration())

	d = Duration{Hours: 1}
	assert.Equal(t, time.Hour, d.ToDuration())

	d = Duration{Minutes: 1}
	assert.Equal(t, time.Minute, d.ToDuration())

	d = Duration{Seconds: 1}
	assert.Equal(t, time.Second, d.ToDuration())
}
