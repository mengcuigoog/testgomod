package utils

import "path/filepath"

func CleanPath(p string) string {
    return filepath.Clean(p)
}


func JoinPath(a,b string) string {
    return filepath.Join(a,b)
}

