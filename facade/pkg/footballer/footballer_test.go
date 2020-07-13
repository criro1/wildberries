package footballer

import (
	"testing"
)

func TestFootballer(t *testing.T) {
	f := NewFootballer("Varan", "Hazard", "Bale")

	expectedQty := 3
	
	resultQty, err := f.GetQty()
	if err != nil {
		t.Error("Error:", err)
	}
	
	if expectedQty != resultQty {
		t.Error("Error, expectedQty != resultQty")
	}

	expectedChoose := "1 - Varan skips and don't touch the ball\n2 - Hazard skips, but touchs and rolls the ball\n3 - Bale kicks the ball\n"

	resultChoose := ""
	for i := 0; i < 3; i++ {
		r, err := f.Choose(i, resultQty)
		if err != nil {
			t.Error("Error:", err)
		}
		resultChoose += r + "\n"
	}

	if expectedChoose != resultChoose {
		t.Error("Error, expectedChoose != resultChoose")
	}
}
