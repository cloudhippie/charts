package chart

import (
	"errors"
	"fmt"
	"path"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

var (
	ErrNothingToCommit = errors.New("nothing to commit")
)

func CommitMessage(c *Chart, msg string) error {
	r, err := git.PlainOpen(
		".",
	)

	if err != nil {
		return err
	}

	w, err := r.Worktree()

	if err != nil {
		return err
	}

	_, err = w.Add(
		path.Dir(
			c.Path,
		),
	)

	if err != nil {
		return err
	}

	_, err = w.Status()

	if err != nil {
		return err
	}

	_, err = w.Commit(
		msg,
		&git.CommitOptions{
			Author: &object.Signature{
				Name:  "GitHub Actions",
				Email: "github@cloudhippie.de",
				When:  time.Now(),
			},
		},
	)

	if err != nil {
		return err
	}

	return nil
}

func CommitBump(c *Chart) error {
	if c.Bump == NoBump {
		return nil
	}

	if !c.Changed {
		return nil
	}

	switch c.Bump {
	case MajorBump:
		return CommitMessage(
			c,
			fmt.Sprintf(
				"major(%s): raised chart version to %s",
				c.Meta.Name,
				c.Meta.Version,
			),
		)
	case MinorBump:
		return CommitMessage(
			c,
			fmt.Sprintf(
				"minor(%s): raised chart version to %s",
				c.Meta.Name,
				c.Meta.Version,
			),
		)
	case PatchBump:
		return CommitMessage(
			c,
			fmt.Sprintf(
				"patch(%s): raised chart version to %s",
				c.Meta.Name,
				c.Meta.Version,
			),
		)
	default:
		return ErrNothingToCommit
	}
}
