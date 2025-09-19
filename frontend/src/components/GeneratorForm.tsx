import React, { useState, useEffect } from 'react';
import { Play, Copy, Download, RefreshCw } from 'lucide-react';
import { apiService } from '../services/api';
import type { GeneratorInfo, GenerateRequest } from '../types/api';
import { CodePreview } from './CodePreview';
import { ErrorMessage } from './ErrorMessage';

export const GeneratorForm: React.FC = () => {
  const [generators, setGenerators] = useState<GeneratorInfo[]>([]);
  const [selectedTemplate, setSelectedTemplate] = useState<string>('');
  const [inputJson, setInputJson] = useState<string>('');
  const [generatedCode, setGeneratedCode] = useState<string>('');
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [error, setError] = useState<string>('');
  const [isGenerating, setIsGenerating] = useState<boolean>(false);

  useEffect(() => {
    loadGenerators();
  }, []);

  const loadGenerators = async () => {
    try {
      setIsLoading(true);
      const response = await apiService.getGenerators();
      setGenerators(response.generators);
      if (response.generators.length > 0) {
        setSelectedTemplate(response.generators[0].name);
      }
    } catch (err) {
      setError('Failed to load generators');
    } finally {
      setIsLoading(false);
    }
  };

  const handleGenerate = async () => {
    if (!selectedTemplate || !inputJson.trim()) {
      setError('Please select a template and provide JSON input');
      return;
    }

    try {
      setIsGenerating(true);
      setError('');
      
      const request: GenerateRequest = {
        template: selectedTemplate,
        input: inputJson.trim(),
      };

      const response = await apiService.generateCode(request);
      setGeneratedCode(response.code);
    } catch (err: any) {
      setError(err.response?.data?.error || 'Failed to generate code');
    } finally {
      setIsGenerating(false);
    }
  };

  const handleCopyCode = async () => {
    try {
      await navigator.clipboard.writeText(generatedCode);
    } catch (err) {
      console.error('Failed to copy code:', err);
    }
  };

  const handleDownloadCode = () => {
    const blob = new Blob([generatedCode], { type: 'text/plain' });
    const url = URL.createObjectURL(blob);
    const a = document.createElement('a');
    a.href = url;
    a.download = `generated-${selectedTemplate}.go`;
    document.body.appendChild(a);
    a.click();
    document.body.removeChild(a);
    URL.revokeObjectURL(url);
  };

  const exampleJson = `{
  "name": "John Doe",
  "age": 30,
  "email": "john@example.com",
  "active": true,
  "tags": ["developer", "golang"]
}`;

  return (
    <div className="space-y-6">
      <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <div className="card">
          <h2 className="text-xl font-semibold mb-4">Input Configuration</h2>
          
          <div className="space-y-4">
            <div>
              <label className="block text-sm font-medium text-gray-700 mb-2">
                Template
              </label>
              <select
                value={selectedTemplate}
                onChange={(e) => setSelectedTemplate(e.target.value)}
                className="input-field"
                disabled={isLoading}
              >
                {generators.map((gen) => (
                  <option key={gen.name} value={gen.name}>
                    {gen.name} - {gen.description}
                  </option>
                ))}
              </select>
            </div>

            <div>
              <label className="block text-sm font-medium text-gray-700 mb-2">
                JSON Input
              </label>
              <textarea
                value={inputJson}
                onChange={(e) => setInputJson(e.target.value)}
                placeholder="Enter your JSON data here..."
                className="input-field h-64 resize-none font-mono text-sm"
                disabled={isGenerating}
              />
            </div>

            <div className="flex space-x-3">
              <button
                onClick={handleGenerate}
                disabled={isGenerating || !selectedTemplate || !inputJson.trim()}
                className="btn-primary flex items-center space-x-2 disabled:opacity-50 disabled:cursor-not-allowed"
              >
                {isGenerating ? (
                  <RefreshCw className="h-4 w-4 animate-spin" />
                ) : (
                  <Play className="h-4 w-4" />
                )}
                <span>{isGenerating ? 'Generating...' : 'Generate Code'}</span>
              </button>

              <button
                onClick={() => setInputJson(exampleJson)}
                className="btn-secondary"
                disabled={isGenerating}
              >
                Load Example
              </button>
            </div>
          </div>
        </div>

        <div className="card">
          <div className="flex items-center justify-between mb-4">
            <h2 className="text-xl font-semibold">Generated Code</h2>
            {generatedCode && (
              <div className="flex space-x-2">
                <button
                  onClick={handleCopyCode}
                  className="p-2 text-gray-600 hover:text-gray-800 transition-colors"
                  title="Copy to clipboard"
                >
                  <Copy className="h-4 w-4" />
                </button>
                <button
                  onClick={handleDownloadCode}
                  className="p-2 text-gray-600 hover:text-gray-800 transition-colors"
                  title="Download code"
                >
                  <Download className="h-4 w-4" />
                </button>
              </div>
            )}
          </div>

          <CodePreview code={generatedCode} />
        </div>
      </div>

      {error && <ErrorMessage message={error} />}
    </div>
  );
};
