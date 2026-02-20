#!/usr/bin/env node
const os = require('os');
const path = require('path');
const { execSync } = require('child_process');
const fs = require('fs');
const https = require('https');
const { pipeline } = require('stream');
const { promisify } = require('util');

const finished = promisify(pipeline);

// Security: Validate and sanitize output paths
function validateOutputPath(filePath) {
    // Resolve to absolute path
    const absPath = path.resolve(filePath);
    
    // Get current working directory
    const cwd = process.cwd();
    
    // Ensure the path is within the current directory or a subdirectory
    if (!absPath.startsWith(cwd)) {
        throw new Error('Output path must be within current directory or subdirectories');
    }
    
    // Block dangerous patterns
    if (filePath.includes('..') || filePath.includes('~')) {
        throw new Error('Relative paths with ".." or "~" are not allowed');
    }
    
    return absPath;
}

// Security: Validate URL
function validateURL(urlStr) {
    try {
        const url = new URL(urlStr);
        
        // Only allow HTTP/HTTPS
        if (url.protocol !== 'http:' && url.protocol !== 'https:') {
            throw new Error('Only HTTP and HTTPS URLs are allowed');
        }
        
        // Check for private IP ranges (basic SSRF protection)
        const hostname = url.hostname.toLowerCase();
        const privatePatterns = [
            '127.',      // localhost
            '192.168.',  // private network
            '10.',       // private network  
            '172.16.', '172.17.', '172.18.', '172.19.',
            '172.20.', '172.21.', '172.22.', '172.23.',
            '172.24.', '172.25.', '172.26.', '172.27.',
            '172.28.', '172.29.', '172.30.', '172.31.',
            'localhost',
            '::1'        // IPv6 localhost
        ];
        
        for (const pattern of privatePatterns) {
            if (hostname.startsWith(pattern)) {
                throw new Error('Private/internal URLs are not allowed');
            }
        }
        
        return true;
    } catch (error) {
        throw new Error(`Invalid URL: ${error.message}`);
    }
}

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
    
    // Validate the download URL
    validateURL(url);
    
    // Validate and sanitize the binary path
    const binaryPath = validateOutputPath(path.join(binDir, filename));
    
    console.log('📥 Downloading doc-fetch binary...');
    console.log(`   Platform: ${os.platform()} ${os.arch()}`);
    console.log(`   URL: ${url}`);
    
    try {
        const response = await new Promise((resolve, reject) => {
            const req = https.get(url, (res) => {
                if (res.statusCode === 404) {
                    reject(new Error('Binary not found for your platform'));
                } else if (res.statusCode !== 200) {
                    reject(new Error(`HTTP ${res.statusCode}: ${res.statusMessage}`));
                } else {
                    resolve(res);
                }
            }).on('error', reject);
            
            // Set timeout for security
            req.setTimeout(30000, () => {
                req.destroy(new Error('Download timeout'));
            });
        });
        
        // Create write stream with secure permissions
        const writeStream = fs.createWriteStream(binaryPath, { mode: 0o700 });
        await finished(response, writeStream);
        
        console.log('✅ Binary downloaded successfully!');
        return true;
    } catch (error) {
        console.error('❌ Failed to download binary:', error.message);
        console.log('🔄 Falling back to Go build from source...');
        
        // Fallback: build from source if Go is available
        try {
            execSync('go version', { stdio: 'pipe' });
            console.log('🏗️  Building from source...');
            
            // Validate build path
            const buildPath = validateOutputPath(path.join(binDir, 'doc-fetch'));
            execSync(`go build -o ${buildPath} ./cmd/docfetch`, { 
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