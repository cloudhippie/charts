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

func CommitChanges(c *Chart) error {
	if c.Bump == NoBump {
		return nil
	}

	if !c.Changed {
		return nil
	}

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

	msg := ""

	switch {
	case c.Bump == MajorBump:
		msg = fmt.Sprintf(
			"major(%s): raised chart version to %s",
			c.Meta.Name,
			c.Meta.Version,
		)
	case c.Bump == MinorBump:
		msg = fmt.Sprintf(
			"major(%s): raised chart version to %s",
			c.Meta.Name,
			c.Meta.Version,
		)
	case c.Bump == PatchBump:
		msg = fmt.Sprintf(
			"major(%s): raised chart version to %s",
			c.Meta.Name,
			c.Meta.Version,
		)
	default:
		return ErrNothingToCommit
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
