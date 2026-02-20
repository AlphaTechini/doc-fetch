#!/usr/bin/env node
const { spawn } = require('child_process');
const path = require('path');
const os = require('os');

// Determine binary path based on platform
const binDir = path.join(__dirname, '..');
const binaryName = os.platform() === 'win32' ? 'doc-fetch.exe' : 'doc-fetch';
const binaryPath = path.join(binDir, binaryName);

// Check if binary exists
if (!require('fs').existsSync(binaryPath)) {
  console.error('❌ doc-fetch binary not found!');
  console.error('💡 Please run: npm install doc-fetch');
  process.exit(1);
}

const args = process.argv.slice(2);

// Spawn the Go binary
const child = spawn(binaryPath, args, {
  stdio: 'inherit'
});

child.on('error', (err) => {
  if (err.code === 'ENOENT') {
    console.error('❌ doc-fetch binary not found!');
    console.error('💡 Please run: npm install doc-fetch');
  } else {
    console.error('❌ Failed to start doc-fetch:', err.message);
  }
  process.exit(1);
});

child.on('exit', (code) => {
  process.exit(code || 0);
});