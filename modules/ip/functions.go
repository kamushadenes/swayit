package ip

import (
	"fmt"
	"github.com/kamushadenes/swayit/common"
)

func GetIPInfo(ip string) (*IPInfo, error) {
	var info IPInfo

	err := common.GetJson(fmt.Sprintf("http://ip-api.com/json/%s?fields=33292287", ip), &info)
	if err != nil {
		return nil, err
	}
	
	return &info, nil
}