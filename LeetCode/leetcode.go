package LeetCode

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"math/rand"
	"os"
	"reflect"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"time"
)


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return helper(root, p, q)
}
func helper(root, p, q *TreeNode)*TreeNode {
	if root==nil||root==p||root==q{
		return root
	}
	left:=helper(root,p,q)
	right:=helper(root,p,q)
	if left!=nil||right!=nil{

	}
	return nil
}
type ListNode struct {
	Val int
	Next *ListNode
}
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	newhead:=new(ListNode)
	cur:=newhead
	for ;l1!=nil&&l2!=nil;{
		if l1.Val<=l2.Val{
			cur.Next=l1
			l1=l1.Next
		}else {
			cur.Next=l2
			l2=l2.Next
		}
		cur=cur.Next
	}
	for ;l1!=nil;{
		cur.Next=l1
		l1=l1.Next
		cur=cur.Next
	}
	for ;l2!=nil;{
		cur.Next=l2
		l2=l2.Next
		cur=cur.Next

	}
	return newhead.Next
}

//func sortList(head *ListNode) *ListNode {
//	return sort(head, nil)
//}
//
//func sort(head,tail *ListNode)*ListNode{
//	if head==nil{
//		return head
//	}
//	if head.Next==tail{
//		head.Next=nil
//		return head
//	}
//	slow,fast:=head,head
//	for fast!=tail{
//		slow=slow.Next
//		fast=fast.Next
//		if fast!=tail{
//			fast=fast.Next
//		}
//	}
//	mid:=slow
//	return mergeTwoLists(sort(head,mid),sort(mid,tail))
//}

