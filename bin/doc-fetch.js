#!/usr/bin/env node
/**
 * DocFetch CLI Wrapper
 * Spawns the correct platform-specific Go binary
 */

const { spawn } = require('child_process');
const path = require('path');
const fs = require('fs');

const packageDir = path.join(__dirname, '..');
const platform = process.platform;

// Platform-aware binary search - only try binaries that match current platform
function getPlatformSuffix() {
  if (platform === 'win32') {
    return process.arch === 'arm64' ? '_windows_arm64.exe' : '_windows_amd64.exe';
  }
  if (platform === 'darwin') {
    return process.arch === 'arm64' ? '_darwin_arm64' : '_darwin_amd64';
  }
  return process.arch === 'arm64' ? '_linux_arm64' : '_linux_amd64';
}

const platformSuffix = getPlatformSuffix();
const specificName = 'doc-fetch' + platformSuffix;

// Search for binaries in platform-specific order
const possibleBinaries = [];
possibleBinaries.push(specificName, platform === 'win32' ? 'doc-fetch.exe' : 'doc-fetch');

let binaryPath = null;
for (const name of possibleBinaries) {
  const testPath = path.join(packageDir, name);
  if (fs.existsSync(testPath)) {
    binaryPath = testPath;
    break;
  }
}

if (!binaryPath) {
  console.error('❌ doc-fetch binary not found!');
  console.error('');
  console.error('💡 Platform:', platform, '(' + process.arch + ')');
  console.error('💡 Searched in:', packageDir);
  console.error('');
  console.error('💡 Try reinstalling:');
  console.error('   npm uninstall -g doc-fetch-cli');
  console.error('   npm install -g doc-fetch-cli@latest');
  process.exit(1);
}

const args = process.argv.slice(2);

// Spawn the Go binary
const child = spawn(binaryPath, args, {
  stdio: 'inherit'
});

child.on('error', (err) => {
  if (err.code === 'ENOENT') {
    console.error('❌ Failed to execute binary');
    console.error('   Path:', binaryPath);
    console.error('   Error: File not found or no execute permission');
    if (platform !== 'win32') {
      console.error('');
      console.error('💡 Fix permissions: chmod +x "' + binaryPath + '"');
    }
  } else {
    console.error('❌ Failed to start:', err.message);
  }
  process.exit(1);
});

child.on('exit', (code) => {
  process.exit(code || 0);
});
