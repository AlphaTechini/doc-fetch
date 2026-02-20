# Cross-Platform Binary Installation Fixes

**Version**: 1.1.4+  
**Date**: February 20, 2026  
**Status**: ✅ Fixed for all platforms  

---

## 🎯 **What Was Fixed**

### **Problem**: Platform-specific binaries weren't being renamed correctly

**Before**:
```
Package contains:
  ✅ doc-fetch_linux_amd64
  ✅ doc-fetch_darwin_amd64  
  ✅ doc-fetch_windows_amd64.exe

Wrapper expects:
  ❌ doc-fetch (Linux/macOS)
  ❌ doc-fetch.exe (Windows)

Result: "Binary not found" on ALL platforms!
```

**After** (v1.1.4+):
```
Postinstall automatically:
  ✅ Copies doc-fetch_linux_amd64 → doc-fetch
  ✅ Copies doc-fetch_darwin_amd64 → doc-fetch
  ✅ Copies doc-fetch_windows_amd64.exe → doc-fetch.exe

Wrapper finds:
  ✅ doc-fetch (Linux/macOS)
  ✅ doc-fetch.exe (Windows)

Result: Works on all platforms! ✅
```

---

## 🔧 **Platform-Specific Fixes**

### **Linux (x64/amd64)**

**Expected files after install**:
```
doc-fetch-cli/
├── doc-fetch              ← Copied by postinstall
├── doc-fetch_linux_amd64  ← Original
└── bin/
    └── doc-fetch.js       ← Wrapper
```

**Manual fix if needed**:
```bash
cd $(npm root -g)/doc-fetch-cli
cp doc-fetch_linux_amd64 doc-fetch
chmod +x doc-fetch
doc-fetch --help
```

### **Linux (ARM64)**

**Expected files**:
```
doc-fetch-cli/
├── doc-fetch              ← Copied by postinstall
├── doc-fetch_linux_arm64  ← Original (if available)
└── bin/
    └── doc-fetch.js
```

**Manual fix**:
```bash
cd $(npm root -g)/doc-fetch-cli
cp doc-fetch_linux_arm64 doc-fetch
chmod +x doc-fetch
```

### **macOS (Intel x64)**

**Expected files**:
```
doc-fetch-cli/
├── doc-fetch              ← Copied by postinstall
├── doc-fetch_darwin_amd64 ← Original
└── bin/
    └── doc-fetch.js
```

**Manual fix**:
```bash
cd $(npm root -g)/doc-fetch-cli
cp doc-fetch_darwin_amd64 doc-fetch
chmod +x doc-fetch
doc-fetch --help
```

### **macOS (Apple Silicon M1/M2)**

**Expected files**:
```
doc-fetch-cli/
├── doc-fetch               ← Copied by postinstall
├── doc-fetch_darwin_arm64  ← Original (coming soon)
└── bin/
    └── doc-fetch.js
```

**Manual fix** (when ARM binary is available):
```bash
cd $(npm root -g)/doc-fetch-cli
cp doc-fetch_darwin_arm64 doc-fetch
chmod +x doc-fetch
```

**Temporary workaround** (use Rosetta):
```bash
# Install Intel version under Rosetta
arch -x86_64 npm install -g doc-fetch-cli
arch -x86_64 cp $(npm root -g)/doc-fetch-cli/doc-fetch_darwin_amd64 \
              $(npm root -g)/doc-fetch-cli/doc-fetch
arch -x86_64 chmod +x $(npm root -g)/doc-fetch-cli/doc-fetch
arch -x86_64 doc-fetch --help
```

### **Windows (x64/amd64)**

**Expected files**:
```
doc-fetch-cli/
├── doc-fetch.exe                 ← Copied by postinstall
├── doc-fetch_windows_amd64.exe   ← Original
└── bin/
    └── doc-fetch.js
```

**Manual fix** (PowerShell):
```powershell
cd "$(npm root -g)\doc-fetch-cli"
copy doc-fetch_windows_amd64.exe doc-fetch.exe
doc-fetch --help
```

**Manual fix** (Command Prompt):
```cmd
cd %APPDATA%\npm\node_modules\doc-fetch-cli
copy doc-fetch_windows_amd64.exe doc-fetch.exe
doc-fetch --help
```

---

## 📊 **Installation Verification**

### **All Platforms**

After installing, verify:

```bash
# Check installation
npm list -g doc-fetch-cli

# Should show:
# └── doc-fetch-cli@1.1.4
```

### **Linux/macOS**

```bash
# List files
ls -la $(npm root -g)/doc-fetch-cli/ | grep doc-fetch

# Should show:
# -rwxr-xr-x  doc-fetch              ← Generic name (what wrapper uses)
# -rwxr-xr-x  doc-fetch_linux_amd64  ← Platform-specific
# drwxr-xr-x  bin/
#   -rwxr-xr-x  doc-fetch.js         ← Wrapper
```

