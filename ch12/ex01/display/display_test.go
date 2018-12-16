package display

func ExampleDisplay_keyIsArray() {
	m := map[[2]string]string{
		[2]string{"a", "b"}: "c",
	}
	Display("m", m)

	// Output:
	// Display m (map[[2]string]string):
	// m[[2]string {"a", "b"}] = "c"
}

func ExampleDisplay_keyIsStruct() {
	type K struct {
		a string
		b int
	}
	m := map[K]string{
		K{"a", 1}: "a1",
	}
	Display("m", m)

	// Output:
	// Display m (map[display.K]string):
	// m[display.K {a: "a", b: 1}] = "a1"
}
