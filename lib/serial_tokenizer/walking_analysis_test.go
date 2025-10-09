package serial_tokenizer

import (
	"borderlands_4_serials/lib/b85"
	"fmt"
	"testing"
)

func TestByWalkingAnalysis(t *testing.T) {
	serials := []string{
		"@Ug!pHG38o5YT`HzQ)h-nP",
		"@Ug!pHG38o5YZ7QZg)h-nP",
		"@Ug!pHG38o5YOe&^9)h-nP",
		"@Ug!pHG38o6@O)92A)h-nP",
		"@Ug!pHG38o5YPb#KC)h-nP",
		"@Ug!pHG38o5YMJlF2)h-nP",
		"@Ug!pHG38o4tO)92A)h-nP",
		"@Ug!pHG38o5Y4JxKV)h-nP",
	}

	serialInts := make(map[string][]int64)

	for _, s := range serials {
		data, _ := b85.Decode(s)
		tok := NewTokenizer(data)

		for i := 0; i < len(data)*8; i++ {
			tok.bs.SetPos(i)

			//token, err := tok.nextToken()
			//if err != nil {
			//	continue
			//}

			{
				val, err := tok.readVarInt()
				if err == nil {
					serialInts[s] = append(serialInts[s], int64(val))
				}
			}
			{
				val, err := tok.readVarBit()
				if err == nil {
					serialInts[s] = append(serialInts[s], int64(val))
				}
			}
		}
		fmt.Println("Serial:", s)
		fmt.Println("Ints:", serialInts[s])
		fmt.Println()
	}

}