func reverseList(head *ListNode) *ListNode {
	if head==nil||head.Next==nil{
		return head
	}
	cur :=head
	var newhead *ListNode

	for cur!=nil{
		next:=cur.Next
		cur.Next=newhead
		newhead=cur
		cur=next
	}
	return newhead
}
func mergeKLists2(lists []*ListNode,l,r int) *ListNode {
	if l==r{
		return lists[l]
	}
	if l>r{
		return nil
	}
	mid:=(l+r)>>1
	return mergeTwoLists(mergeKLists2(lists,l,mid),mergeKLists2(lists,mid+1,r))
}
func mergeKLists(lists []*ListNode) *ListNode {
	return mergeKLists2(lists,0,len(lists)-1)
}
func firstMissingPositive(nums []int) int {
	n:=len(nums)
	for i:=0;i<n;i++{
		for nums[i]>0&&nums[i]<=n&&nums[i]!=nums[nums[i]-1]{
			tmp:=nums[i]
			nums[i]=nums[nums[i]-1]
			nums[tmp-1]=tmp
		}
	}
	for i:=0;i<n;i++{
		if nums[i]!=i+1{
			return i+1
		}
	}
	return n+1
}
func hasCycle(head *ListNode) bool {
	if head==nil||head.Next==nil{
		return false
	}
	slow,fast:=head,head.Next
	for slow!=fast{
		if slow==nil||fast==nil{
			return false
		}
		slow=slow.Next
		fast=fast.Next.Next
	}
	return true

}
func search(nums []int, target int) int {
	n:=len(nums)
	l:=0
	r:=n-1
	for l<=r{
		mid :=(l+r)/2
		if nums[mid]==target{
			return mid
		}
		if nums[mid]>=nums[0]{
			if target>=nums[0]&&target<nums[mid]{
				r=mid-1
			}else {
				l=l+1
			}

		}else {
			if target>nums[mid]&&target<nums[r]{
				l=mid+1
			}else {
				r=mid-1
			}
		}
	}
	return -1
}
func swapPairs(head *ListNode) *ListNode {
	dummyHead:=&ListNode{0,head}
	temp:=dummyHead
	for temp.Next!=nil&&temp.Next.Next!=nil{
		node1:=temp.Next
		node2:=temp.Next.Next
		temp.Next=node2
		node1.Next=node2.Next
		node2.Next=node1
		temp=node1
	}
	return dummyHead.Next
}
func nextPermutation(nums []int)  {
	n:=len(nums)
	i:=n-2
	for i>=0&&nums[i]>=nums[i+1]{
		i--
	}
	if i>=0{
		j:=n-1
		for j>=0&&nums[i]>=nums[j]{
			j--
		}
		nums[i],nums[j]=nums[j],nums[i]
	}
	reverse(nums[i+1:])
}
func reverse(nums []int){
	for i,n:=0,len(nums);i<n/2;i++{
		nums[i],nums[n-1-i]=nums[n-1-i],nums[i]
	}
}
func isValidBST2(root *TreeNode,lower,upper int) bool {
	if root==nil{
		return true
	}
	if root.Val<=lower||root.Val>=upper{
		return false
	}
	return isValidBST2(root.Left,lower,root.Val) && isValidBST2(root.Right,root.Val,upper)
}
func isValidBST(root *TreeNode) bool {
	return isValidBST2(root,math.MinInt64,math.MaxInt64)
}
func isValid(s string) bool {
	n:=len(s)
	if n % 2 == 1 {
		return false
	}
	pairs:=map[byte]byte{
		')':'(',
		']':'[',
		'}':'{',
	}
	stack:=[]byte{}
	for i:=0;i<n;i++{
		if pairs[s[i]]>0{
			if len(stack)==0||stack[len(stack)-1]!=pairs[s[i]]{
				return false
			}
			stack=stack[:len(stack)-1]
		}else{
			stack=append(stack, s[i])
		}
	}
	return len(stack)==0
}
func backtrack(first int,nums []int,res *[][]int,n int){
	if first==n{
		fmt.Println(nums)
		tmp:=make([]int,len(nums))
		copy(tmp,nums)
		*res=append(*res,tmp)
		fmt.Println(*res)
		return
	}
	for i:=first;i<n;i++{
		( nums)[i],( nums)[first]=( nums)[first],( nums)[i]
		backtrack(first+1,nums,res,n)
		( nums)[i],( nums)[first]=( nums)[first],( nums)[i]
	}
}
func permute(nums []int) [][]int {
	res:=make([][]int,0)
	backtrack(0,nums,&res,len(nums))
	return res
}
func change(amount int, coins []int) int {
	dp:=make([]int,amount+1)
	dp[0]=1
	for _,coin:=range coins{
		for i:=coin;i<amount+1;i++{
			dp[i]+=dp[i-coin]
		}
	}
	return dp[amount]
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	newhead :=&ListNode{
		0,
		nil,
	}
	cur:=newhead
	jin:=0
	for l1!=nil&&l2!=nil{
		sum:=l1.Val+l2.Val+jin
		cur.Next=&ListNode{
			sum%10,
			nil,
		}
		jin+=sum/10
		cur=cur.Next
		l1=l1.Next
		l2=l2.Next
	}
	for l1!=nil{
		sum:=l1.Val+jin
		cur.Next=&ListNode{
			sum%10,
			nil,
		}
		cur=cur.Next
		l1=l1.Next
		jin+=sum/10

	}
	for l2!=nil{
		sum:=l2.Val+jin
		jin+=sum/10
		cur.Next=&ListNode{
			sum%10,
			nil,
		}
		cur=cur.Next
		l2=l2.Next
		jin+=sum/10
	}
	if jin!=0{
		cur.Next=&ListNode{
			jin,
			nil,
		}
		cur=cur.Next

	}

	return newhead.Next
}
//func check(root1 *TreeNode,root2 *TreeNode) bool {
//	if root1==nil&&root2==nil{
//		return true
//	}
//	if root1==nil||root2==nil{
//		return false
//	}
//	return root1.Val==root2.Val&&check(root1.Left,root2.Right)&&check(root1.Right,root2.Left)
//}
//func isSymmetric(root *TreeNode) bool {
//	return check(root,root)
//}
func rand7() int {
	return 0
}
func rand10() int {
	xx:=100
	for {
		row:=rand7()
		clo:=rand7()
		xx:=(row-1)*7+clo
		if xx<=40{
			break
		}
	}
	return 1+xx%10
}
type mytreenode struct {
	P *TreeNode
	I int
}
func isCompleteTree(root *TreeNode) bool {
	if root==nil{
		return true
	}

	q:=make([]mytreenode,0)
	q=append(q,mytreenode{root,1})
	//max:=1
	for i:=0;i<len(q);i++{
		tmp:=q[i]
		//q:=q[1:]
		left:=tmp.P.Left
		if left!=nil{
			q=append(q,mytreenode{left,2*tmp.I})
			//max=2*tmp.I
		}
		right:=tmp.P.Right

		if right!=nil{
			q=append(q,mytreenode{right,2*tmp.I+1})
			//max=2*tmp.I+1
		}
	}
	return q[len(q)-1].I==len(q)

}
func reverseLinkedList(head *ListNode) {
	pre:=&ListNode{-1,nil}
	cur:=head
	for cur!=nil{
		next:=cur.Next
		cur.Next=pre
		pre=cur
		cur=next
	}
	return
}
func reverseBetween(head *ListNode, left, right int) *ListNode {
	dummyNode:=&ListNode{-1,nil}
	dummyNode.Next=head
	pre:=dummyNode
	for i:=0;i<left-1;i++{
		pre=pre.Next
	}
	rightNode:=pre
	for i:=0;i<right-left+1;i++{
		rightNode=rightNode.Next
	}
	leftNode:=pre.Next
	curr:=rightNode.Next
	pre.Next=nil
	rightNode.Next=nil
	reverseLinkedList(leftNode)
	pre.Next=rightNode
	leftNode.Next=curr
	return dummyNode.Next
}
func findPeakElement(nums []int) int {
	l:=0
	r:=len(nums)-1
	for l<r{
		mid:=(l+r)/2
		if nums[mid]>nums[mid+1]{
			r=mid
		}else {
			l=mid+1
		}
	}
	return l
}
func flatten(root *TreeNode)  {
	curr:=root
	for curr!=nil{
		if curr.Left!=nil{
			next:=curr.Left
			predecessor:=next
			for predecessor.Right!=nil{
				predecessor=predecessor.Right
			}
			predecessor.Right=curr.Right
			curr.Left,curr.Right=nil,next
		}
		curr=curr.Right
	}
}
type mynum [][]int
func (num mynum) Len() int{
	return len(num)
}
func (num mynum) Swap(i,j int) {
	num[i],num[j]=num[j],num[i]
}
func (num mynum) Less(i,j int) bool{
	return num[j][0]>num[i][0]
}
//func merge(intervals [][]int) [][]int {
//	//sort.Sort(mynum(intervals))
//	sort.SliceStable(intervals,func(i,j int) bool {return intervals[i][0]<intervals[j][0]})
//	res:=make([][]int,0)
//	for i:=0;i<len(intervals);i++{
//		new_:=make([]int,len(intervals[i]))
//		copy(new_ ,intervals[i])
//		for i<len(intervals)&&new_[1]>=intervals[i][0]{
//			new_[0]=int(math.Min(float64(new_[1]),float64(intervals[i][0])))
//			new_[1]=int(math.Max(float64(new_[1]),float64(intervals[i][1])))
//			i++
//		}
//		res=append(res,new_)
//	}
//	return res
//}
//func rob(nums []int) int {
//	if nums==nil{
//		return 0
//	}
//	dp:=make([]int,len(nums))
//	dp[0]=nums[0]
//	dp[1]=int(math.Max(float64(nums[0]),float64(nums[1])))
//	n:=len(nums)
//	for i:=2;i<n;i++{
//		dp[i]=int(math.Max(float64(dp[i-1]),float64(dp[i-2]+nums[])))
//	}
//	return dp[n-1]
//}
func maxProfit(prices []int) int {
	if len(prices)<=1{
		return 0
	}
	res:=0
	for i:=1;i<len(prices);i++{
		res+=int(math.Max(float64(0),float64(prices[i]-prices[i-1])))
	}
	return res
}
func dailyTemperatures(T []int) []int {
	n:=len(T)
	ans:=make([]int,n)
	stack:=make([]int,0)
	for i:=0;i<n;i++{
		t:=T[i]
		for len(stack)>0 && t>T[stack[len(stack)-1]]{
			previndex:=stack[len(stack)-1]
			stack=stack[:len(stack)-1]
			ans[previndex]=i-previndex
		}
		stack=append(stack,i)
	}
	return ans
}
func maxSubArray(nums []int) int {
	res:=nums[0]
	sum:=0
	for i:=0;i<len(nums);i++{
		if sum+nums[i]>=0{
			sum+=nums[i]
			res=int(math.Max(float64(sum),float64(res)))
		}else{
			sum=0
			res=int(math.Max(float64(nums[i]),float64(res)))
		}
	}
	return res
}
type pair struct {
	X int
	Y int
}
var directions=[]pair{
	{-1,0},
	{1,0},
	{0,-1},
	{0,1},
}
func exist(board [][]byte, word string) bool {
	if word==""{
		return true
	}
	if board==nil{
		return false
	}
	h,w:=len(board),len(board[0])
	vis:=make([][]bool,h)
	for i:=range vis{
		vis[i]=make([]bool,w)
	}
	var check func(i,j,k int)(bool)
	check=func(i,j,k int)(bool){
		if board[i][j]!=word[k]{
			return false
		}
		if k==len(word)-1{
			return true
		}
		vis[i][j]=true
		defer func() {vis[i][j]=false}()
		for _,dir :=range directions{
			if newx,newy:=i+dir.X,j+dir.Y;0<=newx&&newx<h&&0<=newy&&newy<w&&!vis[newx][newy]{
				if check(newx, newy, k+1){
					return true
				}
			}
		}
		return false
	}
	for i,row:=range board{
		for j:=range row{
			if check(i,j,0){
				return true
			}
		}
	}
	return false
}
func pathSum2(root *TreeNode, targetSum int,res *[][]int,sum int,ans []int)  {
	if root.Left==nil&&root.Right==nil{
		if sum==targetSum{
			tmp:=make([]int,len(ans))
			copy(tmp,ans)
			*res=append(*res,tmp)
		}
		return
	}
	if root.Left!=nil{
		pathSum2(root.Left,targetSum,res,sum+root.Left.Val,	append(ans,root.Left.Val))

	}
	if root.Right!=nil{
		pathSum2(root.Right,targetSum,res,sum+root.Right.Val,append(ans,root.Right.Val))
	}

}
func pathSum(root *TreeNode, targetSum int) [][]int {
	res:=make([][]int,0)
	if root==nil{
		return res
	}
	pathSum2(root,targetSum,&res,root.Val,[]int{root.Val})
	return res
}
type mytreenode2 struct{
	T *TreeNode
	I int
}
func widthOfBinaryTree(root *TreeNode) int {
	q:=make([]mytreenode2,0)
	q=append(q,mytreenode2{root,1})
	res:=1
	for len(q)>0{
		q2:=q
		q=make([]mytreenode2,0)
		for _,tmp :=range q2{
			if tmp.T.Left!=nil{
				q=append(q,mytreenode2{tmp.T.Left,tmp.I*2})
			}
			if tmp.T.Right!=nil{
				q=append(q,mytreenode2{tmp.T.Right,tmp.I*2+1})
			}
		}
		if q==nil{
			res=int(math.Max(float64(q[len(q)-1].I-q[0].I+1),float64(res)))
		}
	}
	return res
}
func lengthOfLIS(nums []int) int {
	n:=len(nums)
	if n<1{
		return 0
	}
	if n<2{
		return 1
	}
	dp:=make([][]int,2)
	dp[0]=make([]int,n)
	dp[1]=make([]int,n)
	dp[0][0]=0
	dp[1][0]=1
	res:=1
	for i:=1;i<n;i++{
		dp[0][i]=int(math.Max(float64(dp[0][i-1]),float64(dp[1][i-1])))
		dp[1][i]=1
		for j:=0;j<i;j++{
			if nums[j]<nums[i]{
				dp[1][i]=int(math.Max(float64(dp[1][i]),float64(dp[1][j]+1)))
			}
		}
		res=int(math.Max(float64(dp[0][i]),float64(res)))
		res=int(math.Max(float64(dp[1][i]),float64(res)))
	}
	return res
}
func minWindow(s string, t string) string {
	ori, cnt := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(t); i++ {
		ori[t[i]]++
	}

	sLen := len(s)
	len := math.MaxInt32
	ansL, ansR := -1, -1

	check := func() bool {
		for k, v := range ori {
			if cnt[k] < v {
				return false
			}
		}
		return true
	}
	for l, r := 0, 0; r < sLen; r++ {
		if r < sLen && ori[s[r]] > 0 {
			cnt[s[r]]++
		}
		for check() && l <= r {
			if (r - l + 1 < len) {
				len = r - l + 1
				ansL, ansR = l, l + len
			}
			if ori[s[l]]>0 {
				cnt[s[l]]--
			}
			l++
		}
	}
	if ansL == -1 {
		return ""
	}
	return s[ansL:ansR]

}
func diameterOfBinaryTree2(root *TreeNode,res *int) int {
	if root==nil{
		return 0
	}
	leftnum:=0
	rightnum:=0
	if root.Left!=nil{
		leftnum=diameterOfBinaryTree2(root.Left,res)
	}
	if root.Right!=nil{
		rightnum=diameterOfBinaryTree2(root.Right,res)
	}
	*res=int(math.Max(float64(*res),float64(leftnum+rightnum+1)))
	return int(math.Max(float64(leftnum+1),float64(rightnum+1)))
}
func diameterOfBinaryTree(root *TreeNode) int {
	if root==nil{
		return 0
	}
	res:=0
	diameterOfBinaryTree2(root,&res)
	return res-1
}
func numIslands(grid [][]byte) int {
	m:=len(grid)
	n:=len(grid[0])
	vis:=make([][]int,m)
	for i:=0;i<m;i++{
		vis[i]=make([]int,n)
	}
	var expand func(i,j int)
	expand=func(i,j int){
		for _,dir:= range directions{
			newi,newj:=dir.X+i,dir.Y+j
			if 0<=newi && newi<m && newj>=0 && newj<n && vis[newi][newj]==0&& grid[newi][newj]=='1'{
				vis[newi][newj]=1
				expand(newi,newj)
			}
		}
	}
	res:=0
	for i:=0;i<m;i++{
		for j:=0;j<n;j++{
			if vis[i][j]==0&& grid[i][j]=='1'{
				expand(i,j)
				res+=1
			}
		}
	}
	return res
}
func coinChange(coins []int, amount int) int {
	dp:=make([]int,amount+1)
	dp[0]=0
	for i:=1;i<=amount;i++{
		dp[i]=math.MaxInt32
	}
	//res:=math.MaxInt32
	for i:=1;i<=amount;i++{
		for _,coin :=range coins{
			if i>=coin{
				dp[i]=int(math.Min(float64(dp[i]),float64(dp[i-coin]+1)))
			}
		}
	}
	if dp[amount]>amount{
		return -1
	}
	return dp[amount]
}
func subarraySum(nums []int, k int) int {
	m:=make(map[int]int)
	m[0]=1
	n:=len(nums)
	pre:=0
	cnt:=0
	for i:=0;i<n;i++{
		pre+=nums[i]
		if m[pre-k]>0{
			cnt+=m[pre-k]
		}
		m[pre]++
	}
	return cnt
}
func detectCycle(head *ListNode) *ListNode {
	dummy:=&ListNode{-1,nil}
	dummy.Next=head
	slow,fast:=head,head.Next
	for fast!=slow{
		fast=fast.Next
		slow=slow.Next
		if fast!=nil{
			fast=fast.Next
		}
	}
	slow=dummy
	for slow!=fast{
		slow=slow.Next
		fast=fast.Next
	}
	return slow
}
func longestConsecutive(nums []int) int {
	m:=make(map[int]int)
	n:=len(nums)
	res:=0
	for i:=0;i<n;i++{
		m[nums[i]]=1
	}
	for k,_ :=range m{
		if _,ok:=m[k-1];!ok{
			cur:=k
			cutcnt:=1
			for m[cur+1]>0{
				cutcnt++
				cur++
			}
			if res<cutcnt{
				res=cutcnt
			}
		}
	}
	return res
}
func maximumSwap(num int) int {
	nums:=make([]int,0)
	tmp:=num
	for tmp/10>0{
		nums=append(nums,tmp%10)
		tmp/=10
	}
	nums=append(nums,tmp%10)
	n:=len(nums)
	m:=make(map[int][]int)
	for i:=0;i<n;i++{
		m[nums[i]]=append(m[nums[i]],i)
	}
	for i:=n-1;i>=0;i--{
		for j:=9;j>nums[i]&&j>=0;j--{
			if v,ok:=m[j];ok {
				for k := range (v) {
					if v[k] < i {
						nums[i], nums[v[k]] = nums[v[k]], nums[i]
						goto label
					}
				}
			}
		}
	}
label:res:=0
	for i:=n-1;i>=0;i--{
		res+=nums[i]
		res*=10
	}
	return res/10
}
func searchMatrix(matrix [][]int, target int) bool {
	m:=len(matrix)
	if m==0{
		return false
	}
	n:=len(matrix[0])
	if n==0{
		return false
	}
	startx:=m-1
	starty:=0
	for startx>=0&&starty<n{
		if matrix[startx][starty]>target{
			startx--
		}else if  matrix[startx][starty]<target{
			starty++
		}else {
			return true
		}
	}
	return false
}
func findDuplicate(nums []int) int {
	n:=len(nums)
	m:=make(map[int]int)
	for i:=0;i<n;i++{
		m[nums[i]]++
		if m[nums[i]]>0{
			return nums[i]
		}
	}
	return 0
}
func rotate(matrix [][]int)  {
	n:=len(matrix)
	if n==0{
		return
	}
	for i:=0;i<n/2;i++{
		for j:=0;j<(n+1)/2;j++{
			matrix[i][j],matrix[j][n-1-i],matrix[n-1-i][n-1-j],matrix[j][n-1-i]=
				matrix[j][n-1-i],matrix[n-1-i][n-1-j],matrix[j][n-1-i],matrix[i][j]
		}

	}

}
func reverseWords2(s []byte) {
	n:=len(s)
	for i:=0;i<n/2;i++{
		s[i],s[n-1-i]=s[n-1-i],s[i]
	}
	//fmt.Println(string(s))
}
func reverseWords(s string) string {
	//fmt.Println(string(s))

	s1:=[]byte(s)
	for i:=1;i<len(s1);i++{
		if s1[i]==' '&&s1[i-1]==' '{
			s1=append(append([]byte{},s1[:i]...),s1[(i+1):]...)
			i--
		}
	}
	for i:=0;i<len(s1);i++{
		if s1[0]==' '{
			s1=s1[1:]
			i--
		}else{
			break
		}
	}
	for s1[len(s1)-1]==' '{
		s1=s1[:len(s1)-1]
	}
	//fmt.Println(string(s1))
	reverseWords2(s1)
	//fmt.Println(string(s1))
	n:=len(s1)
	for i:=0;i<n;i++{
		if s1[i]==' '{
			continue
		}
		index:=i
		for i<n&&s1[i]!=' '{
			i++
		}
		reverseWords2(s1[index:i])
	}
	return string((s1))
}
func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res:=make([][]int,0)
	var dfs func(sum,index int,sub []int)
	dfs=func(sum,index int,sub []int){
		if candidates[index]+sum>target{
			return
		}else if candidates[index]+sum==target{
			tmp:=make([]int,len(sub)+1)
			copy(tmp,append(sub,candidates[index]))
			res=append(res,tmp)
		}else{
			for i:=index;i<len(candidates);i++{
				dfs(sum+candidates[index],i,append(sub,candidates[index]))
			}
		}
	}
	for i:=0;i<len(candidates);i++{
		dfs(0,i,[]int{})
	}
	return res
}
var res [][]int
func test(sub []int) [][]int{
	res=append(res,append(sub,2))
	test(sub)
	return res
}
func sortedArrayToBST2(nums []int,left,right int) *TreeNode {
	if left>right{
		return nil
	}
	mid:=(left+right)/2
	root:=&TreeNode{Val: nums[mid]}
	root.Left=sortedArrayToBST2(nums,left,mid-1)
	root.Right=sortedArrayToBST2(nums,mid+1,right)
	return root
}
func sortedArrayToBST(nums []int) *TreeNode {
	return sortedArrayToBST2(nums, 0, len(nums) - 1)
}
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	cur := head
	for cur.Next!=nil{
		if cur.Next!=nil{
			cur.Next=cur.Next.Next
		}else {
			cur=cur.Next
		}
	}

	return head
}
func longestValidParentheses(s string) int {
	maxAns := 0
	dp := make([]int, len(s))
	for i:=1;i<len(s);i++{
		if s[i]==')' {
			if s[i-1] == '('{
				if i >= 2 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			}else if i-1-dp[i-1] >= 0 && s[i-1-dp[i-1]] == '(' {
				if i-2-dp[i-1] >= 0 {
					dp[i] = dp[i-1] + 2 + dp[i-2-dp[i-1]]
				} else {
					dp[i] = dp[i-1] + 2
				}
			}
		}
		maxAns=int(math.Max(float64(maxAns),float64(dp[i])))
	}

	return maxAns
}
func canCompleteCircuit(gas []int, cost []int) int {
	n:=len(gas)
	for i:=0;i<n;{
		sumgas,sumcos,cnt:=0,0,0
		for cnt<n{
			sumcos+=cost[(cnt+i)%n]
			sumgas+=gas[(cnt+i)%n]
			if sumcos>sumgas{
				break
			}
			cnt++
		}
		if cnt==n{
			return i
		}else{
			i+=cnt+1
		}

	}
	return -1
}
func minPathSum(grid [][]int) int {
	m:=len(grid)
	if m<1{
		return 0
	}
	n:=len(grid[0])
	dp:=make([][]int,m)
	dp[0]=make([]int,n)
	for i:=1;i<n;i++{
		dp[0][i]=dp[0][i-1]+grid[0][i-1]
	}
	for i:=1;i<m;i++{
		dp[i]=make([]int,n)
		dp[i][0]=dp[i-1][0]+grid[i-1][0]
		for j:=1;j<n;j++{
			dp[i][j]=int(math.Min(float64(dp[i-1][j]+grid[i-1][j]),float64(dp[i][j-1]+grid[i][j-1])))
		}
	}
	fmt.Println(dp)
	return dp[m-1][n-1]+grid[m-1][n-1]
}
func climbStairs(n int) int {
	dp:=make([]int,n+1)
	dp[1]=1
	dp[0]=1
	for i:=2;i<=n;i++{
		dp[i]=dp[i-1]+dp[i-2]
	}
	return dp[n]
}
func findOrder1(numCourses int, prerequisites [][]int) []int {
	var edges=make([][]int,numCourses)
	for _,v:=range prerequisites{
		edges[v[1]]=append(edges[v[1]],v[0])
	}
	visited:=make([]int,numCourses)
	res:=make([]int,0)
	valid:=true
	var dfs func(u int)
	dfs= func(u int) {
		visited[u]=1
		for _,v :=range edges[u]{
			if visited[v]==0{
				dfs(v)
				if !valid{
					return
				}
			}else if visited[v]==1{
				valid=false
				return
			}

		}
		visited[u]=2
		res=append(res,u)
	}
	for i:=0;i<numCourses&&valid;i++{
		if visited[i]==0{
			dfs(i)
		}
	}
	if !valid{
		return []int{}
	}
	for i:=0;i<len(res)/2;i++{
		res[i],res[numCourses-i-1]=res[numCourses-i-1],res[i]
	}
	return res
}

