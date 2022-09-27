package cleanpath

import (
    "path/filepath"
)

func CleanPath(param string) string {
  	return filepath.Clean(filepath.Join("/", param))
}
