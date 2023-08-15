package stati

import (
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	s := New(&Config{
		Name:          "Duyurubu Realtime Stats",
		FetchInterval: time.Second,
	})

	if err := s.Start(3000); err != nil {
		t.Error(err)
	}
}
