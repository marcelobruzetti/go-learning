package twofer

// ShareWith returns the two-fer saying for the given name.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return "One for " + name + ", one for me."
}
