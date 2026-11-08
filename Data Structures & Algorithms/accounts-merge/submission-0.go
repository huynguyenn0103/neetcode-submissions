import "slices"

type Item struct{
	email string
	account string
}

type UnionFind struct {
	parents map[Item]Item
	ranks map[Item]int
}

func compareItem(i1,i2 Item)bool {
	return i1.email == i2.email && i1.account == i2.account
}

func NewUnion() *UnionFind{
	return &UnionFind{
		parents: make(map[Item]Item),
		ranks: make(map[Item]int),
	}
}

func (u *UnionFind)Find(item Item)Item{
	if _, existed := u.parents[item]; !existed {
		u.parents[item] = item
		return u.parents[item]
	} else {
		if u.parents[item] == item {
			return u.parents[item]
		}
		u.parents[item] = u.Find(u.parents[item])
		return u.parents[item]
	}
}

func (u *UnionFind) Union(item1, item2 Item)bool {
	root1 := u.Find(item1)
	root2 := u.Find(item2)
	if compareItem(root1,root2){return false}
	rank1 := u.ranks[root1]
	rank2 := u.ranks[root2]
	if rank1 < rank2 {
		u.parents[root1]=root2
	} else if rank1 > rank2 {
		u.parents[root2]=root1
	} else {
		u.ranks[root2]++
		u.parents[root1]=root2
	}
	return true
}

func (u *UnionFind) GetResult()[][]string{
	mapRootToAllItems := make(map[Item][]string)
	for item := range u.parents{
		root := u.Find(item)
		if mapRootToAllItems[root]==nil{
			mapRootToAllItems[root]=make([]string,0)
		}
		mapRootToAllItems[root]=append(mapRootToAllItems[root],item.email)
	}
	result := make([][]string, 0)
	for root := range mapRootToAllItems{
		slices.SortFunc(mapRootToAllItems[root], func(s1,s2 string)int{
			return strings.Compare(s1,s2)
		})
		tmp := []string{root.account}
		tmp = append(tmp, mapRootToAllItems[root]...)
		result = append(result, tmp)
	}
	return result
}
func accountsMerge(accounts [][]string) [][]string {
	union := NewUnion()
	for _, groups := range accounts {
		account := groups[0]
		n := len(groups)
		if n == 2 {
			item := Item{account: account, email: groups[1]}
			union.Union(item,item)
		} else {
			for i := 2; i < n; i++ {
				curItem := Item{account: account, email: groups[i]}
				prevItem := Item{account: account, email: groups[i-1]}
				union.Union(curItem,prevItem)
			}
		}
	}
	return union.GetResult()
}
