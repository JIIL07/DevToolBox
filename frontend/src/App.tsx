import React from 'react';
import { GeneratorForm } from './components/GeneratorForm';
import { Header } from './components/Header';
import { Footer } from './components/Footer';

function App() {
  return (
    <div className="min-h-screen flex flex-col">
      <Header />
      <main className="flex-1 container mx-auto px-4 py-8">
        <div className="max-w-4xl mx-auto">
          <div className="text-center mb-8">
            <h1 className="text-4xl font-bold text-gray-900 mb-4">
              DevToolBox
            </h1>
            <p className="text-xl text-gray-600">
              Generate code from JSON schemas with ease
            </p>
          </div>
          <GeneratorForm />
        </div>
      </main>
      <Footer />
    </div>
  );
}

export default App;
