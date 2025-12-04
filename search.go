package gomanuf

import (
	"fmt"
	"strings"
)

func checkSlash28(mac string) string {
	mac = mac[:10] + "0:00:00/28"
	if manuf, ok := slash28[mac]; ok {
		return manuf
	}
	return ""
}

func checkSlash36(mac string) string {
	mac = mac[:13] + "0:00/36"
	if manuf, ok := slash36[mac]; ok {
		return manuf
	}
	return ""
}

func Search(mac string) (string, error) {
	mac = strings.Replace(strings.ToUpper(mac), "-", ":", -1)

	if !macRegex.MatchString(mac) {
		return "Invalid MAC address", fmt.Errorf("invalid MAC address format: %s", mac)
	}

	for address, manuf := range data {
		if strings.HasPrefix(mac, address) {
			if manuf == "IEEE Registration Authority" {
				if check28 := checkSlash28(mac); check28 != "" {
					return check28, nil
				}
				if check36 := checkSlash36(mac); check36 != "" {
					return check36, nil
				}
			}

			return manuf, nil
		}
	}

	return "unknown", nil
}
