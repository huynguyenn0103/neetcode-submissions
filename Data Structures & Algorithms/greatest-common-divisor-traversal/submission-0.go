type UnionFind struct {
	parents []int
	ranks []int
}

func NewUnion(n int) *UnionFind{
	parents := make([]int, n)
	for node := range parents{
		parents[node]=node
	}
	return &UnionFind{
		parents: parents,
		ranks: make([]int,n),
	}
}

func (u *UnionFind) Find(v int)int{
	if u.parents[v] == v{
		return v
	} else {
		u.parents[v] = u.Find(u.parents[v])
		return u.parents[v]
	}
}

func (u *UnionFind) Union(v1,v2 int)bool{
	root1 := u.Find(v1)
	root2 := u.Find(v2)
	if root1 == root2 {return false}
	rank1 := u.ranks[root1]
	rank2 := u.ranks[root2]
	if rank1 < rank2 {
		u.parents[root1]=root2
	} else if rank1 > rank2{
		u.parents[root2]=root1
	} else {
		u.ranks[root2]++
		u.parents[root1]=root2
	}
	return true
}

func (u *UnionFind) Connected(v1,v2 int)bool{
	return u.Find(v1) == u.Find(v2)
}

func gcd(a, b int)int{
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func canTraverseAllPairs(nums []int) bool {
	n := len(nums)
	union := NewUnion(n)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++  {
			if gcd(nums[i],nums[j]) > 1{
				union.Union(i,j)
			}
		}
	}

	for i := 0; i < n; i++{
		for j := i + 1; j < n; j++{
			if !union.Connected(i,j){
				return false
			}
		}
	}
	return true
}
