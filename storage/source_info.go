package storage

import (
	"github.com/go-git/go-git/v5/plumbing"

	"github.com/ava-labs/apm/types"
)

// SourceInfo represents a repository, its source, and the last synced commit.
type SourceInfo struct {
	Alias  string        `yaml:"alias"`
	URL    string        `yaml:"url"`
	Commit plumbing.Hash `yaml:"commit"`
}

// List is a list of repositories that support a single plugin alias.
// e.g. foo/plugin, bar/plugin => plugin: [foo, bar]
type List struct {
	Repositories []string `yaml:"repositories"`
}

// Definition stores a plugin definition alongside the plugin-repository's commit
// it was downloaded from.
// TODO gc plugins
type Definition[T types.Definition] struct {
	Definition T             `yaml:"definition"`
	Commit     plumbing.Hash `yaml:"commit"`
}