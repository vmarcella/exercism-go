package tree

// Node represents a node in the tree
type Node struct {
	ID       int
	Children []Node
}

// Record to hold the parent node
type Record struct {
	ID     int
	Parent Node
}

// Build builds a tree out of given records
func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil
	}

	for _, node := range records {

	}
}
