package validation

import "errors"

func WhiteParam(attribute string, order string) (resultParam string, errr error) {
	whiteParam := []string{"name", "asc", "tail_length", "desc"}
	var resultAttribute string
	var resultOrder string
	if attribute != "" || order != "" {

		for _, param := range whiteParam {
			if attribute == param {
				resultAttribute += attribute
			}
			if order == param {
				resultOrder += order
			}
		}
		resultParam = resultAttribute + " " + resultOrder
		if resultParam == " " {
			return "", errors.New("400")
		} else {
			return "ORDER BY" + " " + resultParam, nil
		}
	}
	return "", nil
}
