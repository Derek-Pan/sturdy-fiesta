package week03

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type InPostList struct {
	inorder         []int
	postorder       []int
	inOrderIdxMap   map[int]int
	postOrderIdxMap map[int]int
}

func buildTree(inorder []int, postorder []int) *TreeNode {

	build := &InPostList{
		inorder:         inorder,
		postorder:       postorder,
		inOrderIdxMap:   make(map[int]int),
		postOrderIdxMap: make(map[int]int),
	}
	//维护一个值到索引的map
	for i := 0; i < len(inorder); i++ {
		build.inOrderIdxMap[inorder[i]] = i
		build.postOrderIdxMap[postorder[i]] = i
	}

	return build.build(0, len(inorder)-1, 0, len(inorder)-1)

	// 采用数组copy的方式
	// 和利用索引的方式相比,消耗的时间长很多
	// if len(inorder) == 0 || len(postorder) == 0 {
	//     return nil
	// }

	// inOrderIdxMap := make(map[int]int)
	// postOrderIdxMap := make(map[int]int)

	// for i := 0; i < len(inorder); i++ {
	//     inOrderIdxMap[inorder[i]] = i
	//     postOrderIdxMap[postorder[i]] = i
	// }

	// rootVal := postorder[len(postorder)-1]
	// root := &TreeNode{
	//     Val: rootVal,
	//     Left: nil,
	//     Right: nil,
	// }

	// leftInOrder := inorder[0:inOrderIdxMap[rootVal]]
	// rightInOrder := inorder[inOrderIdxMap[rootVal]+1:len(inorder)]
	// leftPostOrder := postorder[0:len(leftInOrder)]
	// rightPostOrder := postorder[len(leftInOrder): len(postorder)-1]

	// root.Left = buildTree(leftInOrder, leftPostOrder)
	// root.Right = buildTree(rightInOrder, rightPostOrder)

	// return root
}

func (this *InPostList) build(inL, inR, postL, postR int) *TreeNode {
	//当中序,后序的右索引小于左索引时,结束
	if inR < inL || postR < postL {
		return nil
	}

	//后序索引的尾部必然是根节点
	root := &TreeNode{
		Val:   this.postorder[postR],
		Left:  nil,
		Right: nil,
	}

	inOrderIdx := this.inOrderIdxMap[this.postorder[postR]]
	//计算左节点数量
	leftSize := inOrderIdx - inL
	//计算右节点数量
	rightSize := inR - inOrderIdx

	//遍历左子树
	root.Left = this.build(inL, inL+leftSize-1, postL, postL+leftSize-1)
	//遍历右子树
	root.Right = this.build(inR-rightSize+1, inR, postR-rightSize, postR-1)
	return root
}
