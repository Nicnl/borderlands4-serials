# Journal

## Reverse Engineering Journal

### Preamble

It all began with Base85-encoded strings from Borderlands 4 that appeared to represent item data - weapons, shields, and more.  
The strings didn’t line up with the actual in-game data, so they had to be converted into bits and bytes first.  
The fact that @Sparkie just posted a script to convert these Base85-encoded strings into binary and us getting sick of uncontrolled LLM shenanigans started this journey.

### Phase 1 - First data: level

Comparing different items in binary or hex didn’t reveal much, as all items were too different from one another.  
So we began flipping all bits in the serial data individual and observed how the item changed.

This “bit-swapping” process created a lot of broken weapons, but also led to meaningful discoveries - such as item level, parts, and type.  
By extracting functional weapons and repeatedly marking and swapping bits, we eventually identified the field that controlled item level.

### Phase 2 - Varints and the bitstream

Between the level bits, we found a stray “1” bit, which hinted at the existence of varints.  
With this discovery, most weapons could have their levels modified successfully - though not all items behaved consistently. Still, it was progress.

We then tried locating other fields by bit-swapping and managed to reliably alter the scope on a sniper rifle.  
However, the bit locations weren’t always byte-aligned, strengthening our theory that the data was organized as a single bitstream rather than fixed-width parts.

### Phase 3 - The selling accident

While testing, @Nicnl accidentally sold the wrong gun to a vendor and then bought it back.  
Surprisingly, the serial changed completely - even becoming longer, as if new data had been added.  
We confirmed this reproducibly: about 20 bits were added upon buyback, shifting the bitstream and altering the entire serial.

Viewing the data in binary revealed widespread changes, even though the item itself should have been identical.  
This made us suspect that the Base85 decoding wasn’t quite correct - close, but not perfect.

Repeating the scope swap on a bought-back sniper revealed scope bits in two separate locations.  
This suggested the bitstream was shifting across byte boundaries and end up in 2 places due to incorrect endianness.

### Phase 4 - Fixing base85 to binary

Our theory was that buyback added data in one or two places within the serial, shifting - but not rewriting - most of the data.  
The item header remained unchanged, letting us pinpoint where the new bits were inserted. Beyond that point, we observed familiar patters from before the buyback but also seemingly random changes across the serial.

We ruled out encryption, hashing and dependencies on time, location or player state, so we started flipping endianness (little ↔ big) during decoding. This improved consistency but didn’t eliminate the mismatches. We could clearly see repeating structural mismatches: five nibbles identical, followed by three different ones - over and over.  
It wasn’t random; something systematic was off.

The breakthrough came when examining hex again.  
Twenty added bits should shift 5 hexadecimal characters, but instead, 3 shifted while 5 remained - a reversal that could only happen if the bit order itself was mirrored.  
After mirroring each byte and testing again, we found the correct combination of bit order and byte endianness, clearly showing the inserted data and the stable surrounding bitstream.

### Phase 5 - Item type and manufacturer

With the corrected binary, comparing items and bit-swapping became straightforward.  
We discovered that, for guns, item type and manufacturer were stored together in the header as a single varint.  
However, for non-gun items, the header used a standard 8-bit integer instead.

This led to an early but incorrect theory of “Type A” and “Type B” serials.  
Curiously, the transition occurred between class mods: Vex and Amon used varints (values 254 and 255), while Rafa and Harlow used normal integers (0 and 3).  
We couldn’t explain this discrepancy yet.

### Phase 6 - Closing the source

Up until around this point, we had worked completely in public.  
Our notes, the source code for the Base85-to-binary converter, all chat discussions - everything - was shared openly in the Discord channel or linked to a publicly available (but self-hosted) GitLab and BookStack instance.

Eventually, we decided to move to private messages, since we were flooding the channel with binary sequences and wild theories.  
However, we continued to report every major discovery in Discord, and the notes remained publicly accessible.

That changed when we discovered a random repository that had copied and pasted everything - chat messages, notes, theories, source code - without notice or permission.  
We objected, pointing out that this wasn’t cool and that the person could at least have asked first.  
At that time, many of our notes still contained incorrect assumptions, meaning this copy was effectively spreading misinformation.

A few people also seemed unhappy with our work, dismissing it entirely and posting nasty comments.  
To avoid further drama, we decided to make our sources and notes private, sharing only high-level progress updates without detailed information.

