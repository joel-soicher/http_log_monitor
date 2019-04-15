package main

import "testing"

type MockAlerts struct {
	triggered   bool
	untriggered bool
}

func (a *MockAlerts) Trigger(float64) {
	a.triggered = true
}

func (a *MockAlerts) Untrigger() {
	a.untriggered = true
}

func (a *MockAlerts) DisplayAlerts() string {
	return ""
}

func addRequestsAndCheck(t *testing.T, a *Alerter, m *MockAlerts, nb int, expectedTriggered bool, expectedUntriggered bool) {
	for i := 0; i < nb; i++ {
		a.AddRequest(nil)
	}
	a.Compute()
	if m.triggered != expectedTriggered || m.untriggered != expectedUntriggered {
		t.Error()
	}
	m.triggered, m.untriggered = false, false
}

func TestAlerter(t *testing.T) {
	cfg := &Config{
		Tick:       1,
		AlertDelay: 4,
		MaxReq:     20,
	}
	mock := &MockAlerts{}

	a := NewAlerter(cfg, mock)
	addRequestsAndCheck(t, a, mock, 10, false, false) //[10, 0, 0, 0] -> avg = 10, not triggered
	addRequestsAndCheck(t, a, mock, 10, false, false) //[10, 10, 0, 0] -> avg = 10, not triggered
	addRequestsAndCheck(t, a, mock, 10, false, false) //[10, 10, 10, 0] -> avg = 10, not triggered
	addRequestsAndCheck(t, a, mock, 10, false, false) //[10, 10, 10, 10] -> avg = 10, not triggered
	addRequestsAndCheck(t, a, mock, 30, false, false) //[30, 10, 10, 10] -> avg = 15, not triggered
	addRequestsAndCheck(t, a, mock, 30, true, false)  //[30, 30, 10, 10] -> avg = 20, triggered
	addRequestsAndCheck(t, a, mock, 30, false, false) //[30, 30, 30, 10] -> avg = 25, already triggered
	addRequestsAndCheck(t, a, mock, 30, false, false) //[30, 30, 30, 30] -> avg = 30, already triggered
	addRequestsAndCheck(t, a, mock, 10, false, false) //[10, 30, 30, 30] -> avg = 25, already triggered
	addRequestsAndCheck(t, a, mock, 10, false, false) //[10, 10, 30, 30] -> avg = 20, already triggered
	addRequestsAndCheck(t, a, mock, 10, false, true)  //[10, 10, 10, 30] -> avg = 15, untriggered
	addRequestsAndCheck(t, a, mock, 10, false, false) //[10, 10, 10, 10] -> avg = 10, already untriggered
}
