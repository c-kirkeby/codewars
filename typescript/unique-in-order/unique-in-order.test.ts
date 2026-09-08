import { expect, test } from 'vitest'
import { uniqueInOrder } from './unique-in-order.js'

test('Basic tests', () => {
  expect(uniqueInOrder('AAAABBBCCDAABBB')).toStrictEqual(['A', 'B', 'C', 'D', 'A', 'B']);
  expect(uniqueInOrder([1, 2, 3, 3])).toStrictEqual([1, 2, 3])
});

