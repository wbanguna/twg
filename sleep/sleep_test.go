package sleep

import (
	"testing"
	"time"
)

func TestTimeExecutable(t *testing.T) {
	time.Sleep(time.Minute)
}
