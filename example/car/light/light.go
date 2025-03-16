package light

import "fmt"

type LightType int

const (
	LightTypeNone = iota
	LightTypeHead
	LightTypeTail
)

func (l LightType) String() string {
	switch l {
	case LightTypeHead:
		return "Head"
	case LightTypeTail:
		return "Tail"
	default:
		return "Unknown"
	}
}

type Light struct {
	LightType LightType
}

func (l *Light) LightOn() {
	fmt.Printf("%s Light On\n", l.LightType)
}
