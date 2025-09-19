# Contributing to DevToolBox

Thank you for your interest in contributing to DevToolBox! This guide will help you get started.

## 🤝 How to Contribute

### Types of Contributions

- 🐛 **Bug Reports**: Report issues you've found
- 💡 **Feature Requests**: Suggest new features
- 📝 **Documentation**: Improve docs and examples
- 🔧 **Code**: Fix bugs or implement features
- 🔌 **Plugins**: Create new code generators
- 🧪 **Tests**: Add or improve test coverage

## 🚀 Getting Started

### 1. Fork and Clone

```bash
# Fork the repository on GitHub, then clone
git clone https://github.com/YOUR_USERNAME/devtoolbox.git
cd devtoolbox

# Add upstream remote
git remote add upstream https://github.com/JIIL07/devtoolbox.git
```

### 2. Set Up Development Environment

```bash
# Install dependencies
make install-deps

# Start development environment
make docker-dev-up

# Verify everything works
make test
```

### 3. Create a Branch

```bash
# Create feature branch
git checkout -b feature/your-feature-name

# Or bugfix branch
git checkout -b bugfix/issue-number
```

## 📝 Development Guidelines

### Code Style

#### Go Code

- Follow [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- Use `gofmt` for formatting
- Add comments for exported functions
- Write tests for new functionality

```bash
# Format Go code
go fmt ./...

# Run linter
golangci-lint run
```

#### TypeScript/React Code

- Follow [React Style Guide](https://react.dev/learn/thinking-in-react)
- Use TypeScript strict mode
- Add PropTypes or TypeScript interfaces
- Write component tests

```bash
# Format frontend code
cd frontend && npm run format

# Run linter
cd frontend && npm run lint
```

#### Python Code

- Follow [PEP 8](https://pep8.org/)
- Use type hints where possible
- Add docstrings for functions
- Write unit tests

```bash
# Format Python code
black plugins/ tests/python/

# Run linter
flake8 plugins/ tests/python/
```

### Testing

#### Write Tests

- **Go**: Add tests in `tests/go/`
- **Frontend**: Add tests in `frontend/src/components/__tests__/`
- **Python**: Add tests in `tests/python/`

#### Test Coverage

```bash
# Run all tests
make test

# Run specific test suites
go test ./internal/core/...
cd frontend && npm test
python -m pytest tests/python/
```

### Documentation

- Update relevant documentation files
- Add examples for new features
- Update API documentation if needed
- Include screenshots for UI changes

## 🔌 Creating Plugins

### Plugin Structure

```python
def generate(input_json: str) -> str:
    """
    Generate code from JSON input.
    
    Args:
        input_json: JSON string containing the schema
        
    Returns:
        Generated code as string
    """
    # Implementation here
    pass
```

### Plugin Guidelines

1. **Error Handling**: Always handle invalid input gracefully
2. **Documentation**: Add clear docstrings and comments
3. **Testing**: Write comprehensive tests
4. **Examples**: Provide usage examples
5. **Naming**: Use descriptive, consistent names

### Submitting Plugins

1. Create plugin in `plugins/community/`
2. Add tests in `tests/python/`
3. Update documentation
4. Submit pull request

## 🐛 Reporting Issues

### Bug Reports

Use the [GitHub Issues](https://github.com/JIIL07/devtoolbox/issues) template:

```markdown
**Bug Description**
A clear description of the bug.

**Steps to Reproduce**
1. Go to '...'
2. Click on '....'
3. See error

**Expected Behavior**
What you expected to happen.

**Actual Behavior**
What actually happened.

**Environment**
- OS: [e.g., Windows 10]
- Go Version: [e.g., 1.23]
- Node Version: [e.g., 18.17.0]
- Python Version: [e.g., 3.10.0]

**Additional Context**
Any other context about the problem.
```

### Feature Requests

```markdown
**Feature Description**
A clear description of the feature.

**Use Case**
Why would this feature be useful?

**Proposed Solution**
How should this feature work?

**Alternatives**
Any alternative solutions you've considered.

**Additional Context**
Any other context about the feature request.
```

## 🔄 Pull Request Process

### 1. Prepare Your Changes

```bash
# Make your changes
# Add tests
# Update documentation

# Run tests
make test

# Format code
make format
```

### 2. Commit Changes

```bash
# Use conventional commits
git add .
git commit -m "feat: add new code generator for Java classes"
git commit -m "fix: resolve plugin loading issue"
git commit -m "docs: update installation guide"
```

### 3. Push and Create PR

```bash
# Push to your fork
git push origin feature/your-feature-name

# Create pull request on GitHub
```

### 4. PR Template

```markdown
## Description
Brief description of changes.

## Type of Change
- [ ] Bug fix
- [ ] New feature
- [ ] Breaking change
- [ ] Documentation update

## Testing
- [ ] Tests pass locally
- [ ] Added new tests
- [ ] Updated existing tests

## Checklist
- [ ] Code follows style guidelines
- [ ] Self-review completed
- [ ] Documentation updated
- [ ] No breaking changes (or documented)
```

## 📋 Review Process

### What We Look For

- ✅ **Functionality**: Does it work as expected?
- ✅ **Code Quality**: Is the code clean and maintainable?
- ✅ **Tests**: Are there adequate tests?
- ✅ **Documentation**: Is documentation updated?
- ✅ **Performance**: Any performance implications?
- ✅ **Security**: Any security concerns?

### Review Timeline

- **Initial Review**: Within 2-3 business days
- **Follow-up Reviews**: Within 1-2 business days
- **Merge**: After approval and CI passes

## 🏷️ Release Process

### Versioning

We use [Semantic Versioning](https://semver.org/):

- **MAJOR**: Breaking changes
- **MINOR**: New features (backward compatible)
- **PATCH**: Bug fixes (backward compatible)

### Release Schedule

- **Patch Releases**: As needed for bug fixes
- **Minor Releases**: Monthly for new features
- **Major Releases**: Quarterly for breaking changes

## 🎯 Development Roadmap

### Current Priorities

- [ ] Plugin marketplace
- [ ] Real-time collaboration
- [ ] Advanced code generation
- [ ] Performance optimizations
- [ ] Mobile app

### Long-term Goals

- [ ] Multi-language support
- [ ] Cloud deployment
- [ ] Enterprise features
- [ ] AI-powered generation

## 💬 Community

### Communication Channels

- 💬 [GitHub Discussions](https://github.com/JIIL07/devtoolbox/discussions)
- 🐛 [GitHub Issues](https://github.com/JIIL07/devtoolbox/issues)
- 📧 Email: your-email@example.com

### Code of Conduct

We follow the [Contributor Covenant](https://www.contributor-covenant.org/):

- Be respectful and inclusive
- Welcome newcomers
- Focus on constructive feedback
- Respect different viewpoints

## 🏆 Recognition

### Contributors

Contributors are recognized in:

- README.md contributors section
- Release notes
- Annual contributor highlights

### Types of Contributions

- **Code Contributors**: Direct code contributions
- **Documentation Contributors**: Documentation improvements
- **Community Contributors**: Help with issues and discussions
- **Plugin Contributors**: Community plugins

## 📚 Resources

### Development Resources

- [Go Documentation](https://golang.org/doc/)
- [React Documentation](https://react.dev/)
- [Python Documentation](https://docs.python.org/)
- [Docker Documentation](https://docs.docker.com/)

### Project Resources

- [Architecture Overview](architecture.md)
- [API Documentation](api-reference.md)
- [Plugin Development](plugins-guide.md)
- [Testing Guide](testing-guide.md)

## ❓ Questions?

If you have questions about contributing:

1. Check existing [GitHub Discussions](https://github.com/JIIL07/devtoolbox/discussions)
2. Search [GitHub Issues](https://github.com/JIIL07/devtoolbox/issues)
3. Create a new discussion or issue
4. Contact maintainers directly

Thank you for contributing to DevToolBox! 🎉
