type UnionFind struct {
	parents []int
	ranks []int
}

func NewUnion(n int) *UnionFind{
	parents := make([]int, n)
	for i := 0; i < n; i++ {
		parents[i]=i
	}
	return &UnionFind{
		parents: parents,
		ranks: make([]int, n),
	}
}

func (u *UnionFind)Find(v int)int{
	if u.parents[v] == v {
		return v
	} else {
		u.parents[v] = u.Find(u.parents[v])
		return u.parents[v]
	}
}

func (u *UnionFind)Union(v1,v2 int)bool{
	root1 := u.Find(v1)
	root2 := u.Find(v2)
	if root1 == root2{
		return false
	}
	rank1 := u.ranks[root1]
	rank2 := u.ranks[root2]
	if rank1 < rank2{
		u.parents[root1]=root2
	} else if rank1 > rank2{
		u.parents[root2]=root1
	} else {
		u.ranks[root2]++
		u.parents[root1]=root2
	}
	return true
}

func (u *UnionFind) NumOfRoot()int{
	visited := make(map[int]bool)
	result := 0
	for node := range u.parents {
		root := u.Find(node)
		if visited[root]{continue}
		visited[root]=true
		result++
	}
	return result
}

func countComponents(n int, edges [][]int) int {
    cUnion := NewUnion(n)
	for _, edge := range edges {
		cUnion.Union(edge[0], edge[1])
	}
	return cUnion.NumOfRoot()
}
