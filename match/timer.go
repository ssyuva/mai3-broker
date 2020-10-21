package match

import (
	"time"
)

func (m *match) onOrderExpired(orderID string) {
	m.deleteOrderTimer(orderID)
	// TODO cancel order
}

func (m *match) setExpirationTimer(orderID string, expiresAt time.Time) error {
	now := time.Now().UTC()
	if !expiresAt.After(now) {
		go m.onOrderExpired(orderID)
		return nil
	}
	t := time.AfterFunc(expiresAt.Sub(now), func() { m.onOrderExpired(orderID) })
	if t != nil {
		m.timers[orderID] = t
	}

	return nil
}

func (m *match) stopTimers() {
	for k, t := range m.timers {
		t.Stop()
		delete(m.timers, k)
	}
}

func (m *match) deleteOrderTimer(orderID string) bool {
	if t, ok := m.timers[orderID]; ok {
		t.Stop()
		delete(m.timers, orderID)
		return true
	}
	return false
}
