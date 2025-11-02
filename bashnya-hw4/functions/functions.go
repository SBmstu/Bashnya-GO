package functions

import "bashnya-hw4/utils"

func FindUniq(data []string) []string {
	result := make([]string, 0, len(data));
	res_map := utils.GenMapWithCount(data);

	for key, value := range res_map {
		if value == 1 {
			result = append(result, key)
		}
	}

	return result;
}

func FindNonuniq(data []string) []string {
	var result []string;
	res_map := utils.GenMapWithCount(data);

	for key, value := range res_map {
		if value != 1 {
			result = append(result, key)
		}
	}

	return result;
}
