package release

import (
	"fmt"
	"testing"
)

const rcloneVer = "1.62.2"

func TestGithub(t *testing.T) {
	tests{
		{
			Input: "rclone/rclone", F: Github,
			Expected:    fmt.Sprintf("https://github.com/rclone/rclone/releases/download/v%s/rclone-v%s-linux-amd64.deb", rcloneVer, rcloneVer),
			ExpectedErr: nil,
		},
		{
			Input: "myztillx-test-repo", F: Github,
			Expected:    "",
			ExpectedErr: ErrImproperRepo,
		},
		{
			Input: "myztillx/test-repo", F: Github,
			Expected:    "",
			ExpectedErr: ErrNoAssets,
		},
	}.validate(t)
}
