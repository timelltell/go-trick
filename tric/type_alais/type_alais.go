package type_alais

type Order struct {

}
func (this *Order)Setstatus(status int){

}
func (this *Order)Setstatus2(status Mystatus){

}
const (
	Normal = iota + 1
	Deleting
	Deleted
)
type Mystatus int
const (
	Normal1 Mystatus = iota + 1
	Deleting1
	Deleted1
)

func order2(){
	var o = new(Order)
	o.Setstatus(Normal)
	o.Setstatus2(Normal1)
}