However, our Base85-to-binary conversion tool remained publicly available, and since others had already updated their own tools to use the corrected decoding method, there was no real harm in making the rest of our work private.

### Phase 7 - Discovering flag 100 and separators

With the improved decoder, earlier experiments became much more consistent.  
Expanding the level varint by setting its continuation bit and adding zeros still worked - the game simply collapsed it back.

We tested more serials to see which parts behaved as varints, especially in the header and on a purple Jakobs knife.  
Indeed, we found additional varints separated by a 100 binary pattern - likely a type flag.

Since the bitstream wasn’t byte-aligned, the original decoder must have included logic to identify where data structures began and ended.  
All identified varints were prefixed with that three-bit sequence.  
In the header, each prefixed varint was also followed by 01 or 00 - two bits appearing consistently only in headers, not part data.

Our theory was that these acted as separators, determining whether data was split into structures or arrays.  
At the time, it was just a hypothesis - later confirmed to be true.

### Phase 8 - Flag 110 and conversion

After mapping all the standard varints, we found another prefix in the header field for item type and manufacturer.  
This prefix (110) didn’t mark a varint but instead defined a variable-length bit field:  
Flag 110 + varint = number of bits to read as a large integer.

For non-gun items, this meant 9 bits. Suddenly, everything made sense:  
Vex and Amon still decoded as 254 and 255 via varint, but Rafa and Harlow became 256 and 259 as 9-bit integers.  
We confirmed this by replacing the level varint with the new “varbit” and changing the varbit in the manufacturer field to varint - and it worked flawlessly. The game just turned it back to its preferred form.

### Phase 9 - The stupid flag 101

With the previously discovered flags and separators, we could decode nearly everything - except parts.  
A new flag (101) appeared everywhere, making it impossible to tell where one part ended and another began.

Thankfully, the game’s error correction helped: it removed sections it couldn’t parse.  
We intentionally broke items by bit-swapping to trigger this, repeatedly corrupting the same class mod until only one part remained.  
By observing all possible bit combinations the last remaining part, we learned that flag 101 contained a varint and some extra bits.

Eventually, we determined that the full flag structure was about 11 bits long:  
3 bits for the flag, 5 for the varint, and 3 extra bits whose purpose was unclear.  
Typically, those extra bits were 010 or 001 when followed by perk lists (as in Jakobs knives).

This small breakthrough allowed us to decode around 90% of all items - but a few used flag 101 with unknown suffixes (110, 101, 111, 100).  
Further analysis revealed a varint following the suffixes, but with a misplaced continuation bit.  
After much frustration, we assumed it was a normal varint and treated the suffixes as a single flag bit:  
0 = no varint added; 1 = varint appended.  
When using 0, two bits were still reserved, implying 10 = no data following and 01 = list of data following.

It still feels imperfect, but every known case works with this model - so we’re sticking with it until proven otherwise.

### Phase 10 - Flag 111 and surrender

With functional decoding for flag 101, we were able to handle everything - except DLC items.  
Those introduced yet another prefix (111), which didn’t fit into any known flag scheme.

After ten days with very little sleep, we decided to give up on this part for the time being.  
Our plan was to gather more data and build a working item editor first - but even small ID mapping tasks were slow and exhausting.  
Feeling demotivated, we ultimately chose to release the decoding tool to the community, allowing others to figure out the remaining parts and create modded items without relying on the luck required by the LLM-based mutations.

The tool using REST and thus the source not being available was a byproduct of Phase 6, as well as all our scripts were written in Go, and we simply didn’t want to redo everything in JavaScript.  
We were also genuinely curious to see how others would solve this riddle and releasing our research publicly would have meant that no one would ever embark on this journey again.

### Phase 11 - Release and the Community Effort

The tool was released with minimal documentation, yet people immediately began experimenting.  
Players could now edit items freely - crafting perfect god rolls or bizarre abominations.

Before long, the community organized a spreadsheet to map part IDs and combinations.  
Some creations have a special place in my heart, such as "The Void", which broke audio and animation queues by endlessly spawning singularities - or the full skill-tree class mods, which rewrote what “balanced” could possibly mean.  
  
All that’s left to say is: thank you to everyone who supported us and stepped in to help when we were exhausted.  
Your encouragement and collaboration truly made us happy.