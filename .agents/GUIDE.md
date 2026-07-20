# Project Guide

## Release Contract

- Release tags use the `v*` format.
- GitHub Releases must include Linux x64 and ARM64, macOS Intel and Apple Silicon, and Windows x64 and ARM64 binaries.
- The binary release workflow builds and verifies every asset before publishing the release.
- Existing tag-triggered npm and PyPI workflows publish their packages independently.
