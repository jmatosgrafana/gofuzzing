package relative_symlink

import (
    "testing"
    "strings"
    "path/filepath"
)

func FuzzSymlinks(f *testing.F) {
    testcases := []string {"README", "../otherplugin/README", "../otherplugin/../README"}
    for _, tc := range testcases {
        f.Add(tc)
    }
    f.Fuzz(func(t *testing.T, symlinkDestPath string) {
        output := isSymlinkRelativeTo("/base", symlinkDestPath, "/base/plugins/symlink.txt")
        expected := expectedResult("/base", symlinkDestPath, "/base/plugins/symlink.txt")

        //testing output && !expected could be enough: not approving something that should not
        if (output != expected) {
          t.Errorf("Input: %q, Output: %t, Expected: %t", symlinkDestPath, output, expected)
        }
    })
}

func expectedResult(base string, destpath string, origpath string) bool {
  if strings.HasPrefix(destpath, "/") {
    return false //naive implementation instead of filePath.IsAbs
  }

  merged := filepath.Join(filepath.Dir(origpath),destpath)
  if !strings.HasPrefix(merged, base) { //naive check of whether we stay in base folder
    return false
  }

  return true
}
