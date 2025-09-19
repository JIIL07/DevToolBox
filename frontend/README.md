# DevToolBox Frontend

React + TypeScript frontend for DevToolBox code generator.

## Features

- 🚀 Modern React 18 with TypeScript
- 🎨 Beautiful UI with Tailwind CSS
- 📱 Responsive design
- 🧪 Comprehensive testing with Vitest
- 🔧 Hot reload with Vite
- 📦 Optimized build

## Development

```bash
# Install dependencies
npm install

# Start development server
npm run dev

# Build for production
npm run build

# Run tests
npm test

# Run tests with UI
npm run test:ui

# Lint code
npm run lint
```

## Environment Variables

Create `.env.local` file:

```
VITE_API_URL=http://localhost:8080
```

## Project Structure

```
src/
├── components/          # React components
│   ├── __tests__/      # Component tests
│   ├── GeneratorForm.tsx
│   ├── CodePreview.tsx
│   ├── ErrorMessage.tsx
│   ├── Header.tsx
│   └── Footer.tsx
├── services/           # API services
│   └── api.ts
├── types/              # TypeScript types
│   └── api.ts
├── test/               # Test setup
│   └── setup.ts
├── App.tsx             # Main app component
├── main.tsx            # App entry point
└── index.css           # Global styles
```
