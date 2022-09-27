package cleanpath

import (
    "testing"
    "strings"
)

func FuzzCleanPath(f *testing.F) {
    testcases := []string {"README", "../../otherplugin/../README", ""}
    for _, tc := range testcases {
        f.Add(tc)
    }
    f.Fuzz(func(t *testing.T, param string) {
        cleaned := CleanPath(param)

        if !strings.HasPrefix(cleaned, "/") { //CleanPath should enforce that the string starts with a /
          t.Errorf("Orginal input: %q, cleaned up: %q", param, cleaned)
        }

        if strings.Contains(cleaned, "/../") { //CleanPath should have removed all path traversal elements
          t.Errorf("Orginal input: %q, cleaned up: %q", param, cleaned)
        }
    })
}
