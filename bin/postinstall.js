#!/usr/bin/env node
/**
 * Post-install script for doc-fetch-cli
 * Detects and uses the correct platform binary (no copying needed)
 */

const path = require('path');
const os = require('os');
const fs = require('fs');

console.log('🎉 DocFetch CLI installing...\n');

const packageDir = path.join(__dirname, '..');
const platform = os.platform();
const arch = os.arch();

// Define binary names to search for (in priority order)
let searchNames = [];
if (platform === 'win32') {
  searchNames = ['doc-fetch_windows_amd64.exe', 'doc-fetch.exe'];
} else if (platform === 'darwin') {
  searchNames = arch === 'arm64' 
    ? ['doc-fetch_darwin_arm64', 'doc-fetch_darwin_amd64', 'doc-fetch']
    : ['doc-fetch_darwin_amd64', 'doc-fetch_darwin_arm64', 'doc-fetch'];
} else {
  // Linux
  searchNames = arch === 'arm64'
    ? ['doc-fetch_linux_arm64', 'doc-fetch_linux_amd64', 'doc-fetch']
    : ['doc-fetch_linux_amd64', 'doc-fetch_linux_arm64', 'doc-fetch'];
}

console.log(`📦 Platform: ${platform} ${arch}`);
console.log(`🔍 Searching for binary...\n`);

// List all doc-fetch binaries in package
console.log('📋 Available binaries:');
let foundBinary = null;
try {
  const files = fs.readdirSync(packageDir);
  const binaries = files.filter(f => f.includes('doc-fetch') && !f.endsWith('.js'));
  
  if (binaries.length === 0) {
    console.log('   ❌ No binaries found! Package is corrupted.\n');
    process.exit(1);
  }
  
  binaries.forEach(file => {
    const stats = fs.statSync(path.join(packageDir, file));
    const size = (stats.size / 1024 / 1024).toFixed(2);
    
    // Check if this is a valid binary for current platform
    const isValid = searchNames.includes(file);
    const marker = isValid ? '✅' : 'ℹ️ ';
    console.log(`   ${marker} ${file} (${size} MB)`);
    
    // Use first valid match
    if (isValid && !foundBinary) {
      foundBinary = file;
    }
  });
} catch (e) {
  console.log(`   ❌ Error listing directory: ${e.message}\n`);
  process.exit(1);
}

console.log('');

if (!foundBinary) {
  console.error('❌ CRITICAL: No compatible binary found for your platform!');
  console.error(`   Searched for: ${searchNames.join(', ')}`);
  console.error('');
  console.error('💡 This is a packaging error. Please:');
  console.error('   1. Report issue: https://github.com/AlphaTechini/doc-fetch/issues');
  console.error('   2. Include your platform: ' + platform + ' ' + arch);
  console.error('   3. Or install from source (see README)\n');
  process.exit(1);
}

// Update wrapper script to use found binary
const wrapperPath = path.join(packageDir, 'bin', 'doc-fetch.js');
if (fs.existsSync(wrapperPath)) {
  try {
    let wrapper = fs.readFileSync(wrapperPath, 'utf8');
    
    // Update the binary name in wrapper
    wrapper = wrapper.replace(
      /const binaryName = ['"][^'"]+['"];/,
      `const binaryName = '${foundBinary}';`
    );
    
    fs.writeFileSync(wrapperPath, wrapper);
    console.log(`✅ Configured wrapper to use: ${foundBinary}`);
  } catch (e) {
    console.log(`⚠️  Could not update wrapper: ${e.message}`);
    console.log(`   You may need to run: doc-fetch (instead of doc-fetch-cli)`);
  }
}

console.log('');
console.log('✨ Installation complete!');
console.log('');
console.log('Usage:');
console.log('   doc-fetch --url https://docs.example.com --output docs.md');
console.log('');
console.log('Pro tip: Use --llm-txt flag to generate AI-friendly index files!\n');
