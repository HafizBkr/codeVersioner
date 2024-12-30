// internal/objects/commit.go
package objects

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

type Commit struct {
    Hash      string
    Message   string
    Timestamp time.Time
    Parent    string
    Files     map[string]string // filename -> content hash 
}

func NewCommit(message string, parent string, files map[string]string) *Commit {
    commit := &Commit{
        Message:   message,
        Timestamp: time.Now(),
        Parent:    parent,
        Files:     files,
    }
    
    h := sha1.New()
    h.Write([]byte(message + parent + commit.Timestamp.String()))
    commit.Hash = hex.EncodeToString(h.Sum(nil))
    
    return commit
}
// LoadCommit lit et désérialise un fichier commit en structure Commit
func LoadCommit(path string) (*Commit, error) {
	// Lire le fichier commit
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("could not read commit file: %v", err)
	}

	// Désérialiser les données en structure Commit
	var commit Commit
	if err := json.Unmarshal(data, &commit); err != nil {
		return nil, fmt.Errorf("could not unmarshal commit: %v", err)
	}

	return &commit, nil
}