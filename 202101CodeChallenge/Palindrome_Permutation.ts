// 2021 Jan Week 1 Challenge
// Author: Shengwei Zhang
// 思路： 首先考虑字符串的长度， 为了满足回文数的条件， 需要做到：
//       1） 如果字符串长度为偶数，则每个字符必须出现偶数次.
//       2)  如果字符串长度为奇数，则至多有一个字符出现奇数次
function canPermutePalindrome(s: string): boolean {
    let occurenceMap = new Map<string, number>();
    for (let i = 0; i < s.length; i++) {
        let c = s.charAt(i);
        if (occurenceMap.has(c)) {
            let count = occurenceMap.get(c) + 1;
            occurenceMap.set(c, count);
        } else {
            occurenceMap.set(c, 1);
        }
    }
    let occurenceCount = 0;
    occurenceMap.forEach((count, char) => {
        occurenceCount += count %2;
    })
    return occurenceCount <=1

}

// Test
canPermutePalindrome("code")