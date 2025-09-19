export interface GenerateRequest {
  template: string;
  input: string;
}

export interface GenerateResponse {
  code: string;
  error?: string;
}

export interface GeneratorInfo {
  name: string;
  description: string;
}

export interface ListGeneratorsResponse {
  generators: GeneratorInfo[];
}

export interface HealthResponse {
  status: string;
  service: string;
}
