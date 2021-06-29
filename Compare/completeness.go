package Compare
type Shape interface {
	Sides() int
	Area() int
}

//不完整的实现
type Square struct {
	Len int
}
func (s * Square) Sides() int{
	return s.Len
}

//完整的实现
func (s * Square) Area() int{
	return s.Len
}

//通过这句话判断是否完整
var _ Shape=(*Square)(nil)