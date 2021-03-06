package dsl

import (
	"github.com/antonioalfa22/egida/internal/info"
	"github.com/antonioalfa22/go-utils/collections"
	"strings"
)

type Comparation struct {
	Host string
	Value1 interface{}
	Operator string
	Value2 interface{}
}

func (c Comparation) Compare() bool {
	query := strings.Split(c.Value1.(string), ".")[0]
	value := strings.Split(c.Value1.(string), ".")[1]
	switch strings.ToLower(query) {
	case "services":
		if strings.ToLower(c.Value2.(string)) == "stopped" {
			alls, _ := info.GetStoppedServices([]string{c.Host})
			result := alls[0]
			return collections.Find(result.Lines, func(x string) bool {
				return strings.Split(x, "->")[0] == value
			}) != nil
		} else if strings.ToLower(c.Value2.(string)) == "running" {
			alls, _ := info.GetRunningServices([]string{c.Host})
			result := alls[0]
			return collections.Find(result.Lines, func(x string) bool {
				return strings.Split(x, "->")[0] == value
			}) != nil
		} else {
			alls, _ := info.GetAllServices([]string{c.Host})
			result := alls[0]
			return collections.Find(result.Lines, func(x string) bool {
				return strings.Split(x, "->")[0] == value
			}) != nil
		}
	case "packages":
		if strings.ToLower(c.Value2.(string)) == "installed" {
			alls, _ := info.GetAllPackages([]string{c.Host})
			result := alls[0]
			return collections.Find(result.Lines, func(x string) bool {
				return x == value
			}) != nil
		}
	case "hardscores":
		return false
	}
	return false
}
