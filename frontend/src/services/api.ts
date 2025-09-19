import axios from 'axios';
import type { 
  GenerateRequest, 
  GenerateResponse, 
  ListGeneratorsResponse, 
  HealthResponse 
} from '../types/api';

const API_BASE_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080';

const api = axios.create({
  baseURL: API_BASE_URL,
  timeout: 10000,
  headers: {
    'Content-Type': 'application/json',
  },
});

export const apiService = {
  async generateCode(request: GenerateRequest): Promise<GenerateResponse> {
    const response = await api.post<GenerateResponse>('/generate', request);
    return response.data;
  },

  async getGenerators(): Promise<ListGeneratorsResponse> {
    const response = await api.get<ListGeneratorsResponse>('/generators');
    return response.data;
  },

  async checkHealth(): Promise<HealthResponse> {
    const response = await api.get<HealthResponse>('/health');
    return response.data;
  },
};

export default api;
