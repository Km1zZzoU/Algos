//░██████╗░██╗████████╗  ██████╗░██╗░██████╗███████╗░█████╗░████████╗
//██╔════╝░██║╚══██╔══╝  ██╔══██╗██║██╔════╝██╔════╝██╔══██╗╚══██╔══╝
//██║░░██╗░██║░░░██║░░░  ██████╦╝██║╚█████╗░█████╗░░██║░░╚═╝░░░██║░░░
//██║░░╚██╗██║░░░██║░░░  ██╔══██╗██║░╚═══██╗██╔══╝░░██║░░██╗░░░██║░░░
//╚██████╔╝██║░░░██║░░░  ██████╦╝██║██████╔╝███████╗╚█████╔╝░░░██║░░░
//░╚═════╝░╚═╝░░░╚═╝░░░  ╚═════╝░╚═╝╚═════╝░╚══════╝░╚════╝░░░░╚═╝░░░

package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func assert(err error, message string) {
	if err != nil {
		fmt.Println(message)
		log.Fatal(err)
	}
}

func search(arrayhash []plumbing.Hash, repo *git.Repository, commandtest string) (plumbing.Hash, []int) {
	l := 0
	r := len(arrayhash) - 1
	slice := make([]int, 0, r+1)
	for i := 0; i < r+1; i++ {
		slice = append(slice, 0)
	}
	for i := range slice {
		slice[i] = 0
	}
	for l != r {
		indpivot := int((l + r) / 2)
		if checkoutcheck(repo, arrayhash[indpivot], commandtest) {
			slice[indpivot] = 1
			r = indpivot
		} else {
			l = indpivot + 1
			slice[indpivot] = -1
		}
	}
	if checkoutcheck(repo, arrayhash[l], commandtest) {
		slice[r] = 1
		return arrayhash[l-1], slice
	}
	slice[l] = -1
	return arrayhash[r], slice
}

func checkoutcheck(repo *git.Repository, hash plumbing.Hash, commandtest string) bool {

	worktree, err := repo.Worktree()
	if err != nil {
		fmt.Println("error with getting worktree")
		log.Fatal(err)
	}

	err = worktree.Checkout(&git.CheckoutOptions{
		Hash: hash,
	})
	assert(err, "error with checkout")

	_, err = exec.Command("bash", "-c", commandtest).CombinedOutput()
	if err != nil {
		return false
	}
	return true
}

func gitbisect(urlrep, hash1, hash2, commandtest string) {
	newDir := "dir for clone"
	err := os.Mkdir(newDir, 0755)
	if err != nil {
		er := os.RemoveAll(newDir)
		assert(er, "big trouble with creating dir")
		err := os.Mkdir(newDir, 0755)
		assert(err, "small trouble with creating dir")
	}

	cwd, err := filepath.Abs(newDir)
	assert(err, "error with getting path")

	repo, err := git.PlainClone(cwd, false, &git.CloneOptions{
		URL:      urlrep,
		Progress: os.Stdout,
	})
	assert(err, "error with Clone git")

	os.Chdir(newDir)

	commithash1 := plumbing.NewHash(hash1)
	commit1, err := repo.CommitObject(commithash1)
	assert(err, "error with getting commit")

	commithash2 := plumbing.NewHash(hash2)
	commit2, err := repo.CommitObject(commithash2)
	assert(err, "error with getting commit")

	iter, err := repo.Log(&git.LogOptions{
		Since: &commit1.Author.When,
		Until: &commit2.Author.When,
	})
	assert(err, "error with iter")

	var commitHashes []plumbing.Hash
	err = iter.ForEach(func(commit *object.Commit) error {
		commitHashes = append(commitHashes, commit.Hash)
		return nil
	})
	assert(err, "error with commitHashes")

	color.Cyan("between selected commits %d unchecking commit\n\n", len(commitHashes)-2)

	badbadhash, slice := search(commitHashes, repo, commandtest)

	j := 0
	for _, hash := range commitHashes {
		commitobj, err := repo.CommitObject(hash)
		assert(err, "error with convert hash to message")
		if slice[j] == 1 {
			color.Green("%s %s", hash.String(), commitobj.Message)
		}
		if slice[j] == -1 {
			color.Red("%s %s", hash.String(), commitobj.Message)
		}
		if slice[j] == 0 {
			color.White("%s %s", hash.String(), commitobj.Message)
		}
		j++
	}
	commitobj, err := repo.CommitObject(badbadhash)
	assert(err, "error with convert hash to message")

	color.Cyan("A commit that broke everything:\n\n%s %s", badbadhash.String(), commitobj.Message)
}

func main() {
	urlrep := "https://github.com/Km1zZzoU/test-repo-for-minka4"
	commandtest := "go build HelleWorld.go && ./HelleWorld"
	gitbisect(urlrep, "63d80fbc98f88b68447922fb2b00b75718889fce", "a3ed78fb996401df7841cc3a92d00112c06eb1ed", commandtest)
}
