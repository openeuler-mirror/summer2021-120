package docker

import (
	"testing"
)

func TestSplit(t *testing.T) {

	tar := new(TarFile)
	tar.Create("/tmp/nodejs-distro.tar")
	tar.AddAll("1/21", true)
	tar.Close()

}
