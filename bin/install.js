#!/usr/bin/env node
const os = require('os');
const path = require('path');
const { execSync } = require('child_process');
const fs = require('fs');
const https = require('https');
const { pipeline } = require('stream');
const { promisify } = require('util');

const finished = promisify(pipeline);

async function getBinaryUrl() {
  const platform = os.platform();
  const arch = os.arch();
  
  // Map to Go build targets
  let goos, goarch;
  switch(platform) {
    case 'win32': goos = 'windows'; break;
    case 'darwin': goos = 'darwin'; break;
    default: goos = 'linux';
  }
  
  switch(arch) {
    case 'x64': goarch = 'amd64'; break;
    case 'arm64': goarch = 'arm64'; break;
    default: goarch = 'amd64';
  }
  
  return {
    url: `https://github.com/AlphaTechini/doc-fetch/releases/download/v1.0.0/doc-fetch_${goos}_${goarch}`,
    filename: platform === 'win32' ? 'doc-fetch.exe' : 'doc-fetch'
  };
}

async function downloadBinary() {
  const binDir = path.join(__dirname, '..');
  if (!fs.existsSync(binDir)) {
    fs.mkdirSync(binDir, { recursive: true });
  }
  
  const { url, filename } = await getBinaryUrl();
  const binaryPath = path.join(binDir, filename);
  
  console.log('📥 Downloading doc-fetch binary...');
  console.log(`   Platform: ${os.platform()} ${os.arch()}`);
  console.log(`   URL: ${url}`);
  
  try {
    const response = await new Promise((resolve, reject) => {
      https.get(url, (res) => {
        if (res.statusCode === 404) {
          reject(new Error('Binary not found for your platform'));
        } else if (res.statusCode !== 200) {
          reject(new Error(`HTTP ${res.statusCode}: ${res.statusMessage}`));
        } else {
          resolve(res);
        }
      }).on('error', reject);
    });
    
    await finished(response, fs.createWriteStream(binaryPath));
    
    // Make executable on Unix-like systems
    if (os.platform() !== 'win32') {
      fs.chmodSync(binaryPath, 0o755);
    }
    
    console.log('✅ Binary downloaded successfully!');
    return true;
  } catch (error) {
    console.error('❌ Failed to download binary:', error.message);
    console.log('🔄 Falling back to Go build from source...');
    
    // Fallback: build from source if Go is available
    try {
      execSync('go version', { stdio: 'pipe' });
      console.log('🏗️  Building from source...');
      execSync('go build -o bin/doc-fetch ./cmd/docfetch', { 
        stdio: 'inherit',
        cwd: path.join(__dirname, '..')
      });
      console.log('✅ Built successfully from source!');
      return true;
    } catch (buildError) {
      console.error('❌ Go not found or build failed.');
      console.error('💡 Please install Go (https://golang.org/dl/) or download the binary manually.');
      return false;
    }
  }
}

// Run the installation
(async () => {
  try {
    const success = await downloadBinary();
    if (!success) {
      process.exit(1);
    }
  } catch (error) {
    console.error('Installation failed:', error.message);
    process.exit(1);
  }
})();