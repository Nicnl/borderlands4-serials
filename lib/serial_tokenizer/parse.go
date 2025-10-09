package serial_tokenizer

import (
	"borderlands_4_serials/lib/bit_reader"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// -----------------------------
// Public AST node types
// -----------------------------

// VarBit preserves exact bit length for faithful reserialization.
type VarBit struct {
	Len   uint32 `json:"len"`
	Value int64  `json:"value"`
}

// KVEntry is one "111" entry: Key -> Values...
type KVEntry struct {
	Key    any   `json:"key"`    // usually an int (varint), but allow VarBit for future-proofing
	Values []any `json:"values"` // sequence of varints/varbits
}

// KV represents a whole "101 ... 00 00" section.
type KV struct {
	Entries []KVEntry `json:"entries"`
}

// -----------------------------
// Tokenizer wrapper
// -----------------------------

type Tokenizer struct {
	bs *bit_reader.BitReader

	// 1-token pushback buffer to let us "peek" by reading & pushing back
	hasPushback bool
	pushTok     Token
}

func NewTokenizer(data []byte) *Tokenizer {
	return &Tokenizer{bs: bit_reader.NewBitReader(data)}
}

// -----------------------------
// Tokens (must match your existing constants)
// -----------------------------

type Token byte

const (
	TOK_SEP1    Token = iota // "01" soft separator
	TOK_SEP2                 // "00" hard separator
	TOK_VARINT               // "100" ... nibble varint
	TOK_VARBIT               // "110" ... varbit
	TOK_KV_MODE              // "101" ... enter KV section
	TOK_KV_ADD               // "111" ... new KV entry inside section
)

// -----------------------------
// Helpers
// -----------------------------

// expect reads exact bit pattern (debug-only)
func (t *Tokenizer) expect(msg string, bits ...byte) error {
	for _, bit := range bits {
		b, ok := t.bs.Read()
		if !ok {
			return io.EOF
		}
		if b != bit {
			return fmt.Errorf("%s => expected bit %d, got %d", msg, bit, b)
		}
	}
	return nil
}

// pushback & read with pushback

func (t *Tokenizer) pushBack(tok Token) {
	t.hasPushback = true
	t.pushTok = tok
}

func (t *Tokenizer) readTok() (Token, error) {
	if t.hasPushback {
		t.hasPushback = false
		return t.pushTok, nil
	}
	return t.nextToken()
}

// readSkippingSoftSeps consumes any number of TOK_SEP1; returns first non-SEP1 (or eof).
func (t *Tokenizer) readSkippingSoftSeps() (Token, error) {
	for {
		tok, err := t.readTok()
		if err != nil {
			return 0, err
		}
		if tok == TOK_SEP1 {
			continue
		}
		return tok, nil
	}
}

// -----------------------------
// Top-level Parse
// -----------------------------

// Parse walks the whole stream and returns a []any AST.
// Nodes may be: int, VarBit, KV (with possibly empty Entries).
func (t *Tokenizer) Parse() ([]any, error) {
	// Magic "001"
	if err := t.expect("magic", 0, 0, 1); err != nil {
		return nil, fmt.Errorf("magic header: %w", err)
	}

	var out []any

	// Main loop
	for {
		tok, err := t.readSkippingSoftSeps()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			// Some implementations return generic error at EOF:
			// if your nextToken() returns "unexpected end of data", treat as EOF:
			if err.Error() == "unexpected end of data" {
				break
			}
			return nil, err
		}

		switch tok {
		case TOK_SEP2:
			// Hard separator at top level — ignore as boundary.
			// (We don’t need to store separator tokens to reserialize canonically.)
			continue

		case TOK_VARINT:
			v, err := t.readVarInt()
			if err != nil {
				return nil, err
			}
			out = append(out, int(v))

		case TOK_VARBIT:
			val, ln, err := t.readVarBit()
			if err != nil {
				return nil, err
			}
			out = append(out, VarBit{Len: ln, Value: int64(val)})

		case TOK_KV_MODE:
			kv, err := t.parseKVSection()
			if err != nil {
				return nil, err
			}
			out = append(out, kv)

		case TOK_KV_ADD:
			// A bare TOK_KV_ADD at top-level is unexpected: treat as error,
			// but if you want to be permissive, you could push it back and try parseKVSection().
			return nil, fmt.Errorf("unexpected TOK_KV_ADD at top-level")

		default:
			return nil, fmt.Errorf("unknown token %d", tok)
		}
	}

	// Debug (optional)
	if dbg, err := json.MarshalIndent(out, "", "  "); err == nil {
		fmt.Println(string(dbg))
	}

	return out, nil
}

// -----------------------------
// KV Section parser ("101 ... 00 00")
// -----------------------------