func findOrder2(numCourses int, prerequisites [][]int) []int {
	var edges=make([][]int,numCourses)
	inedge:=make([]int,numCourses)
	for _,v:=range prerequisites{
		edges[v[1]]=append(edges[v[1]],v[0])
		inedge[v[0]]++
	}
	res:=make([]int,0)
	q:=make([]int,0)
	for i:=0;i<numCourses;i++{
		if inedge[i]==0{
			q=append(q,i)
		}
	}
	for len(q)>0{
		u:=q[0]
		q=q[1:]
		res=append(res,u)
		for _,v:=range edges[u]{
			inedge[v]--
			if inedge[v]==0{
				q=append(q,v)
			}
		}
	}

	if len(res)!=numCourses{
		return []int{}
	}
	return res
}
func getKthFromEnd(head *ListNode, k int) *ListNode {
	cnt:=0
	cur:=head
	for cur!=nil{
		cur=cur.Next
		cnt++
		if cnt==k{
			break
		}
	}
	if cnt!=k{
		return nil
	}
	slow,fast:=head,cur
	for fast!=nil{
		slow=slow.Next
		fast=fast.Next
	}
	return slow
}
func longestCommonSubsequence(text1 string, text2 string) int {
	m:=len(text1)
	n:=len(text2)
	dp:=make([][]int,m+1)
	for i:=0;i<m+1;i++{
		dp[i]=make([]int,n+1)
	}
	for i:=1;i<m+1;i++{
		for j:=1;j<n+1;j++{
			if text1[i-1]==text2[j-1]{
				dp[i][j]=dp[i-1][j-1]+1
			}else{
				dp[i][j]=int(math.Max(float64(dp[i-1][j]),float64(dp[i][j-1])))
			}
		}
	}
	return dp[m][n]
}
type Node struct {
	Val int
	Next *Node
	Random *Node
}
func copyRandomList1(head *Node,m map[*Node]*Node) *Node {
	if head==nil{
		return nil
	}
	if _,ok:=m[head];ok{
		return m[head]
	}
	node:=&Node{head.Val,nil,nil}
	m[head]=node
	node.Next=copyRandomList1(head.Next,m)
	node.Random=copyRandomList1(head.Random,m)
	return node
}
func copyRandomList(head *Node) *Node {
	m:=make(map[*Node]*Node)
	if head==nil{
		return nil
	}
	node:=copyRandomList1(head,m)
	return node
}
func maximalSquare(matrix [][]byte) int {

	dp := make([][]int, len(matrix))
	maxSide := 0
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[i]))
		for j := 0; j < len(matrix[i]); j++ {
			dp[i][j] = int(matrix[i][j] - '0')
			if dp[i][j] == 1 {
				maxSide = 1
			}
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if dp[i][j]==1{
				dp[i][j]=int(math.Min(float64(dp[i-1][j]),float64(dp[i][j-1])))
				dp[i][j]=int(math.Min(float64(dp[i][j]),float64(dp[i-1][j-1])))+1
				if dp[i][j]>maxSide{
					maxSide=dp[i][j]
				}
			}
		}
	}
	return int(math.Pow(float64(maxSide),float64(2)))
}
func sumNumbers2(root *TreeNode,s int) int {
	if root==nil {
		return 0
	}
	sum:=10*s+root.Val
	if root.Right==nil&&root.Left==nil{
		return sum
	}
	return sumNumbers2(root.Left,sum)+sumNumbers2(root.Right,sum)
}
func sumNumbers(root *TreeNode) int {
	if root==nil{
		return 0
	}
	return sumNumbers2(root,0)
}
func search2(nums []int, target int) int {
	left,right:=0,len(nums)-1
	for left<=right{
		mid:=(left+right)/2
		if nums[mid]==target{
			return mid
		}else if nums[mid]<target{
			left=mid+1
		}else {
			right=mid-1
		}

	}
	return -1
}
func invertTree(root *TreeNode) *TreeNode {
	if root==nil{
		return nil
	}
	if root.Left!=nil{
		invertTree(root.Left)
	}
	if root.Right!=nil{
		invertTree(root.Right)
	}
	tmp:=root.Left
	root.Left=root.Right
	root.Right=tmp
	return root
}
func maxDepth(root *TreeNode) int {
	if root==nil{
		return 0
	}
	return int(math.Max(float64(maxDepth(root.Left)),float64(maxDepth(root.Right))))+1
}
func inorderTraversal(root *TreeNode) []int {
	res:=make([]int,0)
	var inorderTraversal2 func(root  *TreeNode)
	inorderTraversal2=func(root  *TreeNode){
		if root==nil{
			return
		}
		inorderTraversal2(root.Left)
		res=append(res,root.Val)
		inorderTraversal2(root.Right)
	}
	inorderTraversal2(root)
	return res
}
//func isMatch(s string, p string) bool {
//	m:=len(s)
//	n:=len(p)
//	dp:=make([][]bool,m+1)
//	for i:=0;i<m+1;i++{
//		dp[i]=make([]bool,n+1)
//	}
//	match:=func(i,j int) bool{
//		if i==0{
//			return false
//		}
//		if p[j-1]=='.'{
//			return true
//		}
//		return s[i-1]==p[j-1]
//	}
//	dp[0][0]=true
//	for i:=0;i<m+1;i++{
//		for j:=1;j<n+1;j++{
//			if p[j-1]=='*'{
//				dp[i][j]=dp[i][j]||dp[i][j-2]
//				if match(i,j-1){
//					dp[i][j]=dp[i][j]||dp[i-1][j]
//				}
//			}else if match(i,j){
//				dp[i][j]=dp[i][j]||dp[i-1][j-1]
//			}
//		}
//	}
//	return dp[m][n]
//}
type Solution struct {
	data []int
}


