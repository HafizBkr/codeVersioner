// internal/objects/commit.go
package objects

import (
    "crypto/sha1"
    "encoding/hex"
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