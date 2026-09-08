export function high(str: string) {
  const words = str.split(' ');
  const counts = words
    .map(
      element => element.split('').reduce(
        (sum, current) => sum += current.charCodeAt(0) - 96, 0
      )
    );
  return words[counts.indexOf(Math.max(...counts))]
}
