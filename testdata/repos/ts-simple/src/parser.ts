export function tokenizeDialogue(input: string): string[] {
  return input.split('\n');
}

export function parseNode(token: string): { type: string; value: string } {
  return { type: 'text', value: token };
}

export class DialogueParser {
  private tokens: string[] = [];

  parse(input: string): void {
    this.tokens = tokenizeDialogue(input);
  }
}

export const MAX_TOKENS = 1000;
