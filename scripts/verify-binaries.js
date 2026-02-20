#!/usr/bin/env node
/**
 * Pre-publish verification script
 * Ensures all platform binaries are present before publishing to NPM
 */

const fs = require('fs');
const path = require('path');
const os = require('os');

const packageDir = path.join(__dirname, '..');

console.log('🔍 Verifying platform binaries before publish...\n');

const requiredBinaries = [
  { name: 'doc-fetch_linux_amd64', platform: 'linux', arch: 'x64' },
  { name: 'doc-fetch_linux_arm64', platform: 'linux', arch: 'arm64', optional: true },
  { name: 'doc-fetch_darwin_amd64', platform: 'darwin', arch: 'x64' },
  { name: 'doc-fetch_darwin_arm64', platform: 'darwin', arch: 'arm64', optional: true },
  { name: 'doc-fetch_windows_amd64.exe', platform: 'win32', arch: 'x64' },
];

let allPresent = true;
const missing = [];
const present = [];

console.log('📦 Checking binaries:\n');

requiredBinaries.forEach(binary => {
  const binaryPath = path.join(packageDir, binary.name);
  const exists = fs.existsSync(binaryPath);
  
  if (exists) {
    const stats = fs.statSync(binaryPath);
    const size = (stats.size / 1024 / 1024).toFixed(2);
    console.log(`   ✅ ${binary.name.padEnd(35)} (${size} MB) - ${binary.platform} ${binary.arch}`);
    present.push(binary.name);
  } else {
    const marker = binary.optional ? '⚠️ ' : '❌';
    const note = binary.optional ? '(optional)' : '(REQUIRED)';
    console.log(`   ${marker} ${binary.name.padEnd(35)} MISSING ${note}`);
    if (!binary.optional) {
      missing.push(binary.name);
      allPresent = false;
    }
  }
});

console.log('\n📊 Summary:');
console.log(`   Present:  ${present.length}`);
console.log(`   Missing:  ${missing.length}${missing.length > 0 ? ' (' + missing.join(', ') + ')' : ''}`);
console.log('');

if (!allPresent) {
  console.error('❌ CRITICAL: Required binaries are missing!');
  console.error('');
  console.error('💡 Build all platforms before publishing:');
  console.error('');
  console.error('   # Linux x64');
  console.error('   GOOS=linux GOARCH=amd64 go build -o doc-fetch_linux_amd64 ./cmd/docfetch');
  console.error('');
  console.error('   # Linux ARM64 (optional)');
  console.error('   GOOS=linux GOARCH=arm64 go build -o doc-fetch_linux_arm64 ./cmd/docfetch');
  console.error('');
  console.error('   # macOS Intel');
  console.error('   GOOS=darwin GOARCH=amd64 go build -o doc-fetch_darwin_amd64 ./cmd/docfetch');
  console.error('');
  console.error('   # macOS Apple Silicon (optional)');
  console.error('   GOOS=darwin GOARCH=arm64 go build -o doc-fetch_darwin_arm64 ./cmd/docfetch');
  console.error('');
  console.error('   # Windows x64');
  console.error('   GOOS=windows GOARCH=amd64 go build -o doc-fetch_windows_amd64.exe ./cmd/docfetch');
  console.error('');
  console.error('Then run this script again to verify.');
  console.error('');
  process.exit(1);
}

// Verify minimum file sizes (catches corrupted/empty files)
console.log('🔍 Verifying binary integrity...\n');

let allValid = true;
present.forEach(binaryName => {
  const binaryPath = path.join(packageDir, binaryName);
  const stats = fs.statSync(binaryPath);
  
  // Minimum expected size: 5MB (actual binaries are ~8-9MB)
  const minSize = 5 * 1024 * 1024;
  
  if (stats.size < minSize) {
    const sizeMB = (stats.size / 1024 / 1024).toFixed(2);
    console.log(`   ⚠️  ${binaryName.padEnd(35)} Suspiciously small (${sizeMB} MB)`);
    allValid = false;
  } else {
    const sizeMB = (stats.size / 1024 / 1024).toFixed(2);
    console.log(`   ✅ ${binaryName.padEnd(35)} OK (${sizeMB} MB)`);
  }
});

console.log('');

if (!allValid) {
  console.error('⚠️  Warning: Some binaries seem too small and may be corrupted.');
  console.error('   Consider rebuilding them before publishing.\n');
}

// Final verdict
if (allPresent && allValid) {
  console.log('✅ All required binaries present and valid!');
  console.log('✅ Ready to publish to NPM\n');
  process.exit(0);
} else if (allPresent) {
  console.log('⚠️  All binaries present but some may be corrupted.');
  console.log('⚠️  Proceed with caution!\n');
  process.exit(0); // Allow publishing but warn
} else {
  process.exit(1);
}
