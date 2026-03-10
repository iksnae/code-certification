// TypeScript interfaces mirroring Go JSON output types.
// These match the shapes produced by `certify report --format json` and `.certification/` artifacts.

export interface CertifyCard {
  repository: string;
  generated_at: string;
  commit_sha?: string;
  overall_grade: string;
  overall_score: number;
  pass_rate: number;
  total_units: number;
  passing: number;
  failing: number;
  expired: number;
  observations: number;
  grade_distribution: Record<string, number>;
  languages: LanguageDetail[];
  top_issues?: IssueCard[];
}

export interface IssueCard {
  unit_id: string;
  grade: string;
  score: number;
  reason: string;
}

export interface FullReport {
  repository: string;
  commit_sha?: string;
  generated_at: string;
  card: CertifyCard;
  units: UnitReport[];
  dimension_averages: Record<string, number>;
  language_detail: LanguageDetail[];
}

export interface UnitReport {
  unit_id: string;
  unit_type: string;
  path: string;
  language: string;
  symbol?: string;
  status: string;
  grade: string;
  score: number;
  confidence: number;
  dimensions: Record<string, number>;
  observations?: string[];
  actions?: string[];
  certified_at: string;
  expires_at: string;
  source: string;
}

export interface LanguageDetail {
  name: string;
  units: number;
  passing: number;
  average_score: number;
  grade: string;
  grade_distribution: Record<string, number>;
  top_score: number;
  bottom_score: number;
}

export interface IndexEntry {
  id: string;
  language: string;
  path: string;
  type: string;
}

export interface RecordJSON {
  unit_id: string;
  unit_type: string;
  unit_path: string;
  policy_version: string;
  status: string;
  grade: string;
  score: number;
  confidence: number;
  dimensions?: Record<string, number>;
  observations?: string[];
  actions?: string[];
  certified_at: string;
  expires_at: string;
  source: string;
  version: number;
}

export interface BadgeJSON {
  schemaVersion: number;
  label: string;
  message: string;
  color: string;
  style?: string;
  namedLogo?: string;
  logoColor?: string;
}

export interface ModelInfo {
  id: string;
  owned_by?: string;
  context_window?: number;
  created?: number;
}

export interface CertifyConfig {
  mode: string;
  scope?: { include?: string[]; exclude?: string[] };
  agent?: AgentConfig;
  expiry?: { default_window_days: number; min_window_days: number; max_window_days: number };
  issues?: { enabled: boolean; labels?: string[]; grouping?: string };
}

export interface AgentConfig {
  enabled?: boolean;
  provider?: ProviderConfig;
  models?: ModelAssignments;
  rate_limit?: { requests_per_minute: number; retry_max: number; retry_backoff_base_ms: number };
}

export interface ProviderConfig {
  type?: string;
  base_url?: string;
  api_key_env?: string;
  http_referer?: string;
  x_title?: string;
}

export interface ModelAssignments {
  prescreen?: string;
  review?: string;
  scoring?: string;
  decision?: string;
  remediation?: string;
  fallback?: string;
}

export interface ProviderPreset {
  name: string;
  baseURL: string;
  apiKeyEnvVar: string;
  local: boolean;
  description: string;
}
