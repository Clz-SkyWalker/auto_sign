package recyle

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestEveryDayRang(t *testing.T) {
	recyle := NewRecyle(RecyleParam{
		RType: EveryDayRang,
		CTime: "15:55",
		Range: 0,
		Fun: func() {
			fmt.Println("hi")
		},
	})
	recyle.CreateRecyle()
	time.Sleep(1 * time.Second)
	assert.NotEqual(t, recyle.duration, nil, "循环周期为nil")
}
