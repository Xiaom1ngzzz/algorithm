package leetcode432

import (
	"math"
)

type Bucket struct {
	set  []string
	cnt  int
	last *Bucket
	next *Bucket
}

func NewBucket(s string, c int) *Bucket {
	set := make([]string, 0)
	set = append(set, s)
	return &Bucket{
		set:  set,
		cnt:  c,
		last: nil,
		next: nil,
	}
}

func insert(cur, pos *Bucket) {
	cur.next.last = pos
	pos.next = cur.next
	cur.next = pos
	pos.last = cur
}

func remove(cur *Bucket) {
	cur.next.last = cur.last
	cur.last.next = cur.next
	cur.next = nil
	cur.last = nil
}

type AllOne struct {
	head *Bucket
	tail *Bucket
	m    map[string]*Bucket
}

func Constructor() AllOne {
	all := AllOne{
		head: NewBucket("", 0),
		tail: NewBucket("", math.MaxInt),
		m:    make(map[string]*Bucket),
	}
	all.head.next = all.tail
	all.tail.last = all.head
	return all
}

func (this *AllOne) Inc(key string) {
	bucket, ok := this.m[key]
	if !ok {
		// 说明 key 不存在
		// 判断 1 号桶是否存在
		if this.head.next.cnt == 1 {
			this.m[key] = this.head.next
			this.head.next.set = append(this.head.next.set, key)
		} else {
			newBucket := NewBucket(key, 1)
			this.m[key] = newBucket
			insert(this.head, newBucket)
		}
	} else {
		if bucket.next.cnt == bucket.cnt+1 {
			this.m[key] = bucket.next
			bucket.next.set = append(bucket.next.set, key)
		} else {
			newBucket := NewBucket(key, bucket.cnt+1)
			this.m[key] = newBucket
			insert(bucket, newBucket)
		}
		index := 0
		for idx, val := range bucket.set {
			if val == key {
				index = idx
				break
			}
		}
		bucket.set = append(bucket.set[:index], bucket.set[index+1:]...)
		if len(bucket.set) == 0 {
			remove(bucket)
		}
	}
}

func (this *AllOne) Dec(key string) {
	bucket, _ := this.m[key]
	if bucket.cnt == 1 {
		delete(this.m, key)
	} else {
		if bucket.last.cnt == bucket.cnt-1 {
			this.m[key] = bucket.last
			bucket.last.set = append(bucket.last.set, key)
		} else {
			newBucket := NewBucket(key, bucket.cnt-1)
			this.m[key] = newBucket
			insert(bucket.last, newBucket)
		}
	}

	index := 0
	for idx, val := range bucket.set {
		if val == key {
			index = idx
			break
		}
	}
	bucket.set = append(bucket.set[:index], bucket.set[index+1:]...)
	if len(bucket.set) == 0 {
		remove(bucket)
	}
}

func (this *AllOne) GetMaxKey() string {
	return this.tail.last.set[0]
}

func (this *AllOne) GetMinKey() string {
	return this.head.next.set[0]
}
