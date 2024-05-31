package model

import utilapp "github.com/moura1001/ramengo-api/src/util/app"

type Broth struct {
	Id            string      `json:"id"`
	ImageInactive string      `json:"imageInactive"`
	ImageActive   string      `json:"imageActive"`
	Name          string      `json:"name"`
	Description   string      `json:"description"`
	Price         utilapp.BRL `json:"price"`
}

func ListAllBroths() []Broth {
	return []Broth{
		{
			Id:            "1",
			ImageInactive: "https://tech.redventures.com.br/icons/salt/inactive.svg",
			ImageActive:   "https://tech.redventures.com.br/icons/salt/active.svg",
			Name:          "Salt",
			Description:   "Simple like the seawater, nothing more",
			Price:         utilapp.ToBRL(10),
		},
	}
}
