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
console.log(`📦 Expected binary: ${expectedBinary}`);
console.log(`📦 Will copy to: ${binaryName}\n`);

// List all available binaries for debugging
console.log('📋 Available binaries in package:');
let hasPlatformBinary = false;
let foundBinary = null;
try {
  const files = fs.readdirSync(packageDir);
  const binaries = files.filter(f => f.includes('doc-fetch') && !f.endsWith('.js'));
  if (binaries.length === 0) {
    console.log('   ⚠️  No binaries found! Package may be corrupted.');
  } else {
    binaries.forEach(file => {
      const stats = fs.statSync(path.join(packageDir, file));
      const size = (stats.size / 1024 / 1024).toFixed(2);
      
      // Check if this is the correct binary for current platform
      let isCorrect = false;
      if (platform === 'win32' && file.endsWith('.exe')) {
        isCorrect = true;
        foundBinary = file;
      } else if (platform !== 'win32' && !file.endsWith('.exe') && 
                 (file === 'doc-fetch' || file.includes(`_${platform}_`))) {
        isCorrect = true;
        foundBinary = file;
      }
      
      const marker = isCorrect ? '✅' : 'ℹ️ ';
      console.log(`   ${marker} ${file} (${size} MB)`);
      
      if (isCorrect && file === expectedBinary) {
        hasPlatformBinary = true;
      }
    });
  }
} catch (e) {
  console.log(`   ❌ Could not list directory: ${e.message}`);
}
console.log('');

// Smart detection: If we found a binary that works, use it!
if (!hasPlatformBinary && foundBinary && fs.existsSync(path.join(packageDir, foundBinary))) {
  console.log(`✅ Found compatible binary: ${foundBinary}`);
  console.log(`📋 Copying: ${foundBinary} → ${binaryName}`);
  
  try {
    fs.copyFileSync(path.join(packageDir, foundBinary), path.join(packageDir, binaryName));
    
    if (platform !== 'win32') {
      fs.chmodSync(path.join(packageDir, binaryName), 0o755);
    }
    
    console.log(`✅ Binary installed successfully!\n`);
    
    // Verify it works
    try {
      const testPath = path.join(packageDir, binaryName);
      if (fs.existsSync(testPath)) {
        console.log('✨ Installation complete! You can now use: doc-fetch\n');
        return; // Success!
      }
    } catch (e) {
      // Continue to error handling
    }
  } catch (copyError) {
    console.error(`⚠️  Copy failed: ${copyError.message}`);
    // Continue to error handling
  }
}

// If we get here, something really went wrong
if (!fs.existsSync(sourcePath) && !hasPlatformBinary && !foundBinary) {
  console.error('⚠️  CRITICAL: No compatible binary found!');
  console.error(`   Searched for: ${expectedBinary}`);
  console.error(`   Found instead: ${foundBinary || 'nothing compatible'}`);
  console.error('');
  console.error('💡 This is a packaging error - the NPM package was published without your platform binary.');
  console.error('');
  console.error('💡 Immediate workaround - Install from source:');
  console.error('   1. Install Go: https://golang.org/dl/');
  console.error('   2. Run these commands:');
  console.error('      npm uninstall -g doc-fetch-cli');
  console.error('      git clone https://github.com/AlphaTechini/doc-fetch.git');
  console.error('      cd doc-fetch');
  if (platform === 'win32') {
    console.error('      go build -o doc-fetch.exe ./cmd/docfetch');
    console.error('      copy doc-fetch.exe "' + packageDir + '"');
  } else {
    console.error('      go build -o doc-fetch ./cmd/docfetch');
    console.error('      cp doc-fetch "' + packageDir + '"');
    console.error('      chmod +x "' + path.join(packageDir, 'doc-fetch') + '"');
  }
  console.error('');
  console.error('💡 Or wait for fixed version (check for updates):');
  console.error('   npm install -g doc-fetch-cli@latest');
  process.exit(1);
}

console.log(`✅ Found source binary: ${expectedBinary}`);

// Copy the binary to the expected location
try {
  console.log(`📋 Copying: ${expectedBinary} → ${binaryName}`);
  fs.copyFileSync(sourcePath, destPath);
  console.log(`✅ Copy successful`);
  
  // Verify the destination exists
  if (!fs.existsSync(destPath)) {
    throw new Error('Destination file does not exist after copy');
  }
  
  // Make executable on Unix-like systems
  if (platform !== 'win32') {
    fs.chmodSync(destPath, 0o755);
    console.log(`✅ Set executable permissions`);
  } else {
    console.log(`ℹ️  Windows: No chmod needed`);
  }
  
  console.log(`\n✅ Binary installed: ${binaryName}`);
} catch (error) {
  console.error(`\n❌ CRITICAL: Failed to install binary!`);
  console.error(`   Source: ${sourcePath}`);
  console.error(`   Destination: ${destPath}`);
  console.error(`   Error: ${error.message}`);
  console.error('');
  console.error('💡 Manual fix:');
  console.error(`   1. Navigate to: ${packageDir}`);
  console.error(`   2. Rename: ${expectedBinary} → ${binaryName}`);
  if (platform !== 'win32') {
    console.error(`   3. Run: chmod +x ${binaryName}`);
  }
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
