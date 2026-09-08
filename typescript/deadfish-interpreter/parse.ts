/** return the output array and ignore all non-op characters */
export function parse(data: string): number[] {
  let memory = 0;
  let output = [];
  const tokens = data.split('');

  for (let token of tokens) {
    switch (token) {
      case 'i':
        memory += 1;
        break;
      case 'd':
        memory -= 1;
        break;
      case 's':
        memory *= memory;
        break;
      case 'o':
        output.push(memory)
        break;
      default: continue;
    }
  }
  return output;
}