func Constructor(nums []int) Solution {
	data := nums[0:]
	s := &Solution{
		data,
	}
	return *s
}


func (this *Solution) Pick(target int) int {
	count:=0
	index:=-1
	n:=len(this.data)
	rand.Seed(time.Now().UnixNano())
	for i:=0;i<n;i++{
		if this.data[i]==target{
			count++
			if rand.Intn(count)%(count)==0{
				index=i
			}
		}
	}
	return index
}
func myAtoi(s string) int {
	for len(s)>0&&s[0]==' '{
		s=s[1:]
	}
	if len(s)<1{
		return 0
	}
	subtraction:=false
	if s[0]=='-'{
		subtraction=true
		s=s[1:]
	}else if s[0]=='+'{
		s=s[1:]
	}
	sum:=0
	for i:=0;i<len(s);i++{

		if s[i]>'9'||s[i]<'0'{
			break
		}
		var c  = s[i]-'0'
		sum*=10

		fmt.Println(int(c))
		sum+=int(c)
		//sum=int(math.Abs(float64(sum*10+int(c))))
		if subtraction==false&&sum>math.MaxInt32{
			return math.MaxInt32
		}else if subtraction==true&&-sum < math.MinInt32{
			return math.MinInt32
		}
	}
	if subtraction{
		sum=-sum
	}
	return sum
}
func isBipartite(graph [][]int) bool {
	var (
		UNCOLORED, RED, GREEN = 0, 1, 2
		color []int
		valid bool
	)
	n := len(graph)
	valid = true
	color = make([]int, n)
	var dfs func(index,i int)
	dfs=func(index,c int){
		color[index]=c
		cNei:=RED
		if c==RED{
			cNei=GREEN
		}
		for _,neighbor:=range graph[index]{
			if color[neighbor]==UNCOLORED{
				dfs(neighbor,cNei)
				//color[neighbor]=cNei
				if !valid{
					return
				}
			}else if color[neighbor]!=cNei{
				valid=false
				return
			}
		}
	}
	for i := 0; i < n && valid; i++ {
		if color[i] == UNCOLORED {
			dfs(i, RED)
		}
	}
	return valid
}
func removeKdigits(num string, k int) string {
	//s:=make([]byte,0)
	//n:=len(num)
	//s=append(s,'0')
	//for i:=0;i<n;i++{
	//	for k>0&&len(s)>0&&s[len(s)-1]>num[i]{
	//		s=s[:len(s)-1]
	//		k--
	//	}
	//	s=append(s,num[i])
	//}
	//res:=strings.TrimLeft(string(s),"0")
	//if res==""{
	//	res="0"
	//}
	//return res
	stack := []byte{}
	for i := range num {
		digit := num[i]
		for k > 0 && len(stack) > 0 && digit < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, digit)
	}
	stack = stack[:len(stack)-k]
	ans := strings.TrimLeft(string(stack), "0")
	if ans == "" {
		ans = "0"
	}
	return ans
}

