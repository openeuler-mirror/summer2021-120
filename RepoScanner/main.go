package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"openeuler.org/repoScanner/docker"
	ci "openeuler.org/repoScanner/ci"
	cd "openeuler.org/repoScanner/cd"
)

func main() {

    ciInstance = ci.BuildCI()

	for {
	    image = ciInstance.Build()

		if image != nil  {
            cd.DeployWebUI(image)
		} else {
			time.Sleep(time.Duration(10) * time.Second)
		}

	}

}
