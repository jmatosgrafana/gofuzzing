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
        ok := isSymlinkRelativeTo("/base", symlinkDestPath, "/base/plugins/symlink.txt")
        expected := expectedResult("/base", symlinkDestPath, "/base/plugins/symlink.txt")

        //testing ok && !expected could be enough: not approving something that should not
        if (ok != expected) {
          t.Errorf("OK: %q, Expected: %t", symlinkDestPath, expected)
        }
    })
}

func expectedResult(base string, destpath string, origpath string) bool {
  if strings.HasPrefix(destpath, "/") {
    return false //naive implementation, should also check it's not base ? rather filepath Abs?
  }

  merged := filepath.Join(filepath.Dir(origpath),destpath)
  if !strings.HasPrefix(merged, base) { //naive check of whether we stay in base folder
    return false
  }

  return true
}
