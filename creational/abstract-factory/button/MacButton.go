package button

import "fmt"

type MacButton struct{}

func (m *MacButton) BPaint() {
	fmt.Println("Painting Mac Button")
}
