export function formatDate(date: Date): string {
  return date.toISOString();
}

export default function log(message: string): void {
  console.log(message);
}
