package git

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

type GiteeBranch struct {
	Name   string            `json:"name"`
	Commit GiteeCommitstruct `json:"commit"`
}

type GiteeCommitstruct struct {
	Sha string `json:"sha"`
}

type Gitee struct {
	GiteeOwner       string `json:"giteeOwner,omitempty"`
	GiteeRepo        string `json:"giteeRepo,omitempty"`
	Branch           string `json:"branch,omitempty"`
	GiteeAccessToken string `json:"accessToken,omitempty"` //access_token
	Repository       *git.Repository
}

func NewGitee(owner, repo, branch, token, directory string) *Gitee {

	//https://gitee.com/openeuler/website-v2.git
	repository, err := git.PlainClone(directory, false, &git.CloneOptions{
		URL:           "https://gitee.com/" + owner + "/" + repo + ".git",
		ReferenceName: plumbing.ReferenceName("refs/heads/" + branch),
		SingleBranch:  true,
		Progress:      os.Stdout,
	})

	if err != nil {
		fmt.Println(err)
		repository, err = git.PlainOpen(directory)
	}
	return &Gitee{
		GiteeOwner:       owner,
		GiteeRepo:        repo,
		Branch:           branch,
		GiteeAccessToken: token,
		Repository:       repository,
	}
}
func (g *Gitee) PeekLastCommit() *GiteeBranch {
	url := "https://gitee.com/api/v5/repos/" + g.GiteeOwner + "/" + g.GiteeRepo + "/branches/" + g.Branch

	if g.GiteeAccessToken != "" {
		url += "?access_token=" + g.GiteeAccessToken
	}

	fmt.Printf("http.Get(url)r %s\n", url)
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return nil
	}

	data, _ := ioutil.ReadAll(response.Body)

	var result = GiteeBranch{}
	err = json.Unmarshal(data, &result)

	return &result

}

func (g *Gitee) Head() string {
	r := g.Repository
	ref, err := r.Head()

	if err != nil {
		return ""
	}

	fmt.Println(ref.Hash())

	return ref.Hash().String()
}

func (g *Gitee) Checkout(commit string) error {
	r := g.Repository

	w, err := r.Worktree()
	if err != nil {
		return err
	}

	return w.Pull(&git.PullOptions{RemoteName: "origin"})
}
