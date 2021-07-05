package main

type student struct {
	name  string
	score int
}

type students []student

func (s students) Len() int {
	return len(s)
}

func (s students) Less(i, j int) bool {
	return s[i].score < s[j].score
}

func (s students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s *students) Push(x interface{}) {
	*s = append(*s, x.(student))
}

func (s *students) Pop() interface{} {
	old := *s
	n := len(old)
	x := old[n-1]
	*s = old[0 : n-1]
	return x
}
