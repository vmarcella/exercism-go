package twofer

// ShareWith determines who to share with
func ShareWith(name string) string {
	// Determine who's going to be sharing with me
	// Return - A string formatted properly to who's sharing with me
	if len(name) == 0 {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
