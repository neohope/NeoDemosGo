package main

import (
	"fmt"
)

type Widget struct {
	X, Y int
}

type Label struct {
	Widget        // Embedding (delegation)
	Text   string // Aggregation
}

type Button struct {
	Label // Embedding (delegation)
}
func NewButton(X int, Y int, Text string) (*Button, error) {
	return &Button{Label{Widget{X, Y}, Text}}, nil
}

type ListBox struct {
	Widget          // Embedding (delegation)
	Texts  []string // Aggregation
	Index  int      // Aggregation
}

type Painter interface {
	Paint()
}

type Clicker interface {
	Click()
}

func (label Label) Paint() {
	fmt.Printf("%p:Label.Paint(%q)\n", &label, label.Text)
}

//因为这个接口可以通过 Label 的嵌入带到新的结构体，
//所以，可以在 Button 中重载这个接口方法
func (button Button) Paint() { // Override
	fmt.Printf("%p:Button.Paint(%s)\n", &button, button.Text)
}
func (button Button) Click() {
	fmt.Printf("%p:Button.Click(%s)\n", &button, button.Text)
}

func (listBox ListBox) Paint() {
	fmt.Printf("%p:ListBox.Paint(%q)\n", &listBox, listBox.Texts)
}
func (listBox ListBox) Click() {
	fmt.Printf("%p:ListBox.Click(%q)\n", &listBox, listBox.Texts)
}

func main() {
	label := Label{Widget{10, 10}, "State:"}
	label.X = 11
	label.Y = 12
	button1 := Button{Label{Widget{10, 70}, "OK"}}
	button2, _ := NewButton(50, 70, "Cancel")
	listBox := ListBox{Widget{10, 40},
		[]string{"AL", "AK", "AZ", "AR"}, 0}

	for _, painter := range []Painter{label, listBox, button1, button2} {
		painter.Paint()
	}
	fmt.Println()

	for _, widget := range []interface{}{label, listBox, button1, button2} {
		widget.(Painter).Paint()

		// 判定是否实现了Click方法
		if clicker, ok := widget.(Clicker); ok {
			clicker.Click()
		}
		fmt.Println() // print a empty line
	}
}
