package caddy

import "testing"

func TestCaddyAPI_CreateWebsite(t *testing.T) {
	caddyAPI := NewCaddyAPI("http://3.114.44.103:32019", "/home/ubuntu/caddy")
	err := caddyAPI.CreateWebsite("mp4.bitheart.io", "")
	if err != nil {
		t.Errorf("CreateWebsite() error = %v", err)
	} else {
		t.Log("CreateWebsite() success")
	}
}
