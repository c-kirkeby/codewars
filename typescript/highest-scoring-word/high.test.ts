import { expect, test } from 'vitest'
import { high } from './high'

const solutions = [
  ['man i need a taxi up to ubud', 'taxi'],
  ['what time are we climbing up the volcano', 'volcano'],
  ['take me to semynak', 'semynak'],
  ['massage yes massage yes massage', 'massage'],
  ['take two bintang and a dance please', 'bintang'],
  ['aa b', 'aa'],
  ['b aa', 'b'],
  ['bb d', 'bb'],
  ['d bb', 'd'],
  ['aaa b', 'aaa'],
]

test('Basic tests', () => {
  solutions.forEach(([input, expected]) => {
    expect(high(input)).toStrictEqual(expected)
  })
});

