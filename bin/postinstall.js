#!/usr/bin/env node
/**
 * Post-install script for doc-fetch-cli
 * Downloads the latest Go binary from GitHub Releases, with bundled fallback
 */

const path = require('path');
const os = require('os');
const fs = require('fs');
const https = require('https');

console.log('📦 DocFetch CLI setting up...\n');

const packageDir = path.join(__dirname, '..');
const platform = os.platform();
const arch = os.arch();

// Read package.json for version
let version = '2.5.8';
try {
  const pkg = JSON.parse(fs.readFileSync(path.join(packageDir, 'package.json'), 'utf8'));
  version = pkg.version;
} catch (e) {
  // use default
}

// Map platform/arch to binary name
function getBinaryName() {
  const goos = platform === 'win32' ? 'windows' : (platform === 'darwin' ? 'darwin' : 'linux');
  const goarch = arch === 'arm64' ? 'arm64' : 'amd64';
  const name = `doc-fetch_${goos}_${goarch}`;
  return { name, fullName: goos === 'windows' ? name + '.exe' : name };
}

const { name, fullName } = getBinaryName();
const downloadUrl = `https://github.com/AlphaTechini/doc-fetch/releases/download/v${version}/${fullName}`;
const destPath = path.join(packageDir, fullName);

console.log(`📥 Platform: ${platform} ${arch}`);
console.log(`📥 Downloading binary for v${version}...`);
console.log(`   ${downloadUrl}\n`);

function downloadBinary(url, dest) {
  return new Promise((resolve, reject) => {
    const tempDest = dest + '.download';
    const file = fs.createWriteStream(tempDest, { mode: 0o755 });
    https.get(url, { timeout: 30000 }, (res) => {
      if (res.statusCode === 302 || res.statusCode === 301) {
        file.close();
        fs.unlinkSync(tempDest);
        downloadBinary(res.headers.location, dest).then(resolve).catch(reject);
        return;
      }
      if (res.statusCode !== 200) {
        file.close();
        fs.unlinkSync(tempDest);
        reject(new Error(`HTTP ${res.statusCode}`));
        return;
      }
      res.pipe(file);
      file.on('finish', () => {
        file.close();
        fs.renameSync(tempDest, dest);
        resolve();
      });
    }).on('error', (err) => {
      file.close();
      try { fs.unlinkSync(tempDest); } catch (e) {}
      reject(err);
    }).on('timeout', () => {
      file.close();
      try { fs.unlinkSync(tempDest); } catch (e) {}
      reject(new Error('Download timeout'));
    });
  });
}

async function setup() {
  // Try downloading the latest binary from GitHub Releases
  try {
    await downloadBinary(downloadUrl, destPath);
    console.log('✅ Binary downloaded from GitHub Releases\n');
    updateWrapper(fullName);
    return;
  } catch (e) {
    console.log(`⚠️  Download failed (${e.message}), falling back to bundled binary...\n`);
  }

  // Fallback: find bundled binary
  let searchOrder = [];
  if (platform === 'win32') {
    searchOrder = [fullName, 'doc-fetch.exe'];
  } else if (platform === 'darwin') {
    searchOrder = arch === 'arm64'
      ? ['doc-fetch_darwin_arm64', 'doc-fetch_darwin_amd64', 'doc-fetch']
      : ['doc-fetch_darwin_amd64', 'doc-fetch_darwin_arm64', 'doc-fetch'];
  } else {
    searchOrder = arch === 'arm64'
      ? ['doc-fetch_linux_arm64', 'doc-fetch_linux_amd64', 'doc-fetch']
      : ['doc-fetch_linux_amd64', 'doc-fetch_linux_arm64', 'doc-fetch'];
  }

  let foundBinary = null;
  for (const name of searchOrder) {
    const testPath = path.join(packageDir, name);
    if (fs.existsSync(testPath)) {
      foundBinary = name;
      break;
    }
  }

  if (!foundBinary) {
    console.error('❌ No binary found. Please install from source or download manually.');
    console.error('   https://github.com/AlphaTechini/doc-fetch/releases\n');
    process.exit(1);
  }

  console.log(`✅ Using bundled binary: ${foundBinary}\n`);
  updateWrapper(foundBinary);
}

function updateWrapper(binaryName) {
  const wrapperPath = path.join(packageDir, 'bin', 'doc-fetch.js');
  if (fs.existsSync(wrapperPath)) {
    try {
      let wrapper = fs.readFileSync(wrapperPath, 'utf8');
      wrapper = wrapper.replace(
        /const binaryName = ['"][^'"]+['"];/,
        "const binaryName = '" + binaryName + "';"
      );
      fs.writeFileSync(wrapperPath, wrapper);
    } catch (e) {
      console.log(`⚠️  Could not update wrapper: ${e.message}\n`);
    }
  }
}

setup().catch((err) => {
  console.error(`❌ Setup failed: ${err.message}\n`);
  process.exit(1);
});
