package hosts

import (
	"regexp"
	"strings"

	"github.com/go-zoox/errors"
	"github.com/go-zoox/fs"
)

type Hosts struct {
	FilePath string
	Mapping  map[string]string
}

func New(filepath string) *Hosts {
	return &Hosts{
		FilePath: filepath,
		Mapping:  make(map[string]string),
	}
}

func (h *Hosts) Load() error {
	lines, err := fs.ReadFileLines(h.FilePath)
	if err != nil {
		return err
	}

	for _, line := range lines {
		// emty line
		if len(strings.TrimSpace(line)) == 0 {
			continue
			// comment line
		} else if strings.HasPrefix(line, "#") {
			continue
		}

		host := &Host{}
		if err := host.Parse(line); err != nil {
			return err
		}

		if host.IsIPv4 {
			for _, name := range host.Names {
				h.Mapping[name+":4"] = host.IP
			}
		} else {
			for _, name := range host.Names {
				h.Mapping[name+":6"] = host.IP
			}
		}
	}

	return nil
}

func (h *Hosts) LookUp(host string, typ ...int) (string, error) {
	typX := 4
	if len(typ) > 0 {
		typX = typ[0]
	}

	// IPv4
	if typX == 4 {
		return h.lookUpIPv4(host)
	} else if typX == 6 {
		return h.lookUpIPv6(host)
	} else {
		return "", errors.Errorf("invalid lookup type: %d, allow: 4, 6", typX)
	}
}

func (h *Hosts) lookUpIPv4(host string) (string, error) {
	if ip, ok := h.Mapping[host+":4"]; ok {
		return ip, nil
	}

	return "", errors.Errorf("not found: %s", host)
}

func (h *Hosts) lookUpIPv6(host string) (string, error) {
	if ip, ok := h.Mapping[host+":6"]; ok {
		return ip, nil
	}

	return "", errors.Errorf("not found: %s", host)
}

func (h *Hosts) Length() int {
	return len(h.Mapping)
}

type Host struct {
	IP     string
	Names  []string
	IsIPv4 bool
}

func (h *Host) Parse(text string) error {
	re := regexp.MustCompile(`\s+`)
	textX := re.ReplaceAllString(text, " ")
	parts := strings.Split(strings.TrimSpace(textX), " ")
	if len(parts) < 2 {
		return errors.Errorf("failed to parse invalid host: %s", text)
	}

	h.IP = parts[0]
	h.Names = parts[1:]
	h.IsIPv4 = strings.Contains(h.IP, ".")

	return nil
}
