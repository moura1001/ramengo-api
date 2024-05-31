package utilapp

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type BRL int64

func ToBRL(v float64) BRL {
	return BRL(v * 100)
}

func (c BRL) Float64() float64 {
	v := float64(c)
	v = v / 100
	return v
}

func (c BRL) String() string {
	return fmt.Sprintf("R$%.2f", c.Float64())
}

func (c BRL) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Float64())
}

func (c *BRL) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), "\"")
	value, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}
	*c = ToBRL(value)
	return nil
}