// parseKVSection assumes the leading TOK_KV_MODE ("101") has JUST been consumed.
func (t *Tokenizer) parseKVSection() (KV, error) {
	var kv KV
	// We stay inside until we encounter "00 00" (two hard seps with no token between).
	// Soft seps (01) can appear anywhere and are ignored.

	seenFirstHard := false

	for {
		tok, err := t.readSkippingSoftSeps()
		if err != nil {
			if errors.Is(err, io.EOF) || err.Error() == "unexpected end of data" {
				// Unterminated section, but keep what we have (so you can reserialize/repair).
				return kv, nil
			}
			return kv, err
		}

		if tok == TOK_SEP2 {
			// Could be first of the terminating pair
			if seenFirstHard {
				// TOK_SEP2 TOK_SEP2 => end of this section
				return kv, nil
			}
			seenFirstHard = true
			// absorb any number of soft seps between the two hard seps
			// but if ANY non-separator token appears, it's not a terminator pair.
			for {
				tok2, err2 := t.readTok()
				if err2 != nil {
					if errors.Is(err2, io.EOF) || err2.Error() == "unexpected end of data" {
						// treat single trailing 00 as just a boundary and exit section anyway
						return kv, nil
					}
					return kv, err2
				}
				if tok2 == TOK_SEP1 {
					// keep looping; still between potential pair
					continue
				}
				if tok2 == TOK_SEP2 {
					// Confirmed pair -> end section
					return kv, nil
				}
				// Not a terminator pair; push back what we saw and continue parsing normally
				t.pushBack(tok2)
				seenFirstHard = false
				break
			}
			// If we didn’t return above, we fell back to normal flow after pushback.
			continue
		}

		// Any non-hard-sep token resets the pending-pair state
		seenFirstHard = false

		if tok == TOK_KV_ADD {
			entry, err := t.parseKVEntry()
			if err != nil {
				return kv, err
			}
			kv.Entries = append(kv.Entries, entry)
			continue
		}

		// It’s valid for a section to be empty (no entries) and then immediately closed.
		// Any other token type directly inside 101 is unexpected but we’ll try to be permissive:
		switch tok {
		case TOK_VARINT:
			// If the encoder ever allows section metadata: slurp and drop into a dummy empty entry.
			v, err := t.readVarInt()
			if err != nil {
				return kv, err
			}
			kv.Entries = append(kv.Entries, KVEntry{
				Key:    int(v),
				Values: nil,
			})
		case TOK_VARBIT:
			val, ln, err := t.readVarBit()
			if err != nil {
				return kv, err
			}
			kv.Entries = append(kv.Entries, KVEntry{
				Key:    VarBit{Len: ln, Value: int64(val)},
				Values: nil,
			})
		case TOK_KV_MODE:
			// Nested 101? Support recursion:
			child, err := t.parseKVSection()
			if err != nil {
				return kv, err
			}
			kv.Entries = append(kv.Entries, KVEntry{
				Key:    "section",    // sentinel; you can choose something else
				Values: []any{child}, // keep nested section as a value
			})
		case TOK_KV_ADD:
			// handled above (unreachable here)
		default:
			// ignore or error; choose permissive:
			return kv, fmt.Errorf("unexpected token %d inside KV section", tok)
		}
	}
}

// parseKVEntry assumes the leading TOK_KV_ADD ("111") has just been consumed.
// Grammar (per your latest spec):
//
//	Entry := "111" [SEP1]* Key ( [SEP1]* Value )*
//	Key   := VARINT (usually) | VARBIT (be liberal)
//	Value := VARINT | VARBIT
//
// Values stop when we see either another "111" (new entry) or the section terminator "00 00".
func (t *Tokenizer) parseKVEntry() (KVEntry, error) {
	var e KVEntry

	// After 111, skip any soft separators
	tok, err := t.readSkippingSoftSeps()
	if err != nil {
		if errors.Is(err, io.EOF) || err.Error() == "unexpected end of data" {
			// Empty entry with no key; keep as-is
			return e, nil
		}
		return e, err
	}

	// Key
	switch tok {
	case TOK_VARINT:
		v, err := t.readVarInt()
		if err != nil {
			return e, err
		}
		e.Key = int(v)
	case TOK_VARBIT:
		val, ln, err := t.readVarBit()
		if err != nil {
			return e, err
		}
		e.Key = VarBit{Len: ln, Value: int64(val)}
	default:
		// Be tolerant: push back and leave Key nil if encoder does strange stuff
		t.pushBack(tok)
		e.Key = nil
	}

	// Values: zero or more until next 111 or "00 00"
	for {
		// Absorb soft seps
		tok, err := t.readSkippingSoftSeps()
		if err != nil {
			if errors.Is(err, io.EOF) || err.Error() == "unexpected end of data" {
				return e, nil
			}
			return e, err
		}

		// Handle “00 00” → end of section; we need to *return the entry*
		if tok == TOK_SEP2 {
			// Could be first of the pair; check the next tokens
			// (We must leave the second 00 on the input for the section parser to catch.)
			// Strategy: if next non-SEP1 is SEP2, push it back and return the entry.
			for {
				tok2, err2 := t.readTok()
				if err2 != nil {
					if errors.Is(err2, io.EOF) || err2.Error() == "unexpected end of data" {
						return e, nil
					}
					return e, err2
				}
				if tok2 == TOK_SEP1 {
					continue
				}
				// push back what we saw; the section parser will decide
				t.pushBack(tok2)
				break
			}
			// We encountered a hard separator; treat as end of values for this entry
			return e, nil
		}

		// New entry
		if tok == TOK_KV_ADD {
			// Push it back so the caller can start a new entry
			t.pushBack(tok)
			return e, nil
		}

		switch tok {
		case TOK_VARINT:
			v, err := t.readVarInt()
			if err != nil {
				return e, err
			}
			e.Values = append(e.Values, int(v))
		case TOK_VARBIT:
			val, ln, err := t.readVarBit()
			if err != nil {
				return e, err
			}
			e.Values = append(e.Values, VarBit{Len: ln, Value: int64(val)})
		case TOK_KV_MODE:
			// Nested section as a value—support recursion
			child, err := t.parseKVSection()
			if err != nil {
				return e, err
			}
			e.Values = append(e.Values, child)
		default:
			// Unknown token inside an entry; push back and stop values
			t.pushBack(tok)
			return e, nil
		}
	}
}
