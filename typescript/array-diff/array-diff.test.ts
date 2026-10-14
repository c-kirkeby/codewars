import { expect, test } from 'vitest'
import { arrayDiff } from './array-diff';

test('Basic test should work', () => {
  expect(arrayDiff([], [4, 5])).toStrictEqual([]);
  expect(arrayDiff([3, 4], [3])).toStrictEqual([4]);
  expect(arrayDiff([1, 8, 2], [])).toStrictEqual([1, 8, 2]);
  expect(arrayDiff([1, 2, 3], [1, 2])).toStrictEqual([3]);
});
