// Package tree implements a method to build a simple tree data structure.
package tree

import (
	"errors"
	"sort"
)

// Record represent a single record
type Record struct {
	ID     int
	Parent int
}

// Node holds tree node data
type Node struct {
	ID       int
	Children []*Node
}

// Build returns a simple tree from a give array of records
func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}
	// Sort by ID.
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	n := make(map[int]*Node, len(records))
	for i, r := range records {
		if r.ID != i || r.Parent > r.ID || r.ID > 0 && r.Parent == r.ID {
			return nil, errors.New("not in sequence or has bad parent")
		}

		n[r.ID] = &Node{ID: r.ID}
		if r.ID != 0 {
			n[r.Parent].Children = append(n[r.Parent].Children, n[r.ID])
		}
	}

	return n[0], nil
}
