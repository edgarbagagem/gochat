export function onlyUnique(value: unknown, index: number, array: unknown[]) {
  return array.indexOf(value) === index;
}
