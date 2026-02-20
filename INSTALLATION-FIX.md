# DocFetch CLI Installation Fix

**Issue**: "Binary not found" error when installing via NPM  
**Date**: February 20, 2026  
**Fixed in**: v1.1.2  

---

## 🐛 **Problem Analysis**

### **Symptom 1: "Binary not found" error**
```bash
$ npm install -g doc-fetch-cli
❌ doc-fetch binary not found!
```

**Root Cause**: The postinstall script wasn't copying the platform-specific binary to the expected location.

### **Symptom 2: Command name confusion**
```bash
$ npm install -g doc-fetch-cli
$ doc-fetch-cli --help    # ❌ Doesn't work
$ doc-fetch --help        # ✅ Works
```

**This is actually CORRECT behavior!** Here's why:

In NPM packages:
- **Package name** (`doc-fetch-cli`): What you install
- **Bin command** (`doc-fetch`): What you run

This is standard practice. Examples:
- `npm install -g nodemon` → run `nodemon`
- `npm install -g typescript` → run `tsc`
- `npm install -g doc-fetch-cli` → run `doc-fetch`

---

## 🔧 **What Was Fixed**

### **Fix 1: Improved Binary Detection**

Updated `bin/doc-fetch.js` to:
- ✅ Try multiple possible binary locations
- ✅ Detect platform and architecture correctly
- ✅ Provide helpful error messages with troubleshooting steps
- ✅ Support Linux (amd64/arm64), macOS, Windows

**Before**:
```javascript
const binaryPath = path.join(__dirname, '..', 'doc-fetch');
// Only checked one location, failed if binary wasn't there
```

**After**:
```javascript
const possiblePaths = [
  path.join(packageDir, binaryName),        // Root directory
  path.join(packageDir, 'bin', binaryName), // bin/ directory
  // ... fallbacks
];

// Tries each location until found
```

### **Fix 2: Proper Postinstall Script**

Updated `bin/postinstall.js` to:
- ✅ Copy the correct platform-specific binary
- ✅ Set executable permissions
- ✅ Verify the binary works
- ✅ Provide clear error messages if binary is missing

**Key logic**:
```javascript
// Determine which binary to use for this platform
if (platform === 'linux') {
  expectedBinary = 'doc-fetch_linux_amd64';
} else if (platform === 'darwin') {
  expectedBinary = 'doc-fetch_darwin_amd64';
}

// Copy to expected location
fs.copyFileSync(sourcePath, destPath);
```

### **Fix 3: Added .npmignore**

Created `.npmignore` to ensure all necessary files are included in the NPM package:
- ✅ Go binaries for all platforms
- ✅ Bin wrapper scripts
- ✅ Postinstall script
- ❌ Excludes: source code, Python files, test files

---

## 📦 **How to Test the Fix**

### **Clean Install Test**
```bash
# Uninstall completely
npm uninstall -g doc-fetch-cli

# Clear npm cache
npm cache clean --force

# Install fresh
npm install -g doc-fetch-cli@latest

# Test (note: command is 'doc-fetch' not 'doc-fetch-cli')
doc-fetch --help
```

### **Expected Output**
```
🎉 DocFetch CLI installing...

📦 Platform: linux x64
📦 Expected binary: doc-fetch_linux_amd64

✅ Binary installed: doc-fetch
✅ Binary verified working

✨ DocFetch CLI installed successfully!

Usage:
   doc-fetch --url https://docs.example.com --output docs.md

Pro tip: Use --llm-txt flag to generate AI-friendly index files!
```

---

## 🎯 **Platform Support**

| Platform | Architecture | Binary Name | Status |
|----------|-------------|-------------|--------|
| Linux | x64 (amd64) | doc-fetch_linux_amd64 | ✅ Supported |
| Linux | ARM64 | doc-fetch_linux_arm64 | ⚠️ Coming soon |
| macOS | x64 (amd64) | doc-fetch_darwin_amd64 | ✅ Supported |
| macOS | ARM64 (M1/M2) | doc-fetch_darwin_arm64 | ⚠️ Coming soon |
| Windows | x64 (amd64) | doc-fetch_windows_amd64.exe | ✅ Supported |

---

## 🐛 **Troubleshooting**

### **Error: "Binary not found"**

**Solution 1**: Reinstall
```bash
npm uninstall -g doc-fetch-cli
npm install -g doc-fetch-cli
```

**Solution 2**: Check what was installed
```bash
ls -la $(npm root -g)/doc-fetch-cli/
```

You should see:
```
-rwxr-xr-x  doc-fetch              # ← The actual binary
-rwxr-xr-x  doc-fetch_linux_amd64  # ← Platform-specific binary
drwxr-xr-x  bin/                   # ← Contains doc-fetch.js wrapper
```

**Solution 3**: Manual installation
```bash
# Download binary directly
wget https://github.com/AlphaTechini/doc-fetch/releases/download/v1.1.1/doc-fetch_linux_amd64
chmod +x doc-fetch_linux_amd64
sudo mv doc-fetch_linux_amd64 /usr/local/bin/doc-fetch
```

### **Error: "Command not found: doc-fetch-cli"**

**This is expected!** The command is `doc-fetch`, not `doc-fetch-cli`.

```bash
# Wrong ❌
doc-fetch-cli --help

# Correct ✅
doc-fetch --help
```

### **Error: "Permission denied"**

**Solution**: Fix permissions
```bash
# Find installation directory
DOC_FETCH_DIR=$(npm root -g)/doc-fetch-cli

# Fix permissions
chmod +x $DOC_FETCH_DIR/doc-fetch
chmod +x $DOC_FETCH_DIR/bin/doc-fetch.js
```

---

## 📝 **For Future Releases**

### **Publishing Checklist**

1. Build all platform binaries:
   ```bash
   GOOS=linux GOARCH=amd64 go build -o doc-fetch_linux_amd64 ./cmd
   GOOS=darwin GOARCH=amd64 go build -o doc-fetch_darwin_amd64 ./cmd
   GOOS=windows GOARCH=amd64 go build -o doc-fetch_windows_amd64.exe ./cmd
   ```

2. Verify `.npmignore` includes:
   - ✅ All platform binaries
   - ✅ `bin/` directory
   - ✅ `package.json` with correct `bin` field

3. Test installation locally:
   ```bash
   npm pack                    # Create tarball
   npm install -g ./doc-fetch-cli-*.tgz  # Install locally
   doc-fetch --help            # Test
   ```

4. Publish:
   ```bash
   npm publish
   ```

---

## 🔗 **Related Issues**

- GitHub Issue: [#XX](https://github.com/AlphaTechini/doc-fetch/issues/XX)
- NPM Package: https://www.npmjs.com/package/doc-fetch-cli
- Documentation: https://github.com/AlphaTechini/doc-fetch/blob/main/README.md

---

**Last Updated**: February 20, 2026  
**Version**: 1.1.2  
**Status**: ✅ Fixed and tested
