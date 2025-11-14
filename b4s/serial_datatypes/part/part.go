package part

import (
	"fmt"
	"strconv"
	"strings"
)

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
		return "{" + strconv.Itoa(int(p.Index)) + "}"
	case SUBTYPE_INT:
		return "{" + strconv.Itoa(int(p.Index)) + ":" + strconv.Itoa(int(p.Value)) + "}"
	case SUBTYPE_LIST:
		var output strings.Builder
		output.WriteString(fmt.Sprintf("{%d:[", p.Index))

		for i, v := range p.Values {
			if i != 0 {
				output.WriteByte(' ')
			}
			output.WriteString(fmt.Sprintf("%d", v))
		}
		output.WriteString("]}")
		return output.String()
	default:
		return fmt.Sprintf("{ERR_UNKNOWN_PART:%d}", p.SubType)
	}
}
