import { expect, test } from 'vitest'
import { duplicateEncode } from './duplicate-encode'

test('Basic tests', () => {
  expect(duplicateEncode("din")).toBe("(((");
  expect(duplicateEncode("recede")).toBe("()()()");
  expect(duplicateEncode("Success")).toBe(")())())");
  expect(duplicateEncode("(( @")).toBe("))((");
});

