package calc

import (
	"math/rand"
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	sum := Add(3, 2)
	if sum != 5 {
		t.Errorf("sum = %d; want 5", sum)
	}
}

func TestFlake(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	j := rand.Int31n(10)
	if j > 5 {
		t.Errorf("j = %d; want <= 5", j)
	}
}
