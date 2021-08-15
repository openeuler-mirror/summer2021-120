package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"openeuler.org/repoScanner/docker"
	"openeuler.org/repoScanner/git"
)

// Retrieve remote tags without cloning repository
//"4527ddfc25c0a9f7c48b8483bb5bce59"
func main() {

	owner := os.Getenv("GITEE_OWNER")
	repo := os.Getenv("GITEE_REPO")
	branch := os.Getenv("BRANCH")
	accessToken := os.Getenv("GITEE_ACCESS_TOKEN")
	workDir := os.Getenv("GITEE_OWNER")

	if workDir == "" {
		workDir = "repo"
	}

	lastCommitID := os.Getenv("BASE_COMMIT")
	gitee := git.NewGitee(owner, repo, branch, accessToken, workDir)

	if gitee == nil {
		log.Panic()
	}
	for {
		result := gitee.PeekLastCommit()

		fmt.Println(result.Name, result.Commit.Sha)

		if lastCommitID == "" {
			lastCommitID = gitee.Head()
			docker.Build(lastCommitID, "repo")
		}
		if result != nil && gitee.Head() != result.Commit.Sha {
			lastCommitID = result.Commit.Sha

			gitee.Checkout(lastCommitID)

			docker.Build(lastCommitID, "repo")
		} else {
			time.Sleep(time.Duration(10) * time.Second)
		}

	}

}
