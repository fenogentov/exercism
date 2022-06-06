package kindergarten

import (
	"errors"
	"regexp"
	"sort"
	"strings"
)

var IsLower = regexp.MustCompile(`[a-z]`).MatchString

// Define the Garden type here.
type Garden struct {
	children []string
	diagram  string
}

// The diagram argument starts each row with a '\n'.  This allows Go's
// raw string literals to present diagrams in source code nicely as two
// rows flush left, for example,
//
//     diagram := `
//     VVCCGG
//     VVCCGG`

func NewGarden(diagram string, children []string) (*Garden, error) {
	if err := validateGarden(diagram, children); err != nil {
		return nil, err
	}

	return &Garden{children: children, diagram: diagram}, nil
}

// Plants returns list of plants grown by child
func (g *Garden) Plants(child string) ([]string, bool) {
	sort.Strings(g.children)
	idx := g.NumChild(child)
	if idx == -1 {
		return nil, false
	}

	pls := []string{}
	rows := strings.Split(g.diagram, "\n")
	pls = append(pls, decodingPlant(rows[1][idx*2]))
	pls = append(pls, decodingPlant(rows[1][idx*2+1]))
	pls = append(pls, decodingPlant(rows[2][idx*2]))
	pls = append(pls, decodingPlant(rows[2][idx*2+1]))

	return pls, true
}

// NumChild returns index of child in list.
func (g *Garden) NumChild(child string) int {
	idx := -1
	for i := 0; i < len(g.children); i++ {
		if g.children[i] == child {
			idx = i
			break
		}
	}
	return idx
}

// decodingPlant returns the name of the plant by letter.
func decodingPlant(s byte) string {
	switch s {
	case 'G':
		return "grass"
	case 'C':
		return "clover"
	case 'R':
		return "radishes"
	case 'V':
		return "violets"
	default:
		return ""
	}
}

// validate Duplicate Element checks data to create Garden structure.
func validateGarden(diagram string, children []string) error {
	if IsLower(diagram) {
		return errors.New("error")
	}
	if !validateDuplicateElement(children) {
		return errors.New("error")
	}

	d := strings.Split(diagram, "\n")
	ch := len(children)
	if len(d) != 3 || len(d[1]) != ch*2 || len(d[2]) != ch*2 {
		return errors.New("error")
	}

	return nil
}

// validateDuplicateElement checks for duplicate elements from slice.
func validateDuplicateElement(sl []string) bool {
	tmp := make(map[string]struct{})
	for _, item := range sl {
		tmp[item] = struct{}{}
	}
	return len(sl) == len(tmp)
}
