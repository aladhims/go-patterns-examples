package checkbox

import "fmt"

type WinCheckbox struct{}

func (w *WinCheckbox) CPaint() {
	fmt.Println("Painting Windows Checkbox")
}
