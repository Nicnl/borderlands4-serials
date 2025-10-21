# VARINT16

The varint16 encodes integers up to 16-bits into blocks of 5-bits, between one and four.

- <span style="color: rgb(22, 145, 121);">Confirmed (can be swapped with Varbit32, notably in the item level)</span>
- Major type is: `100`
- It is followed by at least one block of 5-bits:
    - LSB is the continuation bit:
        - 1 => there is another block ahead.  
          (Up to 3 additional blocks, 4 blocks total.)
        - 0 => end of payload.
    - The 4 remaining bits is part of the payload in LSB-first.
- All blocks have to be concatenated in the read order, and interpreted as LSB-first integer.

### Example 1: single-block VARINT
```
Type     Stop
v        v
100  11000
     ^^^^
     data
=bin:1100
=int:3
```

### Example 2: two-block VARINT

```
  Continue      Stop
         v      v
100  10001  00010
     ^^^^   ^^^^
     data   data
=bin:1000   0001
=int:129
```

### Example 3: four-block VARINT with useless zero-blocks

This example is special:   
The three continuation blocks are all zero, which is useless but valid.   
This is how we discovered this type, by injecting those zero-blocks and see if they'd collapse when reserialized by the game.

```
                              Stop
         v      v      v      v
100  01001  00001  00001  00000
     ^^^^   ^^^^   ^^^^   ^^^^
     data   data   data   data
=bin:0100   0000   0000   0000
=int:2
```
