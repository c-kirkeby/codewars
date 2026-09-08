import { expect, test } from 'vitest'
import { parse } from './parse.js'

test('Basic tests', () => {
  expect(parse("iiisdoso")).toStrictEqual([8, 64]);
  expect(parse("iiisxxxdoso")).toStrictEqual([8, 64]);
});
