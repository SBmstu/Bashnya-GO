package functions

func FindUniq(data []string) []string {
	result := make([]string, 0, len(data));
	res_map := make(map[string]int);

	for _, s := range data {
		res_map[s]++;
	}

	for key, value := range res_map {
		if value == 1 {
			result = append(result, key)
		}
	}

	return result;
}
