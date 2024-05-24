package populate

var beerAndCiderProductTypes = []Product{
	{
		name:   "Stubby",
		volume: 375,
		price:  3500,
		abv:    5,
	},
	{
		name:   "Longneck",
		volume: 750,
		price:  6500,
		abv:    5,
	},
	{
		name:   "Can",
		volume: 375,
		price:  3000,
		abv:    5,
	},
	{
		name:   "6 Pack",
		volume: 375 * 6,
		price:  18000,
		abv:    5,
	},
	{
		name:   "Carton",
		volume: 375 * 24,
		price:  72000,
		abv:    5,
	},
}

var wineProductTypes = []Product{
	{
		name:   "Bottle",
		volume: 750,
		price:  1500,
		abv:    12,
	},
	{
		name:   "Half Bottle",
		volume: 375,
		price:  800,
		abv:    12,
	},
	{
		name:   "Case",
		volume: 750 * 12,
		price:  18000,
		abv:    12,
	},
}

var spiritAndPreMixProductTypes = []Product{
	{
		name:   "750mL",
		volume: 700,
		price:  4000,
		abv:    40,
	},
	{
		name:   "1L",
		volume: 1000,
		price:  5000,
		abv:    40,
	},
	{
		name:   "1.125L",
		volume: 1125,
		price:  5500,
		abv:    40,
	},
	{
		name:   "Case",
		volume: 700 * 12,
		price:  48000,
		abv:    40,
	},
}
