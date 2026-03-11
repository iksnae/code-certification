import type { ProviderPreset } from './types.js';

// Brand colors from docs/brand.md
export const BRAND_COLORS = {
  graphite: '#1C1C1C',
  slate: '#2C3A40',
  steelBlue: '#4A6B82',
  certified: '#2E8B57',
  observations: '#E0A100',
  probationary: '#F59E0B',
  expired: '#9CA3AF',
  decertified: '#DC2626',
  lightBg: '#F8FAFC',
  mutedBg: '#EEF2F5',
} as const;

// Grade → color mapping (for badges, charts)
export const GRADE_COLORS: Record<string, string> = {
  'A': '#2E8B57',
  'A-': '#3DA06A',
  'B+': '#4A6B82',
  'B': '#4A6B82',
  'C': '#E0A100',
  'D': '#F59E0B',
  'F': '#DC2626',
};

// Grade → emoji mapping
export const GRADE_EMOJI: Record<string, string> = {
  'A': '🟢', 'A-': '🟢',
  'B+': '🟢', 'B': '🟢',
  'C': '🟡',
  'D': '🟠',
  'F': '🔴',
  'expired': '⚪',
};

// All 9 quality dimensions
export const DIMENSION_NAMES = [
  'correctness',
  'maintainability',
  'readability',
  'testability',
  'security',
  'architectural_fitness',
  'operational_quality',
  'performance_appropriateness',
  'change_risk',
] as const;

// Deep analysis metric labels (Sprint 7-13: type-aware cross-file analysis)
export const DEEP_METRIC_LABELS: Record<string, string> = {
  fan_in: 'Fan-In (callers)',
  fan_out: 'Fan-Out (callees)',
  is_dead_code: 'Dead Export',
  dep_depth: 'Dep Depth',
  instability: 'Instability',
  concrete_deps: 'Concrete Deps',
  coupling_score: 'Coupling',
  unused_params: 'Unused Params',
  interface_size: 'Interface Size',
  type_aware_unwrapped: 'Unwrapped Errors',
  cognitive_complexity: 'Cognitive Complexity',
  errors_not_wrapped: 'Errors Not Wrapped',
  unsafe_import_count: 'Unsafe Imports',
  hardcoded_secrets: 'Hardcoded Secrets',
} as const;

// Provider presets — convenience shortcuts, not limitations.
// Any OpenAI-compatible endpoint works via "Custom".
export const PROVIDER_PRESETS: ProviderPreset[] = [
  {
    name: 'OpenRouter',
    baseURL: 'https://openrouter.ai/api/v1',
    apiKeyEnvVar: 'OPENROUTER_API_KEY',
    local: false,
    description: 'Access 200+ models including GPT-4o, Claude, Gemini, Llama, Qwen, Mistral',
  },
  {
    name: 'Groq',
    baseURL: 'https://api.groq.com/openai/v1',
    apiKeyEnvVar: 'GROQ_API_KEY',
    local: false,
    description: 'Ultra-fast inference — Llama, Gemma, Mixtral (generous free tier)',
  },
  {
    name: 'Together',
    baseURL: 'https://api.together.xyz/v1',
    apiKeyEnvVar: 'TOGETHER_API_KEY',
    local: false,
    description: 'Llama, Qwen, DeepSeek, Mixtral',
  },
  {
    name: 'Fireworks',
    baseURL: 'https://api.fireworks.ai/inference/v1',
    apiKeyEnvVar: 'FIREWORKS_API_KEY',
    local: false,
    description: 'Llama, Qwen, Mixtral, custom fine-tunes',
  },
  {
    name: 'OpenAI',
    baseURL: 'https://api.openai.com/v1',
    apiKeyEnvVar: 'OPENAI_API_KEY',
    local: false,
    description: 'GPT-4o, GPT-4o-mini, o1',
  },
  {
    name: 'Google AI Studio',
    baseURL: 'https://generativelanguage.googleapis.com/v1beta/openai',
    apiKeyEnvVar: 'GEMINI_API_KEY',
    local: false,
    description: 'Gemini 2.0 Flash, Gemini Pro',
  },
  {
    name: 'Ollama',
    baseURL: 'http://localhost:11434/v1',
    apiKeyEnvVar: '',
    local: true,
    description: 'Run any model locally — Qwen, Llama, Phi, Gemma, DeepSeek',
  },
  {
    name: 'LM Studio',
    baseURL: 'http://localhost:1234/v1',
    apiKeyEnvVar: '',
    local: true,
    description: 'Local model server with GUI model management',
  },
  {
    name: 'vLLM',
    baseURL: 'http://localhost:8000/v1',
    apiKeyEnvVar: '',
    local: true,
    description: 'High-performance local inference server',
  },
  {
    name: 'Custom',
    baseURL: '',
    apiKeyEnvVar: '',
    local: false,
    description: 'Any OpenAI-compatible API endpoint',
  },
];
