# Borderlands 4 Serials

A Go library for encoding and decoding Borderlands 4 item serial codes.

## Overview

This library provides functionality to parse, manipulate, and serialize Borderlands 4 weapon/item codes.  
It handles the bitstream protocol used by the game to represent items and their parts.

## Features

- **Deserialize** binary item codes into a structured format
- **Serialize** item data back into binary format
- **String representation** for human-readable item codes
- **Parse from string** to reconstruct items from text format

## Supported Features

- Variable-length integer encoding (varint)
- Variable-length bit encoding (varbit)
- Part blocks with three subtypes:
    - Simple parts (index only)
    - Integer subtype (index + single value)
    - List subtype (index + array of values)

## Limitations

- Major type "111" is not supported:
  - It prevents the skin from being handled, and gets stripped of the item as a result. (Example: phosphene skin.)
  - It prevents DLC items from being decoded. Because DLC items are paid, support will not be added.

That's the only known limitation.

## Base85 to bitstream
Before doing anything, the serials must be turned into an usable bitstream.

- **Step 1**: strip the leading `@U` from the serial, as it is not part of the base85 string.
- **Step 2**: convert from base85 to bytes (big endian) using this custom alphabet:  
  ```
  0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!#$%&()*+-;<=>?@^_`{/}~
  ```
- **Step 3**: mirror each byte. (e.g. `0b10000111` becomes `0b11100001`).
- **Step 4**: decode the bitstream, see below.   

**Important:** this data is a single contiguous bitstream.  
This bitstream uses bitpacked datatypes of variable bit length, **with no byte boundaries**.  
Everything has to be read bit by bit **with no regards to byte alignment**.   

Practical example:

```
1. Serial:
  @UgydOV%h><

2. Strip leading @U:
  gydOV%h><

3. Base85 to bytes (big endian):
  84e4595ccbd900

4. Mirror each byte:
  21279a3ad39b00

5. Bitstream:
  00100001001001111001101000111010110100111001101100000000
```

## Decoding the bitstream

### Token-based stream
First, the bitstream always starts with this magic 7-bit header `0010000`.  
Then, it consists of tokens that dictates how to read subsequent data.  
Finally, the bitstream ends with a hard terminator token (`00`).

**Note:** zero padding is used at the end to align to byte boundary, but is not part of the data.

### Token Types (variable length)
The tokens are variable-length, either 2 or 3 bits.   
This is possible because the binary prefixes are unique.   
(See Wikipedia: [Prefix code](https://en.wikipedia.org/wiki/Prefix_code))

- **`00`** (2 bits) - `TOK_SEP1` - Hard separator/terminator, string representation `|`
- **`01`** (2 bits) - `TOK_SEP2` - Soft separator, string representation `,`
- **`100`** (3 bits) - `TOK_VARINT` - Followed by a nibble-based variable integer, string representation `123`
- **`110`** (3 bits) - `TOK_VARBIT` - Followed by a length-prefixed variable integer, string representation `123`
- **`101`** (3 bits) - `TOK_PART` - Followed by a complex part structure, string representation `{54}` or `{54:3}` or `{54:[3 2 1]}`
- **`111`** (3 bits) - `TOK_UNSUPPORTED_111` - Unknown data block, found on specific items. (DLC items + phosphene skin)  
  **Support will not be added for type 111.**

### Data Encoding

#### VARINT (after `100` token)
This is a nibble-based variable integer encoding:

- Reads 4-bit nibbles (LSB-first)
- Each nibble followed by continuation bit:
    - `1` = read next nibble
    - `0` = stop reading
- Max 4 blocks (16 bits total)

Read more about VARINTs [here](_docs/VARINT16.md).

#### VARBIT (after `110` token)
This is a bit-length-prefixed variable integer encoding:

- 5-bit length prefix (LSB-first)
- Followed by N bits of actual value (LSB first)
- Length of 0 either means 0 value or possibly 32 bits, unverified.

Read more about VARBITs [here](_docs/VARBIT32.md).

**Notes about VARINT and VARBIT:**   
This was discovered through testing.   
VARINTs and VARBITs can be used interchangeably, notably in the manufacturer index.   

#### PART (after `101` token)
This is a complex structure representing an item part.   
It links an index to either no value, a single integer value, or a list of integer values.

- **Index**: VARINT (**without** its `100` token)
- **Type flag** (1 bit):
  - `1` → SUBTYPE_INT: VARINT value + `000` terminator
  - `0` → Read 2 more bits:
    - `10` → SUBTYPE_NONE (no data)
    - `01` → SUBTYPE_LIST:
      - `01` token starts list
      - Sequence of `100`/`110` tokens with its VARINT/VARBIT values
      - `00` token ends list

Read more about PARTs [here](_docs/PART.md).

## How did we discover all this?

By staring at bitstreams for 10 days straight, endless tests and sleepless nights.  
Establishing theories, comparing known values.   
Injecting zero-value continuation block into VARINTs.   
Swapping VARINTs with their VARBIT counterparts, etc.

**You have no idea how much trial and error this took.**   

A very small subset of our research notes are available in the `_notes` folder.   
You can read our journal [here](_docs/JOURNAL.md).   
(Thanks to @InflamedSebi for writing it down.)

## Credits

- @Sparkie for his initial base85 alphabet discovery.  
- @Nicnl for his "buyback phenomenon" discovery. (See journal.)
- @InflamedSebi for his byte mirroring discovery.
- @Nicnl for his VARINT discovery.   
- @InflamedSebi for his token prefix discovery.
- @InflamedSebi for his VARBIT discovery.
- @Nicnl for his variable-length token discovery.
- @Nicnl for his 101 PART structure discovery.
- @InflamedSebi for writing down the journal.
