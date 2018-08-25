package main

func topoSort(deps map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visit func(string)

	visit = func(item string) {
		if !seen[item] {
			seen[item] = true
			for _, i := range deps[item] {
				visit(i)
			}
			order = append(order, item)
		}
	}

	for item, _ := range deps {
		visit(item)
	}

	return order
}
