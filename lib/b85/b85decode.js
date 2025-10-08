const B85_CHARSET = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!#$%&()*+-;<=>?@^_`{/}~";

const reverseLookup = new Array(256).fill(0xFF);
for (let i = 0; i < B85_CHARSET.length; i++) {
    reverseLookup[B85_CHARSET.charCodeAt(i)] = i;
}


function b85Decode(string) {
    if (string[0] !== '@' || string[1] !== 'U') return null;
    string = string.slice(2);

    // Initialize reverse lookup table
    const result = [];
    let idx = 0;
    const size = string.length;

    while (idx < size) {
        let workingU32 = 0;
        let charCount = 0;

        // Collect up to 5 valid Base85 characters
        while (idx < size && charCount < 5) {
            const charCode = string.charCodeAt(idx);
            idx++;

            if (charCode >= 0 && reverseLookup[charCode] < 0x56) {
                workingU32 = workingU32 * 85 + reverseLookup[charCode];
                charCount++;
            }
        }

        if (charCount === 0) break;

        // Handle padding for incomplete groups
        if (charCount < 5) {
            const padding = 5 - charCount;
            for (let i = 0; i < padding; i++) {
                workingU32 = workingU32 * 85 + 0x7e; // '~' value
            }
        }

        if (charCount === 5) {
            // Full group - apply byte order transformation
            const standardBytes = [
                (workingU32 >>> 24) & 0xFF,
                (workingU32 >>> 16) & 0xFF,
                (workingU32 >>> 8) & 0xFF,
                (workingU32 >>> 0) & 0xFF
            ];

            const orderedBytes = reverseByteOrder(standardBytes);
            result.push(orderedBytes[0], orderedBytes[1], orderedBytes[2], orderedBytes[3]);
        } else {
            // Partial group - NO byte order transformation, just extract bytes normally
            const byteCount = charCount - 1;
            if (byteCount >= 1) result.push((workingU32 >>> 24) & 0xFF);
            if (byteCount >= 2) result.push((workingU32 >>> 16) & 0xFF);
            if (byteCount >= 3) result.push((workingU32 >>> 8) & 0xFF);
        }
    }

    // Convert byte array to hex string
    return result.map(b => b.toString(16).padStart(2, '0')).join('');
}
