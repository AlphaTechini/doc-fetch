#!/usr/bin/env python3
"""
DocFetch CLI wrapper for Python.

This module provides a Python interface to the Go-based DocFetch binary.
It handles downloading the appropriate binary for your platform and
executing it with the provided arguments.
"""

import os
import sys
import subprocess
import platform
from pathlib import Path

# Get the directory where this script is located
SCRIPT_DIR = Path(__file__).parent
BIN_DIR = SCRIPT_DIR / "bin"
BINARY_NAME = None


def get_binary_name():
    """Get the appropriate binary name for the current platform."""
    system = platform.system().lower()
    machine = platform.machine().lower()
    
    # Map machine architectures
    arch_map = {
        'x86_64': 'amd64',
        'amd64': 'amd64', 
        'arm64': 'arm64',
        'aarch64': 'arm64'
    }
    
    arch = arch_map.get(machine, 'amd64')
    
    if system == 'windows':
        return f'doc-fetch_windows_{arch}.exe'
    elif system == 'darwin':
        return f'doc-fetch_darwin_{arch}'
    else:  # linux and others
        return f'doc-fetch_linux_{arch}'


def download_binary():
    """Download the appropriate binary from GitHub releases."""
    import urllib.request
    import ssl
    
    binary_name = get_binary_name()
    binary_path = BIN_DIR / binary_name
    
    # Create bin directory if it doesn't exist
    BIN_DIR.mkdir(exist_ok=True)
    
    # URL for the binary
    url = f"https://github.com/AlphaTechini/doc-fetch/releases/download/v1.0.0/{binary_name}"
    
    print(f"📥 Downloading doc-fetch binary for {platform.system()} {platform.machine()}...")
    print(f"   URL: {url}")
    
    try:
        # Create SSL context to handle certificates
        ssl_context = ssl.create_default_context()
        
        # Download the binary
        with urllib.request.urlopen(url, context=ssl_context) as response:
            with open(binary_path, 'wb') as f:
                f.write(response.read())
        
        # Make executable on Unix-like systems
        if platform.system() != 'Windows':
            os.chmod(binary_path, 0o755)
        
        print("✅ Binary downloaded successfully!")
        return binary_path
        
    except Exception as e:
        print(f"❌ Failed to download binary: {e}")
        print("💡 Please ensure you have internet access and can reach GitHub.")
        sys.exit(1)


def main():
    """Main entry point for the doc-fetch CLI."""
    global BINARY_NAME
    
    # Get binary path
    binary_name = get_binary_name()
    binary_path = BIN_DIR / binary_name
    
    # Download binary if it doesn't exist
    if not binary_path.exists():
        binary_path = download_binary()
    
    # Execute the binary with all arguments
    try:
        result = subprocess.run([str(binary_path)] + sys.argv[1:], check=False)
        sys.exit(result.returncode)
    except FileNotFoundError:
        print("❌ doc-fetch binary not found!")
        print("💡 This shouldn't happen. Please reinstall the package.")
        sys.exit(1)
    except KeyboardInterrupt:
        print("\n⚠️  Interrupted by user")
        sys.exit(130)
    except Exception as e:
        print(f"❌ Failed to execute doc-fetch: {e}")
        sys.exit(1)


if __name__ == "__main__":
    main()