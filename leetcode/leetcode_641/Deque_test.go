package leetcode641

import (
	"fmt"
	"testing"
)

func TestDeque(t *testing.T) {
	circularDeque := Constructor(3) // 设置容量大小为3
	fmt.Println(circularDeque.limit == 3)

	ret := circularDeque.InsertLast(1) // 返回 true
	if ret == true {
		fmt.Println("circularDeque.InsertLast(1) -> 执行成功")
	} else {
		fmt.Println("circularDeque.InsertLast(1) -> 执行失败")
	}
	ret = circularDeque.InsertLast(2) // 返回 true
	if ret == true {
		fmt.Println("circularDeque.InsertLast(2) -> 执行成功")
	} else {
		fmt.Println("circularDeque.InsertLast(2) -> 执行失败")
	}
	ret = circularDeque.InsertFront(3) // 返回 true
	if ret == true {
		fmt.Println("circularDeque.InsertFront(3) -> 执行成功")
	} else {
		fmt.Println("circularDeque.InsertFront(3) -> 执行失败")
	}
	ret = circularDeque.InsertFront(4) // 已经满了，返回 false
	if ret == false {
		fmt.Println("circularDeque.InsertFront(4) -> 执行成功")
	} else {
		fmt.Println("circularDeque.InsertFront(4) -> 执行失败")
	}
	ret1 := circularDeque.GetRear() // 返回 2
	if ret1 == 2 {
		fmt.Println("circularDeque.GetRear() -> 执行成功")
	} else {
		fmt.Println("circularDeque.GetRear() -> 执行失败")
	}
	ret = circularDeque.IsFull() // 返回 true
	if ret == true {
		fmt.Println("circularDeque.IsFull() -> 执行成功")
	} else {
		fmt.Println("circularDeque.IsFull() -> 执行失败")
	}
	ret = circularDeque.DeleteLast() // 返回 true
	if ret == true {
		fmt.Println("circularDeque.DeleteLast() -> 执行成功")
	} else {
		fmt.Println("circularDeque.DeleteLast() -> 执行失败")
	}
	ret = circularDeque.InsertFront(4) // 返回 true
	if ret == true {
		fmt.Println("circularDeque.InsertFront(4) -> 执行成功")
	} else {
		fmt.Println("circularDeque.InsertFront(4) -> 执行失败")
	}
	ret1 = circularDeque.GetFront() // 返回 4
	if ret1 == 4 {
		fmt.Println("circularDeque.GetFront() -> 执行成功")
	} else {
		fmt.Println("circularDeque.GetFront() -> 执行失败")
	}
}
