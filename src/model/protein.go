package model

import utilapp "github.com/moura1001/ramengo-api/src/util/app"

type Protein struct {
	Id            string      `json:"id"`
	ImageInactive string      `json:"imageInactive"`
	ImageActive   string      `json:"imageActive"`
	Name          string      `json:"name"`
	Description   string      `json:"description"`
	Price         utilapp.BRL `json:"price"`
}

func ListAllProteins() []Protein {
	return []Protein{
		{
			Id:            "1",
			ImageInactive: "https://tech.redventures.com.br/icons/pork/inactive.svg",
			ImageActive:   "https://tech.redventures.com.br/icons/pork/active.svg",
			Name:          "Chasu",
			Description:   "A sliced flavourful pork meat with a selection of season vegetables.",
			Price:         utilapp.ToBRL(10),
		},
	}
}
