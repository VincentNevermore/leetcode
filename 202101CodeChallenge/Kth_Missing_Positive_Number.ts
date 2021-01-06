// 2021 Jan Week 1 Challenge
// Author: Shengwei Zhang

function findKthPositive(arr: number[], k: number): number {
    let count = arr[0] - 1;

    if (arr == null || count >= k) {
        return k;
    }

    if (arr.length == 1 && count < k) {
        return arr[0] + k - count;
    }
    for (let i = 1; i < arr.length; i++) {
        if (count + arr[i] - arr[i - 1] - 1 - k < 0) {
            count = count + arr[i] - arr[i - 1] - 1;
        } else {
            return arr[i - 1] + k - count;
        }
        if (i === arr.length - 1 && count < k) {
            return arr[i] + k - count;
        }
    }
}

// Some special Test Cases
findKthPositive([1], 1);
findKthPositive([2], 1);
findKthPositive([], 100);