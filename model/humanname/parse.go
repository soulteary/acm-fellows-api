package humanname

import (
	"encoding/json"
	"fmt"

	"github.com/soulteary/acm-fellows-api/model/define"
)

func Parse(input string) (ret define.NameData, err error) {
	res, err := ParseNameByRpc(input)
	if err != nil {
		return ret, fmt.Errorf("Parsing Remote response failed: %v", err)
	}

	var name define.NameData
	err = json.Unmarshal([]byte(res), &name)
	if err != nil {
		return ret, fmt.Errorf("Parsing JSON failed: %v", err)
	}
	return name, nil
}
