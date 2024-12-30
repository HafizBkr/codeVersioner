// internal/storage/repository.go
package storage

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"myvcs/internal/objects"
)

type Repository struct {
    Path    string
    Head    string
    Stage   map[string]string
    Objects map[string]*objects.Commit
}

func InitRepository(path string) (*Repository, error) {
    repo := &Repository{
        Path:    path,
        Stage:   make(map[string]string),
        Objects: make(map[string]*objects.Commit),
    }

    dirs := []string{
        filepath.Join(path, ".vcs"),
        filepath.Join(path, ".vcs/objects"),
        filepath.Join(path, ".vcs/refs"),
    }

    for _, dir := range dirs {
        if err := os.MkdirAll(dir, 0755); err != nil {
            return nil, err
        }
    }

    return repo, nil
}

func (r *Repository) StageFile(path string) error {
    content, err := ioutil.ReadFile(path)
    if err != nil {
        return err
    }

    h := sha1.New()
    h.Write(content)
    hash := hex.EncodeToString(h.Sum(nil))

    r.Stage[filepath.Base(path)] = hash

    return ioutil.WriteFile(
        filepath.Join(r.Path, ".vcs/objects", hash),
        content,
        0644,
    )
}

func (r *Repository) Commit(message string) (*objects.Commit, error) {
    if len(r.Stage) == 0 {
        return nil, fmt.Errorf("nothing to commit")
    }

    commit := objects.NewCommit(message, r.Head, r.Stage)
    
    r.Objects[commit.Hash] = commit
    r.Head = commit.Hash
    
    r.Stage = make(map[string]string)
    
    return commit, nil
}