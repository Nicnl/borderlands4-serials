package part

import "fmt"

type PartSubType uint32

const (
	SUBTYPE_NONE PartSubType = iota
	SUBTYPE_INT
	SUBTYPE_LIST
)

type Part struct {
	Index   uint32
	SubType PartSubType
	Value   uint32
	Values  []uint32
}

func (p *Part) String() string {
	switch p.SubType {
	case SUBTYPE_NONE:
		return fmt.Sprintf(" {%d}", p.Index)
	case SUBTYPE_INT:
		return fmt.Sprintf(" {%d:%d}", p.Index, p.Value)
	case SUBTYPE_LIST:
		output := fmt.Sprintf(" {%d:[", p.Index)
		for i, v := range p.Values {
			if i != 0 {
				output += " "
			}
			output += fmt.Sprintf("%d", v)
		}
		output += "]}"
		return output
	default:
		return fmt.Sprintf("{ERR_UNKNOWN_PART:%d}", p.SubType)
	}
}
