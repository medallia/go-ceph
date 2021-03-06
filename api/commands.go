package api

import (
	"fmt"
)

func (cc *CephClient) SetOsdFlag(flag string) error {
	_, err := cc.callApi(fmt.Sprintf("osd/set?key=%s", flag), "PUT")
	if err != nil {
		return err
	}
	return nil
}

func (cc *CephClient) UnsetOsdFlag(flag string) error {
	_, err := cc.callApi(fmt.Sprintf("osd/unset?key=%s", flag), "PUT")
	if err != nil {
		return err
	}
	return nil
}

func (cc *CephClient) SetOsdsState(state string, osdIds ...int) error {
	endpoint := fmt.Sprintf("osd/%s?", state)
	for _, id := range osdIds {
		endpoint += fmt.Sprintf("ids=%d&", id)
	}
	_, err := cc.callApi(endpoint, "PUT")
	return err
}

func (cc *CephClient) BlacklistOp(blacklistAddr string, op string) error {
	endpoint := fmt.Sprintf("osd/blacklist?blacklistop=%s&addr=%s", op, blacklistAddr)
	_, err := cc.callApi(endpoint, "PUT")
	return err
}

func (cc *CephClient) MdsFailover() error {
	_, err := cc.callApi("mds/fail?who=0", "PUT")
	return err
}
