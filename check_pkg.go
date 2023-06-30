package release

import (
	"errors"
	"os/exec"
)

type Pkgs struct {
	List []Pkg
}

type Pkg struct {
	Which string
	Name  string
}

func whichPkg(pkg string) (bool, error) {
	out, err := exec.Command("which", pkg).Output()
	if err != nil {
		return false, errors.New("error running \"which\" command")
	}
	if len(out) > 0 {
		return true, nil
	}

	return false, nil
}

func CheckPkg() (string, error) {
	// List of package managers
	pkgs := Pkgs{List: []Pkg{
		{Which: "dpkg", Name: "deb"},
		{Which: "rpm", Name: "rpm"},
	}}

	var pkgName string

	for _, pkg := range pkgs.List {
		resp, err := whichPkg(pkg.Which)
		if err != nil {
			return "", errors.New("error running whichPkg command")
		}
		if resp {
			pkgName = pkg.Name
			break
		}
	}

	return pkgName, nil
}
