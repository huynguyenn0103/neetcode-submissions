import "slices"

type People struct{
	name string
	height int
}


func sortPeople(names []string, heights []int) []string {
    listOfPpl := make([]People, 0)
	n := len(names)
	for i := 0; i < n; i++ {
		listOfPpl = append(listOfPpl, People{name: names[i], height: heights[i]})
	}
	slices.SortFunc(listOfPpl, func(p1,p2 People)int{return p2.height - p1.height})
	result := make([]string, 0)
	for _, p := range listOfPpl {
		result = append(result, p.name)
	}
	return result
}