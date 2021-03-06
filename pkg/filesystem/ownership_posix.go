// +build !windows,!plan9

// TODO: Figure out what to do for Plan 9. It doesn't have syscall.Stat_t.

package filesystem

import (
	"os"
	"syscall"

	"github.com/pkg/errors"
)

// GetOwnership returns the owning user and group IDs from file metadata.
func GetOwnership(info os.FileInfo) (int, int, error) {
	if stat, ok := info.Sys().(*syscall.Stat_t); !ok {
		return 0, 0, errors.New("unable to extract raw stat information")
	} else {
		return int(stat.Uid), int(stat.Gid), nil
	}
}

// SetOwnership sets the owning user and group IDs for the specified path.
func SetOwnership(path string, uid, gid int) error {
	return os.Lchown(path, uid, gid)
}
