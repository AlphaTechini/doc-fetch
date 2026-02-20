# Security Policy

## Security Features

DocFetch includes several built-in security protections:

### ✅ Path Traversal Protection
- Output files can only be written within the current working directory
- Relative paths (`../`) are blocked
- Absolute paths outside the current directory are rejected

### ✅ SSRF (Server-Side Request Forgery) Protection  
- Only HTTP/HTTPS URLs are allowed
- Private IP addresses (192.168.x.x, 10.x.x.x, etc.) are blocked
- Localhost and loopback addresses are blocked
- Internal network access is prevented

### ✅ Rate Limiting
- Maximum 10 requests per second to avoid overwhelming servers
- Respectful crawling behavior

### ✅ Input Validation
- URL validation and sanitization
- Output path validation
- Parameter bounds checking (max depth: 10, max workers: 20)

### ✅ Content Safety
- HTML content is cleaned of scripts and dangerous elements
- XSS patterns are filtered out
- Only safe markdown is generated

## Safe Usage Guidelines

### Command Line Usage
```bash
# ✅ SAFE - relative path in current directory
doc-fetch --url https://example.com --output docs.md

# ✅ SAFE - subdirectory in current directory  
doc-fetch --url https://example.com --output ./docs/site.md

# ❌ BLOCKED - path traversal attempt
doc-fetch --url https://example.com --output ../../etc/passwd

# ❌ BLOCKED - absolute path outside current directory
doc-fetch --url https://example.com --output /tmp/malicious.md
```

### URL Restrictions
```bash
# ✅ SAFE - public HTTPS site
doc-fetch --url https://golang.org/doc/ --output docs.md

# ❌ BLOCKED - private IP address
doc-fetch --url http://192.168.1.1/admin --output docs.md

# ❌ BLOCKED - localhost
doc-fetch --url http://localhost:8080/api --output docs.md

# ❌ BLOCKED - non-HTTP protocol
doc-fetch --url file:///etc/passwd --output docs.md
```

## Reporting Security Issues

If you discover a security vulnerability in DocFetch, please:

1. **Do not disclose publicly** until it's been addressed
2. Contact the maintainer directly at [your email]
3. Provide detailed reproduction steps
4. Allow reasonable time for patch development

## Security Updates

Security patches will be released as soon as possible after vulnerability confirmation. Users are encouraged to keep DocFetch updated to the latest version.

## Dependencies Security

DocFetch uses the following dependencies with known security track records:
- `github.com/PuerkitoBio/goquery` - HTML parsing
- `github.com/yuin/goldmark` - Markdown processing  
- Standard Go libraries (`net/http`, `sync`, etc.)

All dependencies are regularly audited and kept up-to-date.