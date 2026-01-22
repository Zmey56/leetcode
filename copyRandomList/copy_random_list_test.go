package copyRandomList

import "testing"

func TestCopyRandomList(t *testing.T) {
	tests := []struct {
		name string
		data [][]interface{}
	}{
		{
			name: "Стандартный пример",
			data: [][]interface{}{{7, nil}, {13, 0}, {11, 4}, {10, 2}, {1, 0}},
		},
		{
			name: "Зацикленный random (на себя)",
			data: [][]interface{}{{1, 1}, {2, 1}},
		},
		{
			name: "Пустой список",
			data: [][]interface{}{},
		},
		{
			name: "Один узел без random",
			data: [][]interface{}{{42, nil}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := buildList(tt.data)
			copiedHead := copyRandomList(head)

			if head == nil {
				if copiedHead != nil {
					t.Errorf("Ожидался nil, получен узел")
				}
				return
			}

			if !isDeepCopy(head, copiedHead) {
				t.Errorf("Тест '%s' провален: копия некорректна", tt.name)
			}
		})
	}
}

func buildList(data [][]interface{}) *Node {
	if len(data) == 0 {
		return nil
	}

	nodes := make([]*Node, len(data))
	for i := 0; i < len(data); i++ {
		nodes[i] = &Node{Val: data[i][0].(int)}
	}

	for i := 0; i < len(data); i++ {
		if i < len(data)-1 {
			nodes[i].Next = nodes[i+1]
		}
		randIdx := data[i][1]
		if randIdx != nil {
			nodes[i].Random = nodes[randIdx.(int)]
		}
	}
	return nodes[0]
}

func isDeepCopy(orig, copy *Node) bool {
	currO, currC := orig, copy

	for currO != nil && currC != nil {
		if currO.Val != currC.Val {
			return false
		}
		if currO == currC {
			return false
		}
		if currO.Random != nil {
			if currC.Random == nil || currO.Random.Val != currC.Random.Val || currO.Random == currC.Random {
				return false
			}
		} else if currC.Random != nil {
			return false
		}

		currO = currO.Next
		currC = currC.Next
	}
	return currO == nil && currC == nil
}
