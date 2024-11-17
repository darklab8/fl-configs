package cfgtype

type ColorRGB struct {
	R int `yaml:"r"`
	G int `yaml:"g"`
	B int `yaml:"b"`
}

type SystemOwner struct {
	Name     string    `yaml:"name"`
	ColorHex string    `yaml:"color_hex"` // for example #2299F5
	ColorRGB *ColorRGB `yaml:"color_rgb"` // for example #2299F5

}

type SystemOwnerNick string
