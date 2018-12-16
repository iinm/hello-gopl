package display

func ExampleDisplay_cirularReference() {
	type Cycle struct {
		Value int
		Tail  *Cycle
	}
	var c Cycle
	c = Cycle{1, &c}
	Display("c", c)

	// Output:
	// Display c (display.Cycle):
	// c.Value = 1
	// (*c.Tail).Value = 1
	// (*(*c.Tail).Tail).Value = 1
	// [warn] reached maximum depth
}
