package debounce

import (
	"sync"
	"time"
)

// Debouncer is used to debounce function calls using a delay
type Debouncer struct {
	mu    sync.Mutex
	timer *time.Timer
}

// Debounce calls the function after the delay. If debounce is
// called again before the delay it will not call it ever and
// prepares to call the newly given function instead.
func (d *Debouncer) Debounce(delay time.Duration, f func()) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.timer != nil {
		d.timer.Stop()
	}
	d.timer = time.AfterFunc(delay, f)
}

// Cancel will remove the pending function call of the debouncer.
// It does nothing if there is no call pending.
func (d *Debouncer) Cancel() {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.timer != nil {
		d.timer.Stop()
	}
}

