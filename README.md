***

<div align="center">
    <b><em>Collection</em></b><br>
    implement java useful data structure by golang for go project
</div>

<div align="center">
</div>

## PriorityQueue

implements the JAVA collection priority queue

```go
import (
		"github.com/fzft/collection"
)

func func main() {
    pq := NewPriorityQueue(5, func(t, u interface{}) int {
        val1 := t.(int)
        val2 := u.(int)
        if val1 > val2 {
        return 1
        } else if val1 == val2 {
        return 0
        } else {
        return -1
        }
	})
    pq.Add(3) // [3]
    pq.Add(2) // [2,3]
    pq.Add(5) // [2,3,5]
    pq.Add(4) // [2,3,4,5]
	
	_ = pq.Peek() // 2
	_ = pq.Poll() // 2
}

```
## ArrayList

implements the JAVA collection ArrayList, include removeAll, addAll, retainAll method

```go
import (
		"github.com/fzft/collection"
)

func func main() {
	
	al1 := NewArrayList(0)
	al.Add(1) // [1]
    al.Add(2) // [2,3]
    al.Add(3) // [1,2,3]
    al.Add(4) // [1,2,3,4]
	
    al2 := NewArrayList(0)
    al2.Add(5) // [5]
    al2.Add(6) // [5,6]
    al2.Add(7) // [5,6,7]
    al2.Add(8) // [5,6,7,8]
	
	al1.AddAll(2, al2) // [1,2,5,6,7,8,3,4]
	
	al1.RemoveAll(al2) // [1,2,3,4]
}

```
