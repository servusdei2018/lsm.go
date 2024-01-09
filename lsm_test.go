package lsm

import (
	"strings"
	"testing"
)

const testLSM = `Begin4
Title: Fungimol
Version: 0.2.2
Entered-date: 2000-05-18
Description: Fungimol is an extensible system for designing atomic-scale objects. The intent is to eventually extend it to be a useful system for doing molecular nanotechnology design work. At the moment it's a PDB file viewer and Buckminsterfullerine editor.
Keywords: nanotechnology, hydrocarbon, fungimol, molecular dynamics, pdb, pdb viewer, graphics
Author: tim@infoscreen.com (Tim Freeman),
	brenner@eos.ncsu.edu (Donald Brenner)
Maintained-by: tim@infoscreen.com (Tim Freeman)
Primary-site: http://www.infoscreen.com/fungimol
	2M fungimol-0.2.2-0.i386.rpm
	180K brennermd-0.1.0-0.i386.rpm
	1.1M fungimol-0.2.2-0.src.rpm
	160K brennermd-0.1.0-0.src.rpm
	1.1M fungimol-0.2.2.tgz
	168K brennermd-0.1.0.tgz
	1K fungimol.lsm
Platforms: Built with GNU C++ compiler that supports anonymous namespaces. g++ 2.95.2 definitely works, 2.7.0 might work, prior to 2.7.0 probably won't work. Non-GNU
	C++ compiler might work too, but I haven't tried it. Also needs X window, 16 bits-per-pixel or more TrueColor display.
Copying-policy: Gnu Library General Public License (GLPL) 2.0
End`

func TestParse(t *testing.T) {
	if _, err := Parse(strings.NewReader(testLSM)); err != nil {
		t.Error(err)
	}
}
