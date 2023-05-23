package parse

import (
	"github.com/soulteary/acm-fellows-api/model/define"
	"github.com/soulteary/acm-fellows-api/model/humanname"
)

func GetNameAndYear(fellows []define.Fellow) (map[string]string, error) {
	names := make(map[string]string)
	for _, fellow := range fellows {
		name, err := humanname.Parse(fellow.Name)
		if err != nil {
			return nil, err
		}

		if len(name.Detail.First) == 1 {
			name.Detail.First = name.Detail.First + "."
		}

		if name.Detail.Middle == "" {
			names[name.Detail.First+" "+name.Detail.Last] = fellow.Year
		} else {
			if len(name.Detail.Middle) == 1 {
				name.Detail.Middle = name.Detail.Middle + "."
			}
			names[name.Detail.First+" "+name.Detail.Middle+" "+name.Detail.Last] = fellow.Year
		}
	}
	return names, nil
}
