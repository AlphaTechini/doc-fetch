# Contributing to DocFetch

Thank you for your interest in contributing to DocFetch! 🎉 This guide will help you get started.

## 📋 Quick Overview

1. **Create an issue first** - Always start by opening an issue to discuss your change
2. **Wait for feedback** - Maintainers will respond and provide guidance
3. **Fork and develop** - Once approved, fork the repo and make your changes
4. **Submit a PR** - Open a pull request referencing the issue
5. **Review process** - Maintainers will review and provide feedback
6. **Merge** - Once approved, your contribution will be merged!

---

## 🐛 Before You Start: Create an Issue

**⚠️ IMPORTANT: Always create an issue before submitting a PR!**

This helps us:
- Avoid duplicate work
- Discuss the best approach
- Ensure your contribution aligns with project goals
- Get early feedback from maintainers

### Types of Issues

#### 🐞 Bug Reports
Include:
- Clear description of the bug
- Steps to reproduce
- Expected vs actual behavior
- Environment details (OS, Go version, DocFetch version)
- Sample command that triggers the bug
- Error messages or logs

#### ✨ Feature Requests
Include:
- Clear description of the feature
- Use case / problem it solves
- Example usage
- Any relevant links or references

#### 📝 Documentation Improvements
Include:
- What needs improvement
- Why it's needed
- Suggested changes

#### ⚡ Performance Improvements
Include:
- Current performance metrics
- Proposed improvements
- Benchmark results (if available)

---

## 🚀 Development Setup

### Prerequisites

- Go 1.21 or later
- Git
- Make (optional, for running tests)

### Fork and Clone

```bash
# Fork the repository on GitHub, then:
git clone https://github.com/YOUR_USERNAME/doc-fetch.git
cd doc-fetch

# Add upstream remote
git remote add upstream https://github.com/AlphaTechini/doc-fetch.git
```

### Build from Source

```bash
# Build the binary
go build -o doc-fetch ./cmd/docfetch

# Test it works
./doc-fetch --help
```

### Run Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific package tests
go test ./pkg/fetcher/...
```

---

## 💻 Making Changes

### Branch Naming

Use descriptive branch names:
- `fix/content-extraction-bug`
- `feat/add-pdf-support`
- `docs/update-readme-examples`
- `perf/improve-concurrent-fetching`

### Code Style

Follow Go best practices:
- Run `go fmt` before committing
- Run `go vet` to catch issues
- Write clear, concise comments
- Keep functions small and focused
- Use meaningful variable names

### Testing Requirements

- Add tests for new features
- Ensure existing tests pass
- Include edge cases
- Test with real documentation sites

Example test:
```go
func TestContentExtraction(t *testing.T) {
    doc := createTestDocument()
    content := cleanContent(doc)
    
    if len(content) == 0 {
        t.Error("Expected content to be extracted")
    }
    
    if !strings.Contains(content, "expected text") {
        t.Error("Expected content to contain specific text")
    }
}
```

---

## 📤 Submitting a Pull Request

### PR Checklist

Before submitting your PR, ensure:

- [ ] You created an issue first and referenced it in the PR
- [ ] Your code follows Go style guidelines
- [ ] All tests pass (`go test ./...`)
- [ ] You've added tests for new functionality
- [ ] You've updated documentation if needed
- [ ] Your commit messages are clear and descriptive
- [ ] You've rebased on the latest main branch

### PR Template

When creating your PR, include:

```markdown
## Description
Brief description of changes

## Related Issue
Fixes #123 (or "Related to #123")

## Type of Change
- [ ] Bug fix
- [ ] New feature
- [ ] Breaking change
- [ ] Documentation update
- [ ] Performance improvement
- [ ] Refactoring

## Testing
Describe how you tested this:
- [ ] Unit tests added/updated
- [ ] Manual testing with real docs
- [ ] Tested on: [list platforms]

## Example Usage
Show example command and output if applicable

## Checklist
- [ ] Code follows project guidelines
- [ ] Self-review completed
- [ ] Comments added where needed
- [ ] Tests pass locally
```

---

## 🔍 Review Process

1. **Automated Checks**: CI runs tests and linting
2. **Maintainer Review**: At least one maintainer reviews
3. **Feedback**: You may be asked to make changes
4. **Approval**: Once approved, PR is merged
5. **Release**: Changes included in next release

Typical timeline: 3-7 days for review

---

## 📖 Contribution Ideas

Looking for ways to contribute? Here are some ideas:

### Easy Wins
- Fix typos in documentation
- Add more examples to README
- Improve error messages
- Add unit tests for existing code

### Intermediate
- Add support for new documentation site formats
- Improve content extraction selectors
- Add progress indicators
- Enhance LLM.txt generation

### Advanced
- Add PDF export support
- Implement incremental updates
- Add authentication support for private docs
- Create plugin system for custom extractors

---

## 🤝 Community Guidelines

### Be Respectful
- Treat everyone with respect
- Welcome newcomers
- Provide constructive feedback
- Assume good intentions

### Communication
- Use clear, concise language
- Explain your reasoning
- Ask questions if unsure
- Respond to feedback promptly

### Collaboration
- Work with maintainers, not against them
- Be open to suggestions
- Help other contributors
- Share knowledge

---

## ❓ Questions?

- **General questions**: Open a discussion on GitHub
- **Bug reports**: Create an issue
- **Feature requests**: Create an issue
- **Quick questions**: Check existing issues/discussions first

---

## 🙏 Thank You!

Your contributions make DocFetch better for everyone. Whether it's a typo fix, a new feature, or better documentation - we appreciate your time and effort!

Happy coding! 🚀
