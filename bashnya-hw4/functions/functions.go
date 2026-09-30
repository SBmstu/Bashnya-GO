package functions

import (
	"bashnya-hw4/utils"
	"strconv"
	"strings"
)

func FindAll(data *utils.Data_t, cfg *utils.Config_t) []string {
	result := make([]string, 0, data.Count);
	res_map := utils.GenMapWithCount(data, cfg);

	for key := range res_map {
		result = append(result, key)
	}

	return result;
}

func FindUniq(data *utils.Data_t, cfg *utils.Config_t) []string {
	result := make([]string, 0, data.Count);
	res_map := utils.GenMapWithCount(data, cfg);

	for _, value := range res_map {
		if value.Count == 1 {
			result = append(result, data.Original[value.I])
		}
	}

	return result;
}

func FindNonuniq(data *utils.Data_t, cfg *utils.Config_t) []string {
	result := make([]string, 0, data.Count);
	res_map := utils.GenMapWithCount(data, cfg);

	for _, value := range res_map {
		if value.Count != 1 {
			result = append(result, data.Original[value.I])
		}
	}

	return result;
}

func CountAll(data *utils.Data_t, cfg *utils.Config_t) []string {
	result := make([]string, 0, data.Count);
	res_map := utils.GenMapWithCount(data, cfg);

	for key, value := range res_map {
		result = append(result, strconv.Itoa(value.Count) + " " + key);
	}
	
	return result;
}

func MakeDataLower(data []string) []string {
	lower_data := make([]string, len(data))

	for i, s := range data {
		lower_data[i] = strings.ToLower(s);
	}

	return lower_data;
}
