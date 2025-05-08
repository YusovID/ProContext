package converter

import (
	"fmt"
	"strconv"
	"strings"
	"task_3/internal/model"
)

func ValutesToJSONValutes(valCurs *model.ValCurs) ([]model.JsonValute, []error) {
	if valCurs == nil {
		return nil, []error{fmt.Errorf("input ValCurs is nil")}
	}

	jsonValutes := []model.JsonValute{}
	errors := []error{}

	for _, valute := range valCurs.Valute {
		jsonValute := model.JsonValute{}

		err := valuteToJSONValute(valute, &jsonValute)
		if err != nil {
			errors = append(errors, err)
		} else {
			jsonValutes = append(jsonValutes, jsonValute)
		}
	}

	return jsonValutes, errors
}

func valuteToJSONValute(srcStruct model.Valute, dstStruct *model.JsonValute) error {
	dstStruct.Name = srcStruct.Name

	vunitRateStr := srcStruct.VunitRate
	vunitRateStr = strings.Replace(vunitRateStr, ",", ".", 1)
	value, err := strconv.ParseFloat(vunitRateStr, 64)
	if err != nil {
		return fmt.Errorf("failed to convert Value to float64: %v", err)
	} else {
		dstStruct.VunitRate = value
	}

	return nil
}
