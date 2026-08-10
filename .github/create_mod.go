// Command create_mod packages a module directory into a Go module zip in the
// exact layout served by proxy.golang.org, using golang.org/x/mod/zip.
//
// Usage: create_mod <module-path> <version> <dir> <output.zip>
package main

import (
	"fmt"
	"os"

	"golang.org/x/mod/module"
	"golang.org/x/mod/zip"
)

func main() {
	if len(os.Args) != 5 {
		fmt.Fprintln(os.Stderr, "usage: create_mod <module-path> <version> <dir> <output.zip>")
		os.Exit(2)
	}
	modPath, version, dir, out := os.Args[1], os.Args[2], os.Args[3], os.Args[4]

	f, err := os.Create(out)
	if err != nil {
		fmt.Fprintln(os.Stderr, "create output:", err)
		os.Exit(1)
	}
	defer f.Close()

	m := module.Version{Path: modPath, Version: version}
	if err := zip.CreateFromDir(f, m, dir); err != nil {
		fmt.Fprintln(os.Stderr, "CreateFromDir:", err)
		os.Exit(1)
	}
	fmt.Println("wrote", out)
}
