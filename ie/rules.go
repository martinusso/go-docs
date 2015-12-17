package ie

type rules struct {
	initialMultiplier, finalMultiplier int
	substitute10, substitute11         int
}

var (
	rulesDefault = rules{
		initialMultiplier: 2,
		finalMultiplier:   9,
		substitute10:      0,
		substitute11:      0,
	}
)