func singleNumber(nums []int) int {
	result := 0
	for _,x:=range nums{
		result=result^x
	}
	return result
}
type MySlice []int
func (m MySlice) Less(i,j int) bool{
	a:=strconv.Itoa(m[i])+strconv.Itoa(m[j])
	b:=strconv.Itoa(m[j])+strconv.Itoa(m[i])
	if a<=b{
		return true
	}else{
		return false
	}
}
func (m MySlice) Swap(i,j int){
	m[i],m[j]=m[j],m[i]
}
func (m MySlice) Len(i,j int) int{
	return len(m)
}
type mm []int
func (m mm)MyLess(i,j int) bool{
	a:=strconv.Itoa(m[i])+strconv.Itoa(m[j])
	b:=strconv.Itoa(m[j])+strconv.Itoa(m[i])
	if a<=b{
		return true
	}else{
		return false
	}
}
func minNumber(nums []int) string {
	sort.SliceStable(nums,func (i,j int) bool{
		a:=strconv.Itoa(nums[i])+strconv.Itoa(nums[j])
		b:=strconv.Itoa(nums[j])+strconv.Itoa(nums[i])
		if a<=b{
			return true
		}else{
			return false
		}
	})
	str:=""
	for _,v:=range nums{
		str+=strconv.Itoa(v)
	}
	return str
}
func compare(a, b int) bool {
	//int to string
	m := strconv.Itoa(a) + strconv.Itoa(b)
	n := strconv.Itoa(b) + strconv.Itoa(a)
	//string to int
	mm, _ := strconv.Atoi(m)
	nn, _ := strconv.Atoi(n)
	if mm > nn {
		return true
	}
	return false
}
func maxSlidingWindow(nums []int, k int) []int {
	n:=len(nums)
	q:=make([]int,0)
	res:=make([]int,0)
	for i:=0;i<n;i++{
		if len(q)!=0&& i>=q[0]+k{
			q=q[1:]
		}
		for len(q)!=0&&nums[q[len(q)-1]]<=nums[i]{
			q=q[:len(q)-1]
		}
		q=append(q,i)
		if i>=k-1{
			res=append(res,nums[q[0]])
		}
	}
	return res
}
func isNumber(s string) bool {
	var (
		blank  = 0 // 空格
		sign1  = 1 // +/- 无e前缀
		digit1 = 2 // 数字(0-9) 无前缀
		point  = 4 // '.'
		digit2 = 5 // 数字(0-9) 有符号前缀
		e      = 6 // 'e'
		sign2  = 7 // +/- 有e前缀
		digit3 = 8 // 数字(0-9) 有e前缀
	)


	s = strings.TrimRight(s, " ")
	dfa := [][]int{
		[]int{blank, sign1, digit1, point, -1},
		[]int{-1, -1,digit1,  point, -1},
		[]int{-1, -1, digit1, digit2, e},
		[]int{-1, digit2, -1, -1, e},
		[]int{-1, digit2, -1, -1, -1},
		[]int{-1, digit2, -1, -1, e},
		[]int{-1, digit3, sign2, -1, -1},
		[]int{-1, digit3, -1, -1, -1},
		[]int{-1, digit3, -1, -1, -1},
	}

	state := 0 // blank start
	for i := 0; i < len(s); i++ {
		var newState int
		switch s[i] {
		case ' ':
			newState = 0
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			newState = 2
		case '+', '-':
			newState = 1
		case '.':
			newState = 3
		case 'e':
			newState = 4
		case 'E':
			newState = 4
		default:
			return false
		}
		state = dfa[state][newState]
		if state == -1 {
			return false
		}
	}
	return state == digit1 || state == digit2 || state == digit3
}
func Permutation( str string ) []string {
	res:=make([]string,0)
	m:=make(map[string]int)
	var dfs func(substr string, d int,v []int)
	dfs=func(substr string,d int,v []int){
		if d>=len(str){
			res=append(res,substr)
			return
		}
		for i:=0;i<len(str);i++{
			if v[i]==1{
				continue
			}
			tmp:=substr+string(str[i])
			_,ok:=m[tmp]
			if ok{
				continue
			}
			v[i]=1
			m[tmp]=1
			dfs(tmp,d+1,v)
			v[i]=0

		}
	}
	for i:=0;i<len(str);i++{
		_,ok:=m[string(str[i])]
		if ok{
			continue
		}
		m[string(str[i])]=1
		v:=make([]int,len(str))
		v[i]=1
		dfs(string(str[i]),1, v)
		v[i]=0
	}
	return res
}
func lengthOfLongestSubstring(s string) int {
	m:=make(map[byte]int)
	res:=0
	index:=-1
	for i:=0;i<len(s);i++{
		_,ok:= m[s[i]]
		if ok && m[s[i]]>index{
			index=m[s[i]]
		}
		m[s[i]]=i
		if res<i-index{
			res=i-index
		}
	}
	return res
}
func isMatch(s string, p string) bool {
	m:=len(s)
	n:=len(p)
	dp:=make([][]bool,m+1)
	for i:=0;i<m+1;i++{
		dp[i]=make([]bool,n+1)
	}
	match:=func(i,j int) bool{
		if i==0{
			return false
		}
		if p[j-1]=='.'{
			return true
		}
		return s[i-1]==p[j-1]
	}
	dp[0][0]=true
	for i:=0;i<m+1;i++{
		for j:=1;j<n+1;j++{
			if p[j-1]=='*'{
				dp[i][j]=dp[i][j]||dp[i][j-2]
				if match(i,j-1){
					dp[i][j]=dp[i][j]||dp[i-1][j]
				}
			}else if match(i,j){
				dp[i][j]=dp[i][j]||dp[i-1][j-1]
			}
		}
	}
	return dp[m][n]
}
func strStr(s string,p string) int{
	n:=len(s)
	m:=len(p)
	if m==0{
		return  0
	}
	s=string(append([]byte{' '},s...))
	p=string(append([]byte{' '},p...))
	next:=make([]int,m+1)
	for i,j:=2,0;i<=m;i++{
		for j!=0 && p[i]!=p[j+1]{
			j=next[j]
		}
		if p[i]==p[j+1]{
			fmt.Println(string(p[i]))
			j++
		}
		next[i]=j
	}
	fmt.Println(next)
	for i,j:=1,0;i<=n;i++{
		for j!=0 && s[i]!=p[j+1]{
			j=next[j]
		}
		if s[i]==p[j+1]{
			j++
		}
		if j==m {
			return i-m
		}
	}
	return -1
}
func file_read(){
	type mystruct struct{
		BpId int `json:"BpId"`
		CloudId string `json:"CloudId"`
		CosRegions string `json:"CosRegions"`
	}
	f1,err:=os.Open("/Users/timellchen/Desktop/base.json")
	defer func(){
		fmt.Println("close file")
		f1.Close()
	}()
	if err!=nil{
		fmt.Println("读取文件失败")
		fmt.Println(err.Error())
	}
	//defer fmt.Println("11111")
	//defer fmt.Println("22222")
	s:=make([]byte,0)
	for{
		b:=make([]byte,1024)
		n,err2:=f1.Read(b)
		if err2==nil{
			fmt.Println(n)
			s=append(s,b[:n]...)
			//fmt.Println(string(b))
		}else if err2==io.EOF{
			fmt.Println("EOF")

			break
		}else{
			fmt.Println("error")
			fmt.Println(err2.Error())
			return
		}
	}
	fmt.Println("start decode")
	var tmp mystruct
	err3:=json.Unmarshal(s,&tmp)
	if err3!=nil{
		fmt.Println(err3.Error())
	}
	fmt.Println(tmp)
	type mystruct2 struct{
		BpId string `json:"region_name"`
		CloudId string `json:"region_name_zh"`
	}
	var tmp2 []mystruct2
	err4:=json.Unmarshal([]byte(tmp.CosRegions),&tmp2)
	if err4!=nil{
		fmt.Println(err4.Error())
	}
	fmt.Println(tmp2)
}
func test_interface(){
	array:=make([]interface{},3)
	array[0]=1
	array[1]="string"
	array[2]=true
	for _,v :=range array{
		switch v.(type) {
		case int:
			fmt.Println("int")
		case string:
			fmt.Println("sting")
		case bool:
			fmt.Println("bool")
		default:
			fmt.Println("none")
		}
	}
	x:=1
	fmt.Println(reflect.TypeOf(x).Name())
}
func CyclePrint(){
	for i:=0;i<5;i++{
		fmt.Println("i:",i)
		var buf [64]byte
		n:=runtime.Stack(buf[:],false)
		fmt.Println("i",i,"goid",(strings.Fields(string(buf[:n])))[1])
		runtime.Goexit()
	}

}

//8 6
//3 2 1 1 2 3
//3 3 2 1 2 3
//2 3 1 1 1 1
//3 1 2 2 1 3
//3 3 3 1 3 1
//1 2 3 2 2 1
//3 1 3 1 1 3
//1 2 3 2 3 3

