"""
DocFetch - Dynamic documentation fetching CLI
Converts entire documentation sites to single markdown files for AI/LLM consumption.
"""

import os
import sys
import platform
from setuptools import setup, find_packages
from setuptools.command.install import install
import urllib.request
import tarfile
import zipfile
import stat
import shutil

# Determine binary name based on platform
def get_binary_name():
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
    else:
        return f'doc-fetch_linux_{arch}'

class InstallCommand(install):
    """Custom install command to download and install the Go binary."""
    
    def run(self):
        # Run the standard install first
        install.run(self)
        
        # Download and install the Go binary
        self.download_binary()
    
    def download_binary(self):
        """Download the appropriate Go binary for the current platform."""
        binary_name = get_binary_name()
        download_url = f"https://github.com/AlphaTechini/doc-fetch/releases/download/v1.0.0/{binary_name}"
        
        # Destination path
        install_dir = os.path.join(self.install_scripts, 'doc-fetch-bin')
        os.makedirs(install_dir, exist_ok=True)
        
        if platform.system() == 'Windows':
            binary_path = os.path.join(install_dir, 'doc-fetch.exe')
        else:
            binary_path = os.path.join(install_dir, 'doc-fetch')
        
        print(f"Downloading DocFetch binary from {download_url}")
        print(f"Platform: {platform.system()} {platform.machine()}")
        
        try:
            # Download the binary
            with urllib.request.urlopen(download_url) as response:
                with open(binary_path, 'wb') as f:
                    f.write(response.read())
            
            # Make executable on Unix-like systems
            if platform.system() != 'Windows':
                st = os.stat(binary_path)
                os.chmod(binary_path, st.st_mode | stat.S_IEXEC)
            
            print(f"Successfully installed DocFetch binary to {binary_path}")
            
        except Exception as e:
            print(f"Failed to download binary: {e}")
            print("Falling back to Go build from source...")
            self.build_from_source()
    
    def build_from_source(self):
        """Build from Go source if binary download fails."""
        try:
            # Check if Go is installed
            import subprocess
            subprocess.run(['go', 'version'], check=True, capture_output=True)
            
            # Build from source
            install_dir = os.path.join(self.install_scripts, 'doc-fetch-bin')
            source_dir = os.path.join(os.path.dirname(__file__), '..')
            
            if platform.system() == 'Windows':
                binary_path = os.path.join(install_dir, 'doc-fetch.exe')
                subprocess.run(['go', 'build', '-o', binary_path, './cmd/docfetch'], 
                             cwd=source_dir, check=True)
            else:
                binary_path = os.path.join(install_dir, 'doc-fetch')
                subprocess.run(['go', 'build', '-o', binary_path, './cmd/docfetch'], 
                             cwd=source_dir, check=True)
            
            # Make executable
            if platform.system() != 'Windows':
                st = os.stat(binary_path)
                os.chmod(binary_path, st.st_mode | stat.S_IEXEC)
                
            print(f"Successfully built DocFetch from source: {binary_path}")
            
        except Exception as e:
            print(f"Failed to build from source: {e}")
            print("Please install Go (https://golang.org/dl/) or ensure the binary is available for your platform.")

# Read README for long description
with open("README.md", "r", encoding="utf-8") as fh:
    long_description = fh.read()

setup(
    name="doc-fetch",
    version="1.0.0",
    author="AlphaTechini",
    author_email="rehobothokoibu@gmail.com",
    description="Dynamic documentation fetching CLI that converts entire documentation sites to single markdown files for AI/LLM consumption",
    long_description=long_description,
    long_description_content_type="text/markdown",
    url="https://github.com/AlphaTechini/doc-fetch",
    packages=find_packages(),
    classifiers=[
        "Development Status :: 5 - Production/Stable",
        "Intended Audience :: Developers",
        "Intended Audience :: Information Technology",
        "License :: OSI Approved :: MIT License",
        "Operating System :: OS Independent",
        "Programming Language :: Python :: 3",
        "Programming Language :: Python :: 3.7",
        "Programming Language :: Python :: 3.8",
        "Programming Language :: Python :: 3.9",
        "Programming Language :: Python :: 3.10",
        "Programming Language :: Python :: 3.11",
        "Topic :: Documentation",
        "Topic :: Internet :: WWW/HTTP",
        "Topic :: Software Development :: Documentation",
        "Topic :: Utilities",
    ],
    python_requires=">=3.7",
    install_requires=[],
    entry_points={
        "console_scripts": [
            "doc-fetch=doc_fetch.cli:main",
        ],
    },
    cmdclass={
        'install': InstallCommand,
    },
    include_package_data=True,
    zip_safe=False,
)