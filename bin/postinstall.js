#!/usr/bin/env node
/**
 * Post-install script for doc-fetch-cli
 * Checks if global bin directory is in PATH and provides helpful instructions
 */

const { execSync } = require('child_process');
const path = require('path');
const os = require('os');
const fs = require('fs');

console.log('🎉 DocFetch CLI installed successfully!\n');

// Get npm global prefix
let globalPrefix;
try {
    globalPrefix = execSync('npm config get prefix', { encoding: 'utf8' }).trim();
} catch (error) {
    console.error('⚠️  Could not determine npm global prefix');
    globalPrefix = null;
}

if (globalPrefix) {
    const binDir = path.join(globalPrefix, 'bin');
    const isWindows = os.platform() === 'win32';
    
    console.log(`📦 Installed to: ${binDir}\n`);
    
    // Check if bin directory is in PATH
    const pathEnv = process.env.PATH || '';
    const pathDirs = pathEnv.split(isWindows ? ';' : ':');
    const isInPath = pathDirs.some(dir => path.resolve(dir) === path.resolve(binDir));
    
    if (!isInPath) {
        console.log('⚠️  WARNING: Global bin directory is not in your PATH!\n');
        console.log('To use doc-fetch-cli, add this directory to your PATH:\n');
        console.log(`   ${binDir}\n`);
        
        // Provide platform-specific instructions
        const shell = process.env.SHELL || '/bin/bash';
        const isZsh = shell.includes('zsh');
        const isBash = shell.includes('bash');
        
        console.log('Quick fix:\n');
        
        if (isWindows) {
            console.log('1. Open System Properties → Environment Variables');
            console.log('2. Edit PATH variable');
            console.log('3. Add this path:');
            console.log(`   ${binDir}`);
            console.log('4. Restart your terminal\n');
        } else if (isZsh) {
            console.log('Add this to your ~/.zshrc:');
            console.log(`   export PATH="${binDir}:$PATH"\n`);
            console.log('Then run: source ~/.zshrc\n');
        } else if (isBash) {
            console.log('Add this to your ~/.bashrc or ~/.bash_profile:');
            console.log(`   export PATH="${binDir}:$PATH"\n`);
            console.log('Then run: source ~/.bashrc\n');
        }
        
        console.log('Alternative: Use npx without installing globally\n');
        console.log('   npx doc-fetch-cli --url https://docs.example.com --output docs.md\n');
    } else {
        console.log('✅ Global bin directory is in your PATH\n');
        console.log('You can now use doc-fetch-cli!\n');
        console.log('Example usage:');
        console.log('   doc-fetch --url https://docs.python.org/3 --output docs.md --llm-txt\n');
        
        // Test if the command works
        try {
            execSync('doc-fetch --version', { encoding: 'utf8', stdio: 'pipe' });
            console.log('✅ Command verified working!\n');
        } catch (error) {
            console.log('⚠️  Command not found in current shell session.\n');
            console.log('Try running: hash -r  (to clear command cache)\n');
            console.log('Or restart your terminal.\n');
        }
    }
}

console.log('📚 Documentation: https://github.com/AlphaTechini/doc-fetch\n');
console.log('✨ Pro tip: Use --llm-txt flag to generate AI-friendly index files!\n');
