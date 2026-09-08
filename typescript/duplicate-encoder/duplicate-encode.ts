export function duplicateEncode(word: string) {
  return word
    .split('')
    .map(char => (word.match(new RegExp(escapeRegExp(char), 'gi')) || []).length > 1 ? ")" : "(").join('')
}

function escapeRegExp(str: string): string {
  return str.replace(/[\\^$.*+?()[\]{}|]/g, '\\$&');
}
