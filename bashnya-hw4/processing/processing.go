package processing

import (
	"bashnya-hw4/functions"
	"bashnya-hw4/utils"
)

func ProcessData(data *utils.Data_t, cfg *utils.Config_t) []string {
	if cfg.C {
		return functions.CountAll(data, cfg);
	}
	if cfg.D {
		return functions.FindNonuniq(data, cfg);
	} else if cfg.U {
		return functions.FindUniq(data, cfg);
	} else {
		return functions.FindAll(data, cfg);
	}
}
