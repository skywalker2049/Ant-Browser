package backend

import "testing"

func TestBuildProxyLocationResolveResultUsesRawCountryCode(t *testing.T) {
	health := ProxyIPHealthResult{
		ProxyId: "proxy-1",
		Ok:      true,
		Country: "Hong Kong SAR China",
		City:    "Hong Kong",
		RawData: map[string]interface{}{
			"countryCode": "HK",
		},
	}

	result := buildProxyLocationResolveResult("proxy-1", health, "ip_health", "2026-07-20T00:00:00Z")
	if !result.Ok {
		t.Fatalf("result.Ok = false, want true: %+v", result)
	}
	if result.Lang != "zh-HK" {
		t.Fatalf("result.Lang = %q, want %q", result.Lang, "zh-HK")
	}
	if result.Timezone != "Asia/Hong_Kong" {
		t.Fatalf("result.Timezone = %q, want %q", result.Timezone, "Asia/Hong_Kong")
	}
	if result.TimezoneSource != "country" {
		t.Fatalf("result.TimezoneSource = %q, want country", result.TimezoneSource)
	}
	if result.Error != "" {
		t.Fatalf("result.Error = %q, want empty", result.Error)
	}
}

func TestBuildProxyLocationResolveResultPrefersRawTimezone(t *testing.T) {
	health := ProxyIPHealthResult{
		ProxyId: "proxy-us-west",
		Ok:      true,
		Country: "United States",
		City:    "",
		RawData: map[string]interface{}{
			"countryCode": "US",
			"timezone":    "America/Los_Angeles",
		},
	}

	result := buildProxyLocationResolveResult("proxy-us-west", health, "ip_health", "2026-07-20T00:00:00Z")
	if !result.Ok {
		t.Fatalf("result.Ok = false, want true: %+v", result)
	}
	if result.Timezone != "America/Los_Angeles" {
		t.Fatalf("result.Timezone = %q, want %q", result.Timezone, "America/Los_Angeles")
	}
	if result.TimezoneSource != "ip" {
		t.Fatalf("result.TimezoneSource = %q, want ip", result.TimezoneSource)
	}
}

func TestBuildProxyLocationResolveResultReadsNestedRawTimezone(t *testing.T) {
	health := ProxyIPHealthResult{
		ProxyId: "proxy-au-adelaide",
		Ok:      true,
		Country: "Australia",
		City:    "Adelaide",
		RawData: map[string]interface{}{
			"countryCode": "AU",
			"location": map[string]interface{}{
				"timezone": "Australia/Adelaide",
			},
		},
	}

	result := buildProxyLocationResolveResult("proxy-au-adelaide", health, "ip_health", "2026-07-20T00:00:00Z")
	if !result.Ok {
		t.Fatalf("result.Ok = false, want true: %+v", result)
	}
	if result.Timezone != "Australia/Adelaide" {
		t.Fatalf("result.Timezone = %q, want %q", result.Timezone, "Australia/Adelaide")
	}
	if result.TimezoneSource != "ip" {
		t.Fatalf("result.TimezoneSource = %q, want ip", result.TimezoneSource)
	}
}

func TestBuildProxyLocationResolveResultIgnoresInvalidRawTimezone(t *testing.T) {
	health := ProxyIPHealthResult{
		ProxyId: "proxy-us-east",
		Ok:      true,
		Country: "United States",
		RawData: map[string]interface{}{
			"countryCode": "US",
			"timezone":    "Not/A_Timezone",
		},
	}

	result := buildProxyLocationResolveResult("proxy-us-east", health, "ip_health", "2026-07-20T00:00:00Z")
	if result.Timezone != "America/New_York" {
		t.Fatalf("result.Timezone = %q, want %q", result.Timezone, "America/New_York")
	}
	if result.TimezoneSource != "country" {
		t.Fatalf("result.TimezoneSource = %q, want country", result.TimezoneSource)
	}
}

func TestDefaultProxyLocationOptionsCoverPriorityRegions(t *testing.T) {
	got := defaultProxyLocationOptions()
	seen := make(map[string]bool, len(got))
	for _, option := range got {
		seen[option.Label] = true
	}

	for _, countryCode := range []string{"US", "GB", "DE", "JP", "KR", "SG", "HK", "TW", "IN", "CA", "RU"} {
		option, ok := countryLocaleDefaults[countryCode]
		if !ok {
			t.Fatalf("countryLocaleDefaults[%q] is missing", countryCode)
		}
		if !seen[option.Label] {
			t.Errorf("default proxy location options do not include %s", countryCode)
		}
	}
}
