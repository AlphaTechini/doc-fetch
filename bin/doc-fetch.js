#!/usr/bin/env node
const { spawn } = require('child_process');
const path = require('path');
const os = require('os');
const fs = require('fs');

// Get the package installation directory
const packageDir = path.join(__dirname, '..');

// Determine binary name based on platform
const platform = os.platform();
const arch = os.arch();
let binaryName;

if (platform === 'win32') {
  binaryName = 'doc-fetch.exe';
} else if (platform === 'darwin') {
  binaryName = 'doc-fetch_darwin_amd64';
} else {
  // Linux and others
  binaryName = arch === 'arm64' ? 'doc-fetch_linux_arm64' : 'doc-fetch_linux_amd64';
}

// Try multiple possible locations
const possiblePaths = [
  path.join(packageDir, binaryName),                    // Root directory
  path.join(packageDir, 'bin', binaryName),             // bin/ directory  
  path.join(packageDir, binaryName.replace('_linux_amd64', '')), // Fallback to generic name
];

// Find the binary
let binaryPath = null;
for (const testPath of possiblePaths) {
  if (fs.existsSync(testPath)) {
    binaryPath = testPath;
    break;
  }
}

if (!binaryPath) {
  console.error('❌ doc-fetch binary not found!');
  console.error('');
  console.error('💡 Troubleshooting steps:');
  console.error('   1. Reinstall: npm uninstall -g doc-fetch-cli && npm install -g doc-fetch-cli');
  console.error('   2. Check installation: ls -la $(npm root -g)/doc-fetch-cli/');
  console.error('   3. Report issue: https://github.com/AlphaTechini/doc-fetch/issues');
  console.error('');
  console.error(`   Expected binary: ${possiblePaths[0]}`);
  console.error(`   Platform: ${platform} ${arch}`);
  process.exit(1);
}

const args = process.argv.slice(2);

// Spawn the Go binary
const child = spawn(binaryPath, args, {
  stdio: 'inherit'
});

child.on('error', (err) => {
  if (err.code === 'ENOENT') {
    console.error('❌ Failed to execute doc-fetch binary');
    console.error(`   Binary path: ${binaryPath}`);
    console.error('   Error: Binary file may be corrupted or missing execute permissions');
    console.error('');
    console.error('💡 Try reinstalling: npm uninstall -g doc-fetch-cli && npm install -g doc-fetch-cli');
  } else {
    console.error('❌ Failed to start doc-fetch:', err.message);
  }
  process.exit(1);
});

child.on('exit', (code) => {
  process.exit(code || 0);
});