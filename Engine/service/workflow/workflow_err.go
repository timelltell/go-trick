package workflow

type WFErr struct {
	e error
}

func (b WFErr) Error() string {
	return b.e.Error()
}

func NewBizErr(e error) error {
	return WFErr{e: e}
}

type RealErr struct {
	e error
}

func (f RealErr) Error() string {
	return f.e.Error()
}

func NewRealErr(e error) error {
	return RealErr{e: e}
}
