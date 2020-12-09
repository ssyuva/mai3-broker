package match

import "time"

func (m *match) checkOrdersMargin() {
	for {
		select {
		case <-m.ctx.Done():
			return
		case <-time.After(10 * time.Second):
			// TODO
			// check margin
			// check gas
		}
	}
}
