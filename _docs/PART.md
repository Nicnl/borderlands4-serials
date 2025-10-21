# PART

The PART encodes a complex structure representing an item part, linking an index to either no value, a single integer value, or a list of integer values.

- Major type is: `101`
- Contains an index followed by optional subtype data

## Structure

### 1. Index (VARINT without token)
After the `101` token, a VARINT is read **without** its `100` token prefix.  
This represents the part index.

### 2. Subtype Flag (1 bit)
A single bit determines the basic subtype:

- **`1`** → SUBTYPE_INT pathway
- **`0`** → Read 2 more bits for further classification

### 3. Subtype Data

#### SUBTYPE_INT (flag = `1`)
```
Type  Index      Flag  Value              Term
vvv   vvvvv...   v     vvvvv...           vvv
101   [varint]   1     [varint w/o token] 000
```

- Contains a VARINT value (without its `100` token)
- Ends with `000` terminator
- Example: `{76:4}` represents index 76 with integer value 4

#### SUBTYPE_NONE (flag = `0`, next 2 bits = `10`)
```
Type  Index      Flag
vvv   vvvvv...   vvv
101   [varint]   010
```

- No additional data after the index
- Example: `{76}` represents index 76 with no value

#### SUBTYPE_LIST (flag = `0`, next 2 bits = `01`)
```
Type  Index      Flag  Start  Values...              End
vvv   vvvvv...   vvv   vv     vvv vvvvv... ...       vv
101   [varint]   001   01     100 [varint] ...       00
                              110 [varbit] ...
```

- List begins with `01` token (TOK_SEP2)
- Followed by a sequence of values, each prefixed with either:
  - `100` → VARINT value
  - `110` → VARBIT value
- List ends with `00` token (TOK_SEP1)
- Example: `{54:[3 2 1]}` represents index 54 with list values [3, 2, 1]


## Examples

### Example 1: Simple part (SUBTYPE_NONE)
```
  Type  Index         Flag
  vvv   vvvvvvvvvv    vvv
  101   0011100100    010
        =varint:76
```
Result: `{76}`

### Example 2: Integer subtype (SUBTYPE_INT)
```
Type  Index         Flag  Value     Term
vvv   vvvvvvvvvv    v     vvvvv     vvv
101   0011100100    1     00100     000
      ^^^^^^^^^^          ^^^^^
      =varint:76          =varint:4
```
Result: `{76:4}`

### Example 3: List subtype (SUBTYPE_LIST)
```
Type  Index         List: Start  Elem 1      Elem 1      Elem 3      End
vvv   vvvvvvvvvv    vvv   vv     vvvvvvvvv   vvvvvvvvv   vvvvvvvvv
101   0110111000    001   01     100 11000   100 01000   100 10000   00
      ^^^^^^^^^^                     ^^^^^       ^^^^^       ^^^^^
      =varint:54                     varint      varint      varint
                                     =3          =2          =3
```
Result: `{54:[3 2 1]}`

## String Representation

- **SUBTYPE_NONE**: `{index}`
- **SUBTYPE_INT**: `{index:value}`
- **SUBTYPE_LIST**: `{index:[value1 value2 value3]}`
