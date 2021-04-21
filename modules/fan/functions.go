package fan

import (
	"fmt"
	"github.com/kamushadenes/swayit/common"
	"github.com/ssimunic/gosensors"
	"strconv"
	"strings"
)

func sum(input ...int64) int64 {
	var s int64
	s = 0

	for i := range input {
		s += input[i]
	}

	return s
}

func run(w *common.WaybarOutput) error {
	sensors, err := gosensors.NewFromSystem()
	if err != nil {
		return err
	}

	var values []int64
	
	for chip := range sensors.Chips {
		for key, value := range sensors.Chips[chip] {
			if strings.HasPrefix(key, "fan") {
				v, err := strconv.ParseInt(strings.Fields(value)[0], 10, 64)
				if err == nil {
					values = append(values, v)
					w.Tooltip += fmt.Sprintf("\uf021 %s: %d RPM\n\n", key, v)
				}
			}
		}
	}

	if len(values) > 0 {
		w.Text = fmt.Sprintf("\uf021 %d", sum(values...)/int64(len(values)))
	}
	
	return nil
}
