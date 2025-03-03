package timesync

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/log/logtest"
)

const d50milli = 50 * time.Millisecond

func TestClock_StartClock(t *testing.T) {
	tick := d50milli
	c := RealClock{}
	then := time.Now()
	ts := NewClock(c, tick, c.Now(), logtest.New(t).WithName(t.Name()))
	tk := ts.Subscribe()
	ts.StartNotifying()

	select {
	case <-tk:
		dur := time.Since(then)
		assert.True(t, tick <= dur)
	case <-time.After(10 * tick):
		assert.Fail(t, "no notification received")
	}
	ts.Close()
}

func TestClock_StartClock_BeforeEpoch(t *testing.T) {
	tick := d50milli
	tmr := RealClock{}

	waitTime := 2 * d50milli
	then := time.Now()
	ts := NewClock(tmr, tick, tmr.Now().Add(2*d50milli), logtest.New(t).WithName(t.Name()))
	tk := ts.Subscribe()
	ts.StartNotifying()

	select {
	case <-tk:
		dur := time.Since(then)
		assert.True(t, waitTime < dur)
	case <-time.After(10 * waitTime):
		assert.Fail(t, "no notification received")
	}

	ts.Close()
}

func TestClock_TickFutureGenesis(t *testing.T) {
	tmr := &RealClock{}
	ticker := NewClock(tmr, d50milli, tmr.Now().Add(2*d50milli), logtest.New(t).WithName(t.Name()))
	assert.Equal(t, types.NewLayerID(0), ticker.lastTickedLayer) // check assumption that we are on genesis = 0
	sub := ticker.Subscribe()
	ticker.StartNotifying()
	defer ticker.Close()
	x := <-sub
	assert.Equal(t, types.NewLayerID(0), x)
	x = <-sub
	assert.Equal(t, types.NewLayerID(1), x)
}

func TestClock_TickPastGenesis(t *testing.T) {
	tmr := &RealClock{}
	start := time.Now()
	ticker := NewClock(tmr, 2*d50milli, start.Add(-7*d50milli), logtest.New(t).WithName(t.Name()))
	expectedTimeToTick := d50milli // tickInterval is 100ms and the genesis tick (layer 0) was 350ms ago
	/*
		T-350 -> layer 0
		T-250 -> layer 1
		T-150 -> layer 2
		T-50  -> layer 3
		T+50  -> layer 4
	*/
	sub := ticker.Subscribe()
	ticker.StartNotifying()
	defer ticker.Close()
	x := <-sub
	duration := time.Since(start)
	assert.Equal(t, types.NewLayerID(4), x)
	assert.True(t, duration >= expectedTimeToTick, "tick happened too soon (%v)", duration)
	assert.True(t, duration < expectedTimeToTick+d50milli, "tick happened more than 50ms too late (%v)", duration)
}

func TestClock_NewClock(t *testing.T) {
	r := require.New(t)
	tmr := &RealClock{}
	ticker := NewClock(tmr, 100*time.Millisecond, tmr.Now().Add(-190*time.Millisecond), logtest.New(t).WithName(t.Name()))
	defer ticker.Close()

	r.Equal(types.NewLayerID(1), ticker.lastTickedLayer)
}

func TestClock_CloseTwice(t *testing.T) {
	ld := d50milli
	clock := NewClock(RealClock{}, ld, time.Now(), logtest.New(t).WithName(t.Name()))
	clock.StartNotifying()
	clock.Close()
	clock.Close()
}