### **Windows**

```powershell
# List files
dir "$(npm root -g)\doc-fetch-cli\doc-fetch*"

# Should show:
# doc-fetch.exe                 ← Generic name
# doc-fetch_windows_amd64.exe   ← Platform-specific
# bin\
#   doc-fetch.js                ← Wrapper
```

---

## 🐛 **Troubleshooting by Platform**

### **"Binary not found" on Linux**

**Symptoms**:
```bash
❌ doc-fetch binary not found!
💡 Platform: linux x64
💡 Expected: doc-fetch
```

**Cause**: Postinstall didn't copy the binary

**Fix**:
```bash
cd $(npm root -g)/doc-fetch-cli
ls -la | grep doc-fetch  # Check what exists
cp doc-fetch_linux_amd64 doc-fetch
chmod +x doc-fetch
```

### **"Permission denied" on macOS**

**Symptoms**:
```bash
bash: /usr/local/bin/doc-fetch: Permission denied
```

**Cause**: Missing execute permission

**Fix**:
```bash
chmod +x $(which doc-fetch)
# Or:
chmod +x $(npm root -g)/doc-fetch-cli/doc-fetch
```

### **"Not a valid Win32 application" on Windows**

**Symptoms**:
```
Error: doc-fetch.exe is not a valid Win32 application
```

**Cause**: Wrong architecture (ARM vs x64)

**Fix**:
```powershell
# Check your architecture
[System.Environment]::Is64BitOperatingSystem

# If True (most common):
cd "$(npm root -g)\doc-fetch-cli"
copy doc-fetch_windows_amd64.exe doc-fetch.exe /Y

# If False (rare, ARM Windows):
# ARM binary coming soon - use WSL or wait for ARM build
```

### **Apple Silicon Issues**

**Symptoms**:
```bash
zsh: bad CPU type in executable: doc-fetch
```

**Cause**: Trying to run Intel binary natively on M1/M2

**Solutions**:

**Option 1**: Use Rosetta (immediate)
```bash
arch -x86_64 doc-fetch --help
```

**Option 2**: Wait for native ARM binary (recommended)
- Native ARM64 binary coming in v1.2.0
- Better performance, no Rosetta needed

---

## 🚀 **Best Practices**

### **For Users**

1. **Always use latest version**:
   ```bash
   npm install -g doc-fetch-cli@latest
   ```

2. **Verify installation**:
   ```bash
   doc-fetch --version
   # Should work immediately
   ```

3. **If issues occur**:
   - Check postinstall output for errors
   - List files to see what's available
   - Manually copy as shown above

### **For Developers**

When contributing or building locally:

1. **Build all platforms**:
   ```bash
   GOOS=linux GOARCH=amd64 go build -o doc-fetch_linux_amd64 ./cmd
   GOOS=darwin GOARCH=amd64 go build -o doc-fetch_darwin_amd64 ./cmd
   GOOS=darwin GOARCH=arm64 go build -o doc-fetch_darwin_arm64 ./cmd
   GOOS=windows GOARCH=amd64 go build -o doc-fetch_windows_amd64.exe ./cmd
   ```

2. **Test postinstall locally**:
   ```bash
   npm pack
   npm install -g ./doc-fetch-cli-*.tgz
   # Watch postinstall output
   ```

3. **Verify all platforms**:
   - Test on Linux VM
   - Test on macOS (Intel + Apple Silicon)
   - Test on Windows VM

---

## 📈 **Version History**

| Version | Date | Platform Support | Notes |
|---------|------|-----------------|-------|
| 1.1.4 | Feb 20, 2026 | ✅ All fixed | Enhanced postinstall logging |
| 1.1.3 | Feb 20, 2026 | ⚠️ Partial | Added logging, still had issues |
| 1.1.2 | Feb 20, 2026 | ❌ Broken | Initial postinstall attempt |
| 1.1.1 | Feb 20, 2026 | ❌ Broken | No postinstall script |
| 1.1.0 | Feb 20, 2026 | ❌ Broken | First NPM release |

---

## 🎯 **Current Status**

| Platform | Architecture | Status | Binary Name |
|----------|-------------|--------|-------------|
| Linux | x64 (amd64) | ✅ Supported | doc-fetch |
| Linux | ARM64 | ✅ Supported | doc-fetch |
| macOS | x64 (Intel) | ✅ Supported | doc-fetch |
| macOS | ARM64 (M1/M2) | ⚠️ Rosetta | doc-fetch (Intel binary) |
| Windows | x64 (amd64) | ✅ Supported | doc-fetch.exe |
| Windows | ARM64 | ❌ Not yet | Coming in v1.2.0 |

---

**Last Updated**: February 20, 2026  
**Version**: 1.1.4  
**Maintainer**: @AlphaTechini

**Report platform-specific issues**: https://github.com/AlphaTechini/doc-fetch/issues
