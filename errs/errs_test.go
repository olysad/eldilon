package errs

import "testing"

type Node[T string | []string] struct {
	id      int
	content string
	with    int
	style   T
}

type render struct {
	inRender    int
	cursorStart int
	cursorEnd   int
	cursor      int
	mapping     M
}

type S []string
type M map[int][2]int

var (
	TestData = [][]*Node[[]string]{
		{
			{
				id:      1,
				content: "1",
				style:   S{"line-end"},
			}, {
				id:      2,
				content: "zone x {",
			},
		},
	}
)

func TestNodePrint(t *testing.T) {
	for _, v := range TestData {

	}
}
