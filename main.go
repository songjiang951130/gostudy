package main

import (
	"fmt"
	"github.com/songjiang951130/gostudy/Data/TreeNode"
)


func main() {
	fmt.Println("hello")
	root := new(TreeNode)
	root.val := 5
}


/**
* Given a binary tree and a sum, determine if the tree has a root-to-leaf path
* such that adding up all the values along the path equals the given sum.
*/
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false;
	}
	sum = sum - root.Val
	if sum == 0 && root.Left ==nil && root.Right == nil{
		return true
	}
	return hasPathSum(root.Left,sum)||hasPathSum(root.Right,sum)
}

/**
算上一个 有多少条路径和等于指定值的
*/
func pathSum(root *TreeNode, sum int) int {
    if root == nil {
		return 0;
	}
	sum = sum - root.Val
	rootSum :=0
	if sum == 0 && root.Left == nil && root.Right == nil{
		rootSum = 1 
	}
	return rootSum + int32(hasPathSum(root.Left,sum))+int32(hasPathSum(root.Right,sum))
}