package Compare

import "fmt"

type Country struct {
	Name string
}
type City struct {
	Name string

}
type Stringable interface {
	ToString() string
}
func (c City) ToString() string{
	return c.Name
}
func (c *Country) ToString() string{
	return c.Name
}
func PrintStr(s Stringable){
	fmt.Println(s.ToString())
}
