package relative_symlink

import (
    "path/filepath"
    "strings"
)

// from https://github.com/grafana/grafana/pull/50537/files
// isSymlinkRelativeTo checks whether symlinkDestPath is relative to basePath.
// symlinkOrigPath is the path to file holding the symbolic link.
func isSymlinkRelativeTo(basePath string, symlinkDestPath string, symlinkOrigPath string) bool {
	if filepath.IsAbs(symlinkDestPath) {
		return false
	} else {
		fileDir := filepath.Dir(symlinkOrigPath)
		cleanPath := filepath.Clean(filepath.Join(fileDir, "/", symlinkDestPath))
		p, err := filepath.Rel(basePath, cleanPath)
		if err != nil {
			return false
		}

    if strings.HasPrefix(p, ".."+string(filepath.Separator)) {
			return false
		}
	}

	return true
}
