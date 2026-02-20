#!/usr/bin/env node
const { spawn } = require('child_process');
const path = require('path');
const os = require('os');
const fs = require('fs');

// Get the package installation directory
const packageDir = path.join(__dirname, '..');

// Determine binary name based on platform (what postinstall creates)
const platform = os.platform();
const binaryName = platform === 'win32' ? 'doc-fetch.exe' : 'doc-fetch';

// Try multiple possible locations
const possiblePaths = [
  path.join(packageDir, binaryName),        // Root directory (postinstall copies here)
  path.join(packageDir, 'bin', binaryName), // bin/ directory (fallback)
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
  console.error(`💡 Platform: ${platform} (${os.arch()})`);
  console.error('💡 Expected: doc-fetch' + (platform === 'win32' ? '.exe' : ''));
  console.error('');
  console.error('💡 Troubleshooting steps:');
  console.error('   1. List files: ls -la $(npm root -g)/doc-fetch-cli/');
  console.error('   2. Look for platform-specific binary:');
  if (platform === 'win32') {
    console.error('      doc-fetch_windows_amd64.exe');
  } else if (platform === 'darwin') {
    console.error('      doc-fetch_darwin_amd64 (Intel) or doc-fetch_darwin_arm64 (M1/M2)');
  } else {
    console.error('      doc-fetch_linux_amd64 (x64) or doc-fetch_linux_arm64 (ARM)');
  }
  console.error('   3. Manually copy: cp <platform-binary> doc-fetch' + (platform === 'win32' ? '.exe' : ''));
  console.error('   4. Make executable: chmod +x doc-fetch (Linux/macOS only)');
  console.error('   5. Reinstall: npm uninstall -g doc-fetch-cli && npm install -g doc-fetch-cli@latest');
  console.error('');
  console.error('📦 Package directory:', packageDir);
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
    console.error('   Error: File not found or no execute permission');
    console.error('');
    if (platform !== 'win32') {
      console.error('💡 Fix permissions: chmod +x "' + binaryPath + '"');
    }
    console.error('💡 Or reinstall: npm uninstall -g doc-fetch-cli && npm install -g doc-fetch-cli@latest');
  } else {
    console.error('❌ Failed to start doc-fetch:', err.message);
  }
  process.exit(1);
});

child.on('exit', (code) => {
  process.exit(code || 0);
});
