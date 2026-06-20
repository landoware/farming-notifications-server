package main

import "time"

type Crop struct {
	Name     string
	Value    string
	Duration time.Duration
}

var cropsByGroup = map[CropGroup][]Crop{
	CropGroupAllotment: {
		{Name: "Potato", Value: "potato", Duration: 40 * time.Minute},
		{Name: "Onion", Value: "onion", Duration: 40 * time.Minute},
		{Name: "Cabbage", Value: "cabbage", Duration: 40 * time.Minute},
		{Name: "Tomato", Value: "tomato", Duration: 40 * time.Minute},
		{Name: "Sweetcorn", Value: "sweetcorn", Duration: 1 * time.Hour},
		{Name: "Strawberry", Value: "strawberry", Duration: 1 * time.Hour},
		{Name: "Watermelon", Value: "watermelon", Duration: 80 * time.Minute},
		{Name: "Snape Grass", Value: "snape_grass", Duration: 70 * time.Minute},
	},
	CropGroupFlower: {
		{Name: "Marigold", Value: "marigold", Duration: 20 * time.Minute},
		{Name: "Rosemary", Value: "rosemary", Duration: 20 * time.Minute},
		{Name: "Nasturtium", Value: "nasturtium", Duration: 20 * time.Minute},
		{Name: "Woad", Value: "woad", Duration: 20 * time.Minute},
		{Name: "Limpwurt", Value: "limpwurt", Duration: 20 * time.Minute},
		{Name: "White Lily", Value: "white_lily", Duration: 20 * time.Minute},
	},
	CropGroupHerb: {
		{Name: "Guam", Value: "guam", Duration: 80 * time.Minute},
		{Name: "Marrentill", Value: "marrentill", Duration: 80 * time.Minute},
		{Name: "Tarromin", Value: "tarromin", Duration: 80 * time.Minute},
		{Name: "Harralander", Value: "harralander", Duration: 80 * time.Minute},
		{Name: "Goutweed", Value: "goutweed", Duration: 80 * time.Minute},
		{Name: "Ranarr", Value: "ranarr", Duration: 80 * time.Minute},
		{Name: "Toadflax", Value: "toadflax", Duration: 80 * time.Minute},
		{Name: "Irit", Value: "irit", Duration: 80 * time.Minute},
		{Name: "Avantoe", Value: "avantoe", Duration: 80 * time.Minute},
		{Name: "Kwuarm", Value: "kwuarm", Duration: 80 * time.Minute},
		{Name: "Snapdragon", Value: "snapdragon", Duration: 80 * time.Minute},
		{Name: "Huasca", Value: "huasca", Duration: 80 * time.Minute},
		{Name: "Cadantine", Value: "cadantine", Duration: 80 * time.Minute},
		{Name: "Lantadyme", Value: "lantadyme", Duration: 80 * time.Minute},
		{Name: "Dwarf Weed", Value: "dwarf_weed", Duration: 80 * time.Minute},
		{Name: "Torstol", Value: "torstol", Duration: 80 * time.Minute},
	},
	CropGroupHops: {
		{Name: "Barley", Value: "barley", Duration: 40 * time.Minute},
		{Name: "Hammerstone", Value: "hammerstone", Duration: 40 * time.Minute},
		{Name: "Asgarnian", Value: "asgarnian", Duration: 50 * time.Minute},
		{Name: "Jute", Value: "jute", Duration: 50 * time.Minute},
		{Name: "Yanillian", Value: "yanillian", Duration: 1 * time.Hour},
		{Name: "Flax", Value: "flax", Duration: 1 * time.Hour},
		{Name: "Krandorian", Value: "krandorian", Duration: 70 * time.Minute},
		{Name: "Wildblood", Value: "wildblood", Duration: 80 * time.Minute},
		{Name: "Hemp", Value: "hemp", Duration: 80 * time.Minute},
		{Name: "Cotton", Value: "cotton", Duration: 100 * time.Minute},
	},
	CropGroupBush: {
		{Name: "Redberry", Value: "redberry", Duration: 100 * time.Minute},
		{Name: "Cadavaberry", Value: "cadavaberry", Duration: 2 * time.Hour},
		{Name: "Dwellberry", Value: "dwellberry", Duration: 140 * time.Minute},
		{Name: "Jangerberry", Value: "jangerberry", Duration: 160 * time.Minute},
		{Name: "Whiteberry", Value: "whiteberry", Duration: 160 * time.Minute},
		{Name: "Poison Ivy", Value: "poison_ivy", Duration: 160 * time.Minute},
	},
	CropGroupTree: {
		{Name: "Oak", Value: "oak", Duration: 160 * time.Minute},
		{Name: "Willow", Value: "willow", Duration: 4 * time.Hour},
		{Name: "Maple", Value: "maple", Duration: 320 * time.Minute},
		{Name: "Yew", Value: "yew", Duration: 400 * time.Minute},
		{Name: "Magic", Value: "magic", Duration: 8 * time.Hour},
	},
	CropGroupFruitTree: {
		{Name: "Apple", Value: "apple", Duration: 16 * time.Hour},
		{Name: "Banana", Value: "banana", Duration: 16 * time.Hour},
		{Name: "Orange", Value: "orange", Duration: 16 * time.Hour},
		{Name: "Curry", Value: "curry", Duration: 16 * time.Hour},
		{Name: "Pineapple", Value: "pineapple", Duration: 16 * time.Hour},
		{Name: "Papaya", Value: "papaya", Duration: 16 * time.Hour},
		{Name: "Palm", Value: "palm", Duration: 16 * time.Hour},
		{Name: "Dragonfruit", Value: "dragonfruit", Duration: 16 * time.Hour},
	},
	CropGroupCactus: {
		{Name: "Cactus", Value: "cactus", Duration: 560 * time.Minute},
		{Name: "Potato Cactus", Value: "potato_cactus", Duration: 70 * time.Minute},
	},
	CropGroupSeaweed: {
		{Name: "Giant Seaweed", Value: "giant_seaweed", Duration: 40 * time.Minute},
	},
	CropGroupMushroom: {
		{Name: "Mushroom", Value: "mushroom", Duration: 4 * time.Hour},
	},
	CropGroupBelladonna: {
		{Name: "Belladonna", Value: "belladonna", Duration: 320 * time.Minute},
	},
	CropGroupCalquat: {
		{Name: "Calquat", Value: "calquat", Duration: 1280 * time.Minute},
	},
	CropGroupCelastrus: {
		{Name: "Celastrus", Value: "celastrus", Duration: 800 * time.Minute},
	},
	CropGroupRedwood: {
		{Name: "Redwood", Value: "redwood", Duration: 6400 * time.Minute},
	},
	CropGroupSpiritTree: {
		{Name: "Spirit Tree", Value: "spirit_tree", Duration: 3840 * time.Minute},
	},
}

func cropsForGroup(group CropGroup) []Crop {
	return cropsByGroup[group]
}

func cropForGroup(group CropGroup, value string) (Crop, bool) {
	for _, crop := range cropsByGroup[group] {
		if crop.Value == value {
			return crop, true
		}
	}

	return Crop{}, false
}

func defaultCropForGroup(group CropGroup) (Crop, bool) {
	crops := cropsForGroup(group)
	if len(crops) == 0 {
		return Crop{}, false
	}

	return crops[0], true
}

func cropOptionRequired(group CropGroup) bool {
	crops := cropsForGroup(group)
	if len(crops) <= 1 {
		return false
	}

	firstDuration := crops[0].Duration
	for _, crop := range crops[1:] {
		if crop.Duration != firstDuration {
			return true
		}
	}

	return false
}
