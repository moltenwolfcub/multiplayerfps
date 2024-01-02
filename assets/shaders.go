package assets

func LoadShader(file string) (string, error) {

	embeddedShader, err := shaders.ReadFile("shaders/" + file)
	if err != nil {
		return "", err
	}
	return string(embeddedShader), nil
}

func MustLoadShader(file string) string {
	img, err := LoadShader(file)
	if err != nil {
		panic("Failed to load shader: " + err.Error())
	}
	return img
}

var (
	QuadTexture string = MustLoadShader("quadTexture.frag")
	TestVert    string = MustLoadShader("test.vert")
)
