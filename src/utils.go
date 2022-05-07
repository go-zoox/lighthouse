package lighthouse

import (
	"errors"
	"strings"
)

func GetWildcardHost(host string) ([]string, error) {
	parts := strings.Split(host, ".")
	if len(parts) < 2 {
		return nil, errors.New("invalid host: " + host)
	} else if len(parts) == 2 {
		return []string{"*." + host}, nil
	}

	var hosts []string
	for i := 1; i < len(parts)-1; i++ {
		hosts = append(hosts, "*."+strings.Join(parts[i:], "."))
	}

	return hosts, nil
}
