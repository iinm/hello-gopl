package main

import "fmt"

func topoSort(deps map[string][]string) ([]string, error) {
	var order []string
	seen := make(map[string]bool)
	var visit func(string, []string) error

	visit = func(item string, path []string) error {
		if len(path) > 0 && contains(path, item) {
			return fmt.Errorf("circular dependency: %q", append(path, item))
		}

		if !seen[item] {
			seen[item] = true
			for _, i := range deps[item] {
				err := visit(i, append(path, item))
				if err != nil {
					return err
				}
			}
			order = append(order, item)
		}
		return nil
	}

	for item, _ := range deps {
		err := visit(item, nil)
		if err != nil {
			return nil, err
		}
	}

	return order, nil
}

func contains(items []string, item string) bool {
	for _, i := range items {
		if item == i {
			return true
		}
	}
	return false
}
