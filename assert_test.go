package assert

import "testing"

func TestMust(t *testing.T) {
	Must(true, "failed").Panic()
	if err := Must(false, "failed due to wtf").Error(); err == nil {
		panic("failed")
	} else {
		t.Log(err)
	}
}

func TestMustf(t *testing.T) {
	Mustf(true, "failed counter=%d", 100).Panic()
	if err := Mustf(false, "failed due to wtf too,retry=%s", []byte("foo")).Error(); err == nil {
		panic("failed")
	} else {
		t.Log(err)
	}
}
