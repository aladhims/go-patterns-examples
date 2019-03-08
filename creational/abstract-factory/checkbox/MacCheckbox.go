package checkbox

import "fmt"

type MacCheckbox struct{}

func (m *MacCheckbox) CPaint() {
	fmt.Println("Painting Mac Checkbox")
}
