package staticarray

const MAXN = 150001

var (
	tree [MAXN][26]int
	end  [MAXN]int
	pass [MAXN]int
	cnt  int
)

func build() {
	cnt = 1
}

func insert(word string) {
	cur := 1
	pass[cur]++
	for i := 0; i < len(word); i++ {
		path := int(byte(word[i]) - 'a')
		if tree[cur][path] == 0 {
			cnt++
			tree[cur][path] = cnt
		}
		cur = tree[cur][path]
		pass[cur]++
	}
	end[cur]++
}

func search(word string) int {
	cur := 1

	for i := 0; i < len(word); i++ {
		path := int(byte(word[i]) - 'a')
		if tree[cur][path] == 0 {
			return 0
		}
		cur = tree[cur][path]
	}
	return end[cur]
}

func prefixNumber(pre string) int {
	cur := 1

	for i := 0; i < len(pre); i++ {
		path := int(byte(pre[i]) - 'a')
		if tree[cur][path] == 0 {
			return 0
		}
		cur = tree[cur][path]
	}
	return pass[cur]
}

func delete(word string) {
	if search(word) > 0 {
		cur := 1

		for i := 0; i < len(word); i++ {
			path := int(byte(word[i]) - 'a')
			pass[tree[cur][path]]--
			if pass[tree[cur][path]] == 0 {
				tree[cur][path] = 0
				return
			}
			cur = tree[cur][path]
		}
		end[cur]--
	}
}

func clear() {
	for i := 0; i <= cnt; i++ {
		for j := 0; j < 26; j++ {
			tree[i][j] = 0
		}
		end[i] = 0
		pass[i] = 0
	}
}
