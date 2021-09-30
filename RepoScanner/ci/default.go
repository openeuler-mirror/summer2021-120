package ci

import (
	"fmt"
	"log"
	"os"
	"time"
)

type CI interface {
	Build(string)
}



func BuildCI() *CI {
	if(os.Getenv("CI_TYPE") == "custom"){
		return &CustomCI{
			Image: os.Getenv("CustomCI_IMAGE"),
			Params: os.Getenv("CustomCI_PARAMS")
		}
	} else{
	    return &InnerCI{
        			Owner: os.Getenv("GITEE_OWNER"),
                    Repo: os.Getenv("GITEE_REPO"),
                    Branch: os.Getenv("BRANCH"),
                    AccessToken: os.Getenv("GITEE_ACCESS_TOKEN"),
                    WorkDir: os.Getenv("GITEE_OWNER")
        		}
	}

}
///
///
type CustomCI struct{
	Image string
	Params string
}


type InnerCI struct {
	Owner        string
	Repo         string
	Branch       string
	AccessToken  string
	WorkDir      string
	LastCommitID string
}

func (c *InnerCI) Build(string) {

	if c.WorkDir == "" {
		c.WorkDir = "repo"
	}

	lastCommitID := os.Getenv("BASE_COMMIT")
	gitee := git.NewGitee(c.Owner, c.Repo, c.Branch, c.AccessToken, c.WorkDir)

	if gitee == nil {
		log.Panic()
	}
	result := gitee.PeekLastCommit()

	fmt.Println(result.Name, result.Commit.Sha)

	if lastCommitID == "" {
		lastCommitID = gitee.Head()
		return docker.Build(lastCommitID, "repo")
	}
	if result != nil && gitee.Head() != result.Commit.Sha {
		lastCommitID = result.Commit.Sha

		gitee.Checkout(lastCommitID)

		return docker.Build(lastCommitID, "repo")
	} else {
		//time.Sleep(time.Duration(10) * time.Second)
		return ""
	}

}


func (c *CustomCI) Build(string){
	return docker.Run(c.Image,c.Params)
}