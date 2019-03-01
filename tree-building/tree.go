package tree

type Node struct {
	ID       int
	Children []Node
}

type Record struct {
	ID     int
	Parent Node
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil
	}

	for _, node := range records {

	}
}
