function numberOfArithmeticSlices(nums: number[]): number {
  if (nums.length < 3) {
    return 0;
  }
  let cnt = 0;
  let consecutiveArithmaticSubArr = -1;

  for (let i = 2; i < nums.length; i++) {
    const diff1 = nums[i] - nums[i - 1];
    const diff2 = nums[i - 1] - nums[i - 2];

    if (diff1 === diff2) {
      cnt++;
      consecutiveArithmaticSubArr++;
      cnt += consecutiveArithmaticSubArr;
      continue;
    }
    consecutiveArithmaticSubArr = -1;
  }
  return cnt;
}