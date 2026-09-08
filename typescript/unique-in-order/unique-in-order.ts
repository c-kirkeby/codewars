export function uniqueInOrder(iterable: string | (string | number)[]): (string | number)[] {
  const input = typeof iterable === 'string' ? iterable.split('') : iterable
  return input.filter((element, index) => element !== input[index + 1])
};
