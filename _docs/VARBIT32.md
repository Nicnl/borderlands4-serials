# VARBIT32

The varbit32 encodes integers up to 32-bits into a length-prefixed bitstream.

- <span style="color: rgb(22, 145, 121);">Confirmed (can be swapped with Varint16, notably in the item level)</span>
- Major type is: `110`
- Is followed by <span style="text-decoration: underline;">exactly one</span> block of 5-bits.
    - =&gt; length of the payload ahead. (LSB-first)
- Is followed by a variable amount of bits =&gt; payload.
    - Payload is ordered LSB-first.

### Example 1

```
Type   Length   Data
vvv    vvvvv    vvv
110    11000    011
       =3       ^^^
                =int:6
```

### Example 2

```
Type   Length   Data
vvv    vvvvv    vvvvvvvvvv
110    01010    0000000001
       =10      ^^^^^^^^^^
                =int:512
```
