package pwm_profile

import (
	"testing"
	"time"
)

func TestRandom(t *testing.T) {
	tFormat := "15:04"
	conf := `
{
	"start":"10:00",
	"end": "19:30"
}
`
	d, err := Random([]byte(conf), 13, 100)
	if err != nil {
		t.Error(err)
	}
	t1, err := time.Parse(tFormat, "10:30")
	if err != nil {
		t.Error(err)
	}
	if int(d.Get(t1)) != 85 {
		t.Error("Expected 85, got:", d.Get(t1))
	}
}
