package model

type student struct {
	Name  string
	score float64
}

// 一个工厂模式的函数，相当于构造函数
func NewStudent(name string) *student {
	return &student{
		Name: name,
	}
}

func (stu *student) GetScore() float64 {
	return stu.score
}

func (stu *student) SetScore(score float64) {
	stu.score = score
}
