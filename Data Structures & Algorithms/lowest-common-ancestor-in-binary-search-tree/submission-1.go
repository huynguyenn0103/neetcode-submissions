/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Item struct {
	freq int
	level int
}


func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	mapTraverseValueToFreq := make(map[*TreeNode]Item)
	valP, valQ := p.Val, q.Val
	
	var find func(node *TreeNode, val int, freq int, level int)
	find = func(node *TreeNode, val int, freq int, level int){
		nextLevel := level + 1
		nextFreq := mapTraverseValueToFreq[node].freq + freq
		mapTraverseValueToFreq[node] = Item{level: nextLevel, freq: nextFreq}
		if node.Val == val {
			return
		} else if node.Val > val {
			find(node.Left, val, freq, nextLevel)
		} else {
			find(node.Right, val, freq, nextLevel)
		}
	}
	find(root, valP, 1, 0)
	find(root, valQ, -1, 0)
	var result *TreeNode
	maxLevel := 0
	for key, val := range mapTraverseValueToFreq{
		if val.freq == 0 {
			if result == nil || maxLevel < val.level {
				result = key
				maxLevel = val.level
			} 
		}
	}
    return result
}
