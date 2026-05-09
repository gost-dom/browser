package dom

type Range struct {
	startContainer Node
	endContainer   Node
	startPos       int
	endPos         int
}

func (r *Range) SetStart(n Node, offset int) error {
	r.startContainer = n
	r.startPos = offset
	return nil
}

func (r *Range) SetEnd(n Node, offset int) error {
	r.endContainer = n
	r.endPos = offset
	return nil
}

func (r Range) StartContainer() Node { return r.startContainer }
func (r Range) EndContainer() Node   { return r.endContainer }
func (r Range) StartOffset() int     { return r.startPos }
func (r Range) EndOffset() int       { return r.endPos }
func (r Range) Collapsed() bool      { return true }
func (r Range) String() string       { return "" }
func (r Range) Detach()              {}
