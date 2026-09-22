package backend

import (
	"strings"
	"testing"

	"ant-chrome/backend/internal/browser"
)

func TestGetDebugPortRejectsProfileWithoutRemoteDebug(t *testing.T) {
	app := NewApp(t.TempDir())
	app.browserMgr = browser.NewManager(nil, app.appRoot)
	app.browserMgr.Profiles["profile-no-debug"] = &browser.Profile{
		ProfileId:          "profile-no-debug",
		Running:            true,
		RemoteDebugEnabled: false,
		DebugPort:          0,
		DebugReady:         false,
	}

	_, err := app.getDebugPort("profile-no-debug")
	if err == nil || !strings.Contains(err.Error(), "未开启远程调试") {
		t.Fatalf("getDebugPort error = %v, want explicit remote-debug-disabled error", err)
	}
}
