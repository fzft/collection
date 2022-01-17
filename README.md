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
