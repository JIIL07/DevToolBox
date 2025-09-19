import { render, screen, waitFor } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { vi } from 'vitest';
import { GeneratorForm } from '../GeneratorForm';
import { apiService } from '../../services/api';

vi.mock('../../services/api');

const mockApiService = vi.mocked(apiService);

describe('GeneratorForm', () => {
  const mockGenerators = [
    { name: 'go-struct', description: 'Generate Go structures' },
    { name: 'ts-interface', description: 'Generate TypeScript interfaces' },
  ];

  beforeEach(() => {
    mockApiService.getGenerators.mockResolvedValue({
      generators: mockGenerators,
    });
  });

  afterEach(() => {
    vi.clearAllMocks();
  });

  it('renders form elements', async () => {
    render(<GeneratorForm />);

    expect(screen.getByText('Input Configuration')).toBeInTheDocument();
    expect(screen.getByText('Generated Code')).toBeInTheDocument();
    expect(screen.getByRole('button', { name: /generate code/i })).toBeInTheDocument();
  });

  it('loads generators on mount', async () => {
    render(<GeneratorForm />);

    await waitFor(() => {
      expect(mockApiService.getGenerators).toHaveBeenCalledTimes(1);
    });
  });

  it('generates code when form is submitted', async () => {
    const mockResponse = {
      code: 'type GeneratedStruct struct {\n\tName string `json:"name"`\n}',
    };

    mockApiService.generateCode.mockResolvedValue(mockResponse);

    const user = userEvent.setup();
    render(<GeneratorForm />);

    await waitFor(() => {
      expect(screen.getByDisplayValue('go-struct - Generate Go structures')).toBeInTheDocument();
    });

    const textarea = screen.getByPlaceholderText('Enter your JSON data here...');
    const generateButton = screen.getByRole('button', { name: /generate code/i });

    await user.clear(textarea);
    await user.type(textarea, '{"name": "test"}');
    await user.click(generateButton);

    await waitFor(() => {
      expect(mockApiService.generateCode).toHaveBeenCalledWith({
        template: 'go-struct',
        input: '{"name": "test"}',
      });
    });

    expect(screen.getByText('type GeneratedStruct struct {')).toBeInTheDocument();
  });

  it('shows error when generation fails', async () => {
    mockApiService.generateCode.mockRejectedValue({
      response: { data: { error: 'Invalid JSON' } },
    });

    const user = userEvent.setup();
    render(<GeneratorForm />);

    await waitFor(() => {
      expect(screen.getByDisplayValue('go-struct - Generate Go structures')).toBeInTheDocument();
    });

    const textarea = screen.getByPlaceholderText('Enter your JSON data here...');
    const generateButton = screen.getByRole('button', { name: /generate code/i });

    await user.type(textarea, 'invalid json');
    await user.click(generateButton);

    await waitFor(() => {
      expect(screen.getByText('Invalid JSON')).toBeInTheDocument();
    });
  });

  it('loads example JSON when button is clicked', async () => {
    const user = userEvent.setup();
    render(<GeneratorForm />);

    await waitFor(() => {
      expect(screen.getByDisplayValue('go-struct - Generate Go structures')).toBeInTheDocument();
    });

    const loadExampleButton = screen.getByRole('button', { name: /load example/i });
    await user.click(loadExampleButton);

    const textarea = screen.getByPlaceholderText('Enter your JSON data here...');
    expect(textarea).toHaveValue(expect.stringContaining('John Doe'));
  });
});
