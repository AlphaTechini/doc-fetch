package fetcher

import (
	"fmt"
	"net"
	"net/url"
	"path/filepath"
	"strings"
)

// ValidateConfig validates the configuration for security issues
func ValidateConfig(config *Config) error {
	if err := validateURL(config.BaseURL); err != nil {
		return fmt.Errorf("invalid URL: %w", err)
	}
	if err := validateOutputPath(config.OutputPath); err != nil {
		return fmt.Errorf("invalid output path: %w", err)
	}
	if config.MaxDepth > 10 {
		return fmt.Errorf("max depth too high (maximum allowed: 10)")
	}
	if config.Workers > 20 {
		return fmt.Errorf("too many concurrent workers (maximum allowed: 20)")
	}
	return nil
}

// validateURL checks if the URL is safe to fetch
func validateURL(urlStr string) error {
	parsed, err := url.Parse(urlStr)
	if err != nil {
		return err
	}

	// Only allow HTTP and HTTPS
	if parsed.Scheme != "http" && parsed.Scheme != "https" {
		return fmt.Errorf("only HTTP and HTTPS URLs are allowed")
	}

	// Block private IP ranges and localhost
	host := parsed.Hostname()
	ip := net.ParseIP(host)
	if ip != nil {
		if ip.IsPrivate() || ip.IsLoopback() || ip.IsLinkLocalUnicast() || ip.IsMulticast() {
			return fmt.Errorf("access to private/internal IP addresses is not allowed")
		}
	}

	// Block dangerous hostnames
	dangerousHosts := []string{"localhost", "127.0.0.1", "0.0.0.0", "::1"}
	for _, dangerous := range dangerousHosts {
		if strings.ToLower(host) == dangerous {
			return fmt.Errorf("access to localhost is not allowed")
		}
	}

	return nil
}

// validateOutputPath ensures the output path is safe
func validateOutputPath(path string) error {
	// Don't allow absolute paths that start with /
	if strings.HasPrefix(path, "/") {
		return fmt.Errorf("absolute paths are not allowed")
	}

	// Don't allow paths that contain ..
	if strings.Contains(path, "..") {
		return fmt.Errorf("relative path traversal (..) is not allowed")
	}

	// Don't allow paths that contain ~
	if strings.Contains(path, "~") {
		return fmt.Errorf("home directory expansion (~) is not allowed")
	}

	// Resolve to absolute path to check final destination
	absPath, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	// Get current working directory
	cwd, err := filepath.Abs(".")
	if err != nil {
		return err
	}

	// Ensure the absolute path is within the current working directory
	if !strings.HasPrefix(absPath, cwd) {
		return fmt.Errorf("output path must be within the current working directory")
	}

	// Check file extension - only allow safe extensions
	allowedExtensions := []string{".md", ".txt", ".llm.txt"}
	ext := filepath.Ext(path)
	isAllowed := false
	for _, allowed := range allowedExtensions {
		if ext == allowed {
			isAllowed = true
			break
		}
	}
	if !isAllowed {
		return fmt.Errorf("only .md, .txt, and .llm.txt file extensions are allowed")
	}

	return nil
}