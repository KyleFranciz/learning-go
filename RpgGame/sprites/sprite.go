package sprites

type Sprite struct {
	// sprite for character core parameters
	name        string
	health      int
	attackPower int
	agility     int
}

func NewSprite(name string, health int, attackPower int, agility int) Sprite {
	return Sprite{
		name:        name,
		health:      health,
		attackPower: attackPower,
		agility:     agility,
	}
}

func (sprite *Sprite) GetName() string {
	// return the name of the sprite
	// NOTE: sprite has to be noted in the function so that the pointer points at the class and returns the current name of the sprite
	return sprite.name
}

func (sprite *Sprite) GetHealth() int {
	// get the health of the character
	return sprite.health
}
