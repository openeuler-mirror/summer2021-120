package docker

import (
	"context"
	"fmt"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/jhoonb/archivex"
)

func Build(commitID, repo string) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	tar := new(archivex.TarFile)
	tar.Create("/tmp/work-" + commitID + ".tar")
	tar.AddAll(repo+"/web-ui/script", true)
	tar.AddAll(repo+"/web-ui/docs", true)
	tar.AddAll(repo+"/web-ui/deploy", true)

	fp, err := os.Open(repo + "/web-ui/Dockerfile")
	if err != nil {
		panic(err)
	}
	tar.Add("Dockerfile", fp, nil)
	fp.Close()
	fp, err = os.Open(repo + "/web-ui/package.json")
	if err != nil {
		panic(err)
	}
	tar.Add("package.json", fp, nil)
	fp.Close()
	tar.Close()
	dockerBuildContext, err := os.Open("/tmp/work-" + commitID + ".tar")
	defer dockerBuildContext.Close()

	//cli.ImageBuild()

	options := types.ImageBuildOptions{
		SuppressOutput: true,
		Tags:           []string{"openeuler-webui:v." + commitID},
	}
	buildResponse, err := cli.ImageBuild(ctx, dockerBuildContext, options)
	if err != nil {
		fmt.Printf("%s", err.Error())
		return
	}
	defer buildResponse.Body.Close()
	fmt.Printf("********* %s **********\n", buildResponse.OSType)
}
