/**
 * Definition for a QuadTree node.
 * type Node struct {
 *     Val bool
 *     IsLeaf bool
 *     TopLeft *Node
 *     TopRight *Node
 *     BottomLeft *Node
 *     BottomRight *Node
 * }
 */

type Boundary struct{
	from []int
	to []int
}

// Order: TL,TR,BL,BR (0,0)(3,3)
func splitGrid(grid Boundary)[]Boundary {
	from, to := grid.from, grid.to
	fromX,fromY,toX,toY := from[0],from[1],to[0],to[1]
	topLeft := Boundary{
		from: grid.from, 
		to: []int{fromX + (toX-fromX)/2, fromY + (toY-fromY)/2},
	}
	topRight := Boundary{
		from: []int{topLeft.from[0], topLeft.to[1]+1},
		to: []int{topLeft.to[0], toY},
	}
	bottomLeft := Boundary{
		from: []int{topLeft.to[0]+1, topLeft.from[1]},
		to: []int{toX, topLeft.to[1]},
	}
	bottomRight := Boundary{
		from: []int{topLeft.to[0]+1,topLeft.to[1]+1},
		to: []int{toX, toY},
	}
	return []Boundary{topLeft, topRight, bottomLeft, bottomRight}
}

func isLeaf(boundary Boundary, grid [][]int)(bool,bool){
	startRow,startCol,endRow,endCol := boundary.from[0],boundary.from[1],boundary.to[0],boundary.to[1]
	val := grid[startRow][startCol]
	for i := startRow; i <= endRow; i++ {
		for j := startCol; j <= endCol; j++ {
			if grid[i][j] != val {
				return false, false
			}
		}
	}
	if val == 0 {
		return false, true
	}
	return true, true
}

func construct(grid [][]int) *Node {
    var createTree func(boundary Boundary) *Node
	createTree = func(boundary Boundary) *Node{
		val, isLeaf := isLeaf(boundary, grid)
		if isLeaf {
			newNode := &Node{Val: val, IsLeaf: true}
			return newNode
		}
		newNode := &Node{Val: val, IsLeaf: false}
		boundaries := splitGrid(boundary)
		newNode.TopLeft = createTree(boundaries[0])
		newNode.TopRight = createTree(boundaries[1])
		newNode.BottomLeft = createTree(boundaries[2])
		newNode.BottomRight = createTree(boundaries[3])
		return newNode
	}
	n := len(grid)
	return createTree(Boundary{from: []int{0,0}, to: []int{n-1,n-1}})
}
