package button

import "fmt"

type WinButton struct{}

func (w *WinButton) BPaint() {
	fmt.Println("Painting Windows Button")
}
