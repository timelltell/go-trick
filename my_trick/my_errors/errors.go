package my_errors

import "errors"

type MyError struct {
	Code int32
	Message string
	//Details []*any.Any

}

type ZooTour2 interface {
	Enter()
	VisitPanda(panda *string)
	VisitTiger(tiger *string)
	Leave()

	Err() error
}
func Tour2(t ZooTour2, panda *string,tiger *string) error{
	t.Enter()
	t.VisitPanda(panda)
	t.VisitTiger(tiger)
	t.Leave()

	if err:=t.Err(); err != nil {
		return errors.New("err")
	}
	return nil
}