# Quick Fix: Windows Binary Not Found

**Symptom**: After installing, you get "binary not found" error

**Root Cause**: The postinstall script didn't rename the binary correctly

---

## 🔧 **Manual Fix (30 seconds)**

### **Step 1: Find Installation Directory**
```powershell
npm root -g
```

This will show something like:
```
C:\Users\YourName\AppData\Roaming\npm\node_modules
```

### **Step 2: Navigate to doc-fetch-cli**
```powershell
cd "C:\Users\YourName\AppData\Roaming\npm\node_modules\doc-fetch-cli"
```

### **Step 3: List Files**
```powershell
dir doc-fetch*
```

You should see:
```
doc-fetch_windows_amd64.exe    ← This exists
doc-fetch.js                    ← Wrapper script
```

But MISSING:
```
doc-fetch.exe                   ← This is what's needed!
```

### **Step 4: Rename the Binary**
```powershell
copy doc-fetch_windows_amd64.exe doc-fetch.exe
```

Or using `ren`:
```powershell
ren doc-fetch_windows_amd64.exe doc-fetch.exe
```

### **Step 5: Verify**
```powershell
doc-fetch --help
```

It should work now! ✅

---

## 🎯 **Why This Happens**

The package includes platform-specific binaries:
- `doc-fetch_windows_amd64.exe` (Windows)
- `doc-fetch_darwin_amd64` (macOS)
- `doc-fetch_linux_amd64` (Linux)

But the wrapper expects just `doc-fetch.exe` on Windows.

The postinstall script SHOULD do this rename automatically, but sometimes it fails due to:
- Permissions issues
- Antivirus blocking file operations
- NPM running in restricted mode

---

## 🚀 **Permanent Fix**

Reinstall with v1.1.3 or later which has better logging:

```powershell
npm uninstall -g doc-fetch-cli
npm install -g doc-fetch-cli@latest
```

Watch the output - it should show:
```
📋 Copying: doc-fetch_windows_amd64.exe → doc-fetch.exe
✅ Copy successful
ℹ️  Windows: No chmod needed
✅ Binary installed: doc-fetch.exe
```

If you still see errors, the detailed logs will tell you exactly what went wrong.

---

## 📞 **Still Having Issues?**

Run this diagnostic:
```powershell
$dir = npm root -g + "\doc-fetch-cli"
Write-Host "Installation directory: $dir"
Write-Host "`nFiles containing 'doc-fetch':"
Get-ChildItem $dir | Where-Object { $_.Name -like "*doc-fetch*" } | Select-Object Name, Length
```

Share the output when reporting the issue at:
https://github.com/AlphaTechini/doc-fetch/issues

---

**Fixed in**: v1.1.3  
**Platform**: Windows only  
**Workaround**: Manual rename (above)
