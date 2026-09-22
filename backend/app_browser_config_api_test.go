package backend

import (
	"ant-chrome/backend/internal/browser"
	"ant-chrome/backend/internal/config"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func newBrowserSettingsTestApp(t *testing.T) *App {
	t.Helper()
	app := NewApp(t.TempDir())
	app.config = &config.Config{}
	app.config.Browser.UserDataRoot = "old-data"
	app.browserMgr = browser.NewManager(app.config, app.appRoot)
	return app
}

func TestSaveBrowserSettingsMigratesRelativeProfileDataRoot(t *testing.T) {
	app := newBrowserSettingsTestApp(t)
	profile := &browser.Profile{ProfileId: "profile-a", UserDataDir: "profile-a"}
	app.browserMgr.Profiles[profile.ProfileId] = profile

	oldDir := filepath.Join(app.appRoot, "old-data", profile.UserDataDir)
	newDir := filepath.Join(app.appRoot, "new-data", profile.UserDataDir)
	if err := os.MkdirAll(oldDir, 0o755); err != nil {
		t.Fatalf("create old profile directory: %v", err)
	}
	if err := os.WriteFile(filepath.Join(oldDir, "marker"), []byte("profile"), 0o644); err != nil {
		t.Fatalf("write profile marker: %v", err)
	}

	settings := app.GetBrowserSettings()
	settings.UserDataRoot = "new-data"
	if err := app.SaveBrowserSettings(settings); err != nil {
		t.Fatalf("SaveBrowserSettings() error = %v", err)
	}
	if app.config.Browser.UserDataRoot != "new-data" {
		t.Fatalf("user data root = %q, want new-data", app.config.Browser.UserDataRoot)
	}
	if _, err := os.Stat(oldDir); !os.IsNotExist(err) {
		t.Fatalf("old profile directory still exists, err = %v", err)
	}
	content, err := os.ReadFile(filepath.Join(newDir, "marker"))
	if err != nil || string(content) != "profile" {
		t.Fatalf("migrated marker = %q, err = %v", string(content), err)
	}
}

func TestSaveBrowserSettingsRejectsRunningProfileDataRootChange(t *testing.T) {
	app := newBrowserSettingsTestApp(t)
	profile := &browser.Profile{ProfileId: "profile-running", UserDataDir: "profile-running", Running: true}
	app.browserMgr.Profiles[profile.ProfileId] = profile

	settings := app.GetBrowserSettings()
	settings.UserDataRoot = "new-data"
	err := app.SaveBrowserSettings(settings)
	if err == nil || !strings.Contains(err.Error(), "正在运行") {
		t.Fatalf("SaveBrowserSettings() error = %v, want running-profile error", err)
	}
	if app.config.Browser.UserDataRoot != "old-data" {
		t.Fatalf("user data root changed after rejection: %q", app.config.Browser.UserDataRoot)
	}
}

func TestBrowserProfileDeleteStopsRunningProfileBeforeDelete(t *testing.T) {
	app := newBrowserSettingsTestApp(t)
	profile := &browser.Profile{ProfileId: "profile-running", ProfileName: "Running", UserDataDir: "profile-running", Running: true}
	app.browserMgr.Profiles[profile.ProfileId] = profile

	if err := app.BrowserProfileDelete(profile.ProfileId); err != nil {
		t.Fatalf("BrowserProfileDelete() error = %v", err)
	}
	if _, ok := app.browserMgr.Profiles[profile.ProfileId]; ok {
		t.Fatalf("profile remains after delete")
	}
	if profile.Running {
		t.Fatalf("profile was deleted while still marked running")
	}
}
