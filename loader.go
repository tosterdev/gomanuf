package gomanuf

import (
	"fmt"
	"strings"
)

func loadManufData() {
	data = make(map[string]string)
	slash28 = make(map[string]string)
	slash36 = make(map[string]string)

	file, err := manufFS.ReadFile("manuf.txt")
	if err != nil {
		panic(fmt.Sprintf("failed to load embedded manuf.txt: %v", err))
	}

	lines := strings.Split(string(file), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		line = strings.Replace(line, "\t\t", "\t", -1)
		fields := strings.Split(line, "\t")

		if len(fields) < 2 {
			continue
		}

		mac := fields[0]
		manuf := fields[1]
		if len(fields) > 2 {
			manuf = fields[2]
		}

		if strings.Contains(mac, ":00/28") {
			slash28[mac] = manuf
		} else if strings.Contains(mac, ":00/36") {
			slash36[mac] = manuf
		}

		data[mac] = manuf
	}
}
