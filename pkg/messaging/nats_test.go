package messaging

import "testing"

func TestPrefix(t *testing.T) {
	c := &Client{cfg: Config{SubjectPrefix: "esf"}}
	if got := c.prefixed("jobs"); got != "esf.jobs" {
		t.Fatalf("unexpected subject: %s", got)
	}
}
