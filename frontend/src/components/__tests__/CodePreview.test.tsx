import { render, screen } from '@testing-library/react';
import { CodePreview } from '../CodePreview';

describe('CodePreview', () => {
  it('renders placeholder when no code is provided', () => {
    render(<CodePreview code="" />);

    expect(screen.getByText('No code generated yet')).toBeInTheDocument();
    expect(screen.getByText('Enter JSON input and click "Generate Code" to see the result')).toBeInTheDocument();
  });

  it('renders code when provided', () => {
    const testCode = 'type Test struct { Name string }';
    render(<CodePreview code={testCode} />);

    expect(screen.getByText(testCode)).toBeInTheDocument();
  });

  it('applies proper styling to code block', () => {
    const testCode = 'test code';
    render(<CodePreview code={testCode} />);

    const codeElement = screen.getByText('test code');
    expect(codeElement.closest('pre')).toHaveClass('bg-gray-900', 'text-gray-100', 'font-mono');
  });
});
