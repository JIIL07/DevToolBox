import React from 'react';
import { Code2 } from 'lucide-react';

export const Header: React.FC = () => {
  return (
    <header className="bg-white shadow-sm border-b border-gray-200">
      <div className="container mx-auto px-4 py-4">
        <div className="flex items-center space-x-3">
          <Code2 className="h-8 w-8 text-primary-600" />
          <h1 className="text-2xl font-bold text-gray-900">DevToolBox</h1>
        </div>
      </div>
    </header>
  );
};
