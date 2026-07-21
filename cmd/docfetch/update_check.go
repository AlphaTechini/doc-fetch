package main

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const (
	releasesURL         = "https://api.github.com/repos/AlphaTechini/doc-fetch/releases?per_page=100"
	updateCheckInterval = 24 * time.Hour
)

type releaseCache struct {
	CheckedAt time.Time `json:"checkedAt"`
	Versions  []string  `json:"versions"`
}

type githubRelease struct {
	TagName    string `json:"tag_name"`
	Draft      bool   `json:"draft"`
	Prerelease bool   `json:"prerelease"`
}

type semanticVersion struct {
	major int
	minor int
	patch int
}

func warnIfOutdated(currentVersion string) {
	if currentVersion == "dev" {
		return
	}

	versions, ok := releaseVersions()
	if !ok {
		return
	}

	ahead := 0
	for _, releaseVersion := range versions {
		if isNewerVersion(releaseVersion, currentVersion) {
			ahead++
		}
	}
	if ahead > 5 {
		_, _ = os.Stderr.WriteString("DocFetch " + currentVersion + " is more than five releases behind " + normalizeVersion(versions[0]) + ". If you have problems using the latest version, run `doc-fetch --doctor`.\n")
	}
}

func releaseVersions() ([]string, bool) {
	if cached, ok := cachedReleaseVersions(); ok {
		return cached, true
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, releasesURL, nil)
	if err != nil {
		return nil, false
	}
	request.Header.Set("Accept", "application/vnd.github+json")
	request.Header.Set("User-Agent", "doc-fetch")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, false
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, false
	}

	var releases []githubRelease
	if err := json.NewDecoder(response.Body).Decode(&releases); err != nil {
		return nil, false
	}

	versions := make([]string, 0, len(releases))
	for _, release := range releases {
		if !release.Draft && !release.Prerelease && release.TagName != "" {
			versions = append(versions, release.TagName)
		}
	}
	if len(versions) == 0 {
		return nil, false
	}

	writeReleaseCache(versions)
	return versions, true
}

func cachedReleaseVersions() ([]string, bool) {
	cachePath, ok := updateCachePath()
	if !ok {
		return nil, false
	}

	contents, err := os.ReadFile(cachePath)
	if err != nil {
		return nil, false
	}

	var cache releaseCache
	if err := json.Unmarshal(contents, &cache); err != nil || len(cache.Versions) == 0 || time.Since(cache.CheckedAt) >= updateCheckInterval {
		return nil, false
	}
	return cache.Versions, true
}

func writeReleaseCache(versions []string) {
	cachePath, ok := updateCachePath()
	if !ok {
		return
	}

	contents, err := json.Marshal(releaseCache{CheckedAt: time.Now(), Versions: versions})
	if err != nil {
		return
	}

	if err := os.MkdirAll(filepath.Dir(cachePath), 0o755); err != nil {
		return
	}
	_ = os.WriteFile(cachePath, contents, 0o600)
}

func updateCachePath() (string, bool) {
	directory, err := os.UserCacheDir()
	if err != nil {
		return "", false
	}
	return filepath.Join(directory, "doc-fetch", "release-check.json"), true
}

func normalizeVersion(value string) string {
	return strings.TrimPrefix(value, "v")
}

func isNewerVersion(candidate, current string) bool {
	candidateVersion, ok := parseVersion(candidate)
	if !ok {
		return false
	}
	currentVersion, ok := parseVersion(current)
	if !ok {
		return false
	}

	if candidateVersion.major != currentVersion.major {
		return candidateVersion.major > currentVersion.major
	}
	if candidateVersion.minor != currentVersion.minor {
		return candidateVersion.minor > currentVersion.minor
	}
	return candidateVersion.patch > currentVersion.patch
}

func parseVersion(value string) (semanticVersion, bool) {
	parts := strings.Split(normalizeVersion(value), ".")
	if len(parts) != 3 {
		return semanticVersion{}, false
	}

	major, majorErr := strconv.Atoi(parts[0])
	minor, minorErr := strconv.Atoi(parts[1])
	patch, patchErr := strconv.Atoi(parts[2])
	if majorErr != nil || minorErr != nil || patchErr != nil {
		return semanticVersion{}, false
	}
	return semanticVersion{major: major, minor: minor, patch: patch}, true
}
