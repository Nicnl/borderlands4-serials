package serial_tokenizer

import "fmt"

func (t *Tokenizer) Expect(msg string, bits ...byte) error {
	for _, bit := range bits {
		b, ok := t.br.Read()
		if !ok {
			return fmt.Errorf("unexpected end of data")
		}
		if b != bit {
			return fmt.Errorf(msg+" => expected bit %d, got %d", bit, b)
		}
	}
	return nil
}
