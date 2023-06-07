package model

type Unit string

const (
	// descriptive units
	TEESPOON   Unit = "tsp"
	TABLESPOON Unit = "tbsp"

	// mass units
	GRAMS     Unit = "g"
	KILOGRAMS Unit = "kg"

	// liquid units
	MILLILITER Unit = "ml"
	LITER      Unit = "l"
)

type Ingredient struct {
	Name         string  `bson:"name,omitempty"`
	NeededAmount float32 `bson:"neededAmount,omitempty"`
	Unit         Unit    `bson:"unit,omitempty"`
}
