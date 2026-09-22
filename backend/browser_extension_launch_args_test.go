package backend

import (
	"slices"
	"strings"
	"testing"
)

func TestBuildBrowserLaunchArgsDoesNotLoadExtensionsFromStartupFlags(t *testing.T) {
	args := buildBrowserLaunchArgs("profile-dir", 9222, "direct://", nil, nil, nil, nil, false)
	if slices.Contains(args, "--load-extension") || slices.Contains(args, "--disable-extensions-except") {
		t.Fatalf("args = %#v, production extensions must not be loaded from startup flags", args)
	}
}

func TestBuildBrowserLaunchArgsOmitsRemoteDebugWhenDisabled(t *testing.T) {
	args := buildBrowserLaunchArgs("profile-dir", 0, "direct://", nil, nil, nil, nil, false)
	for _, arg := range args {
		if strings.HasPrefix(arg, "--remote-debugging-port=") {
			t.Fatalf("args = %#v, remote debug should be omitted when port is disabled", args)
		}
	}
}

func TestBuildBrowserLaunchArgsIncludesRemoteDebugWhenEnabled(t *testing.T) {
	args := buildBrowserLaunchArgs("profile-dir", 9222, "direct://", nil, nil, nil, nil, false)
	if !slices.Contains(args, "--remote-debugging-port=9222") {
		t.Fatalf("args = %#v, want remote debug port", args)
	}
}
