# Borderlands 4 Serials

A Go library for encoding and decoding Borderlands 4 item serial codes.  
It also has a web-based API and UI for easy access.   
A live demo is available here: [borderlands4-deserializer.nicnl.com](https://borderlands4-deserializer.nicnl.com/)


## Legal disclaimer

This project and its contributors are not affiliated, associated, authorized, endorsed by,
or in any way officially connected with Gearbox Software, 2K Games, Inc.,
or any of their subsidiaries or affiliates.
The names "Borderlands," "Gearbox," "2K," as well as any related names, marks, emblems,
and images are registered trademarks of their respective owners.


## Overview

This library provides functionality to deserialize, manipulate, and reserialize Borderlands 4 weapon/item codes.  
It handles the bitstream protocol used by the game to represent items and their parts.

## Features

- **Deserialize** binary item codes into a structured format
- **Serialize** item data back into binary format
- **String representation** for human-readable item codes
- **Parse from string** to reconstruct items from text format
- **High performance** deserialization/serialization is under 1ms, around 200μs on average.

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

This is the only known limitation.

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

The bitstream contains no ciphering, and no checksum.   
As long as you know how to read the tokens and their data, you can decode and encode the entire item.     

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

## String representation

The game engine most probably deserialize/reserialize the item data from/to structures in memory.   
However, we don't have access to this structure.   
This is why we came up with a human-readable string representation of the item data.   

The deserializer turns the item serials into a human-readable string format, example:  
```
Serial:
  @Ugr$WBm/$!m!X=5&qXxA;nj3OOD#<4R

Result:
  267, 0, 1, 22| 2, 274|| {7} {1} {245:[23 39 69 79]}|
```

This string representation is a 1-to-1 mapping of the bitstream structure, using the following rules:
- Tokens are separated by spaces.
- `TOK_SEP1` (`00`) is represented as `|`
- `TOK_SEP2` (`01`) is represented as `,`
- `TOK_VARINT` (`100`) is represented as the integer value (e.g. `123`)
- `TOK_VARBIT` (`110`) is represented as the integer value (e.g. `123`)
- `TOK_PART` (`101`) is represented as:
  - `{index}` for SUBTYPE_NONE
  - `{index:value}` for SUBTYPE_INT
  - `{index:[value1 value2 ...]}` for SUBTYPE_LIST

This allows to easily read and write item data in a textual format, without knowing the underlying structures used by the game engine.

- **Note 1:** VARINT and VARBIT are ambiguous, they use the same string  representation.  
This is because the game engine accepts both interchangeably:   
It collapses from one format to the other depending on which is the shortest during serialization.

- **Note 2:** The same logic should be applied when serializing from string representation.  
If not, the game may collapse VARINTs from/to VARBITs.  
(Not a big deal, result is the same, but it may surprise you if you expected your serial to stay the same.)

  
## Complete deserialization example

The serial below has been deserialized into its bitstream and string representation.  
For clarity, the bitstream has also been splitted into its tokens and data blocks.

```
Serial:
  @Ugy3L+2}TYg%$yC%i7M2gZldO)@}cgb!l34$a-qf{00

Splitted bitstream:
  0010000  1000001110000  01  10000000  01  10010000  01  1000100111000  00  10001000  01  100110011100110110  00  00  1010011100100010  10101000010  10111000010  1011101100100010  1011001111000010  1010011111000010  1011101111000010  1010000110000010  1011001110000010  1010011101000010  1011000111000010  00  0000000
  Header   Varint:24      ,   Varbit:0  ,   Varint:1  ,   Varint:50      |   Varint:2  |   Varint:3379         |   |   Part:76           Part: 2      Part:3       Part:75           Part:57           Part:60           Part:59           Part:16           Part:25           Part:44           Part:49           |   Padding

String representation:
  24, 0, 1, 50| 2, 3379|| {76} {2} {3} {75} {57} {60} {59} {16} {25} {44} {49}|
```

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
