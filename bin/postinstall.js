#!/usr/bin/env node
/**
 * Post-install script for doc-fetch-cli
 * Copies the correct platform-specific binary and sets up PATH
 */

const { execSync } = require('child_process');
const path = require('path');
const os = require('os');
const fs = require('fs');

console.log('🎉 DocFetch CLI installing...\n');

const packageDir = path.join(__dirname, '..');
const platform = os.platform();
const arch = os.arch();

// Determine which binary to use
let binaryName;
let expectedBinary;

if (platform === 'win32') {
  binaryName = 'doc-fetch.exe';
  expectedBinary = 'doc-fetch_windows_amd64.exe';
} else if (platform === 'darwin') {
  binaryName = 'doc-fetch';
  expectedBinary = 'doc-fetch_darwin_amd64';
} else {
  // Linux
  binaryName = 'doc-fetch';
  expectedBinary = arch === 'arm64' ? 'doc-fetch_linux_arm64' : 'doc-fetch_linux_amd64';
}

const sourcePath = path.join(packageDir, expectedBinary);
const destPath = path.join(packageDir, binaryName);

console.log(`📦 Platform: ${platform} ${arch}`);
console.log(`📦 Expected binary: ${expectedBinary}\n`);

// Check if the expected binary exists
if (!fs.existsSync(sourcePath)) {
  console.error(`⚠️  Warning: Expected binary not found: ${expectedBinary}`);
  console.error('');
  console.error('💡 This might be because:');
  console.error('   1. The package was published without binaries');
  console.error('   2. Your platform/architecture is not supported');
  console.error('');
  console.error('Supported platforms:');
  console.error('   - Linux x64 (amd64)');
  console.error('   - macOS x64 (amd64)');  
  console.error('   - Windows x64 (amd64)');
  console.error('');
  console.error('💡 Workaround: Install from source');
  console.error('   npm uninstall -g doc-fetch-cli');
  console.error('   git clone https://github.com/AlphaTechini/doc-fetch.git');
  console.error('   cd doc-fetch && go build -o doc-fetch ./cmd/docfetch');
  console.error('   sudo cp doc-fetch /usr/local/bin/');
  process.exit(1);
}

// Copy the binary to the expected location
try {
  fs.copyFileSync(sourcePath, destPath);
  
  // Make executable on Unix-like systems
  if (platform !== 'win32') {
    fs.chmodSync(destPath, 0o755);
  }
  
  console.log(`✅ Binary installed: ${binaryName}`);
} catch (error) {
  console.error(`❌ Failed to install binary: ${error.message}`);
  process.exit(1);
}

// Verify installation
try {
  const result = execSync(`"${destPath}" --help`, { encoding: 'utf8', stdio: ['pipe', 'pipe', 'ignore'] });
  console.log('✅ Binary verified working\n');
} catch (error) {
  console.error('⚠️  Warning: Could not verify binary execution');
  console.error(`   Error: ${error.message}\n`);
}

console.log('✨ DocFetch CLI installed successfully!\n');
console.log('Usage:');
console.log('   doc-fetch --url https://docs.example.com --output docs.md\n');
console.log('Pro tip: Use --llm-txt flag to generate AI-friendly index files!\n');
