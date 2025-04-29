package render

import (
	"log"
	"regexp"
)

func (r *RealRenderer) renderCustom(md []byte, config *RenderConfig) ([]byte, error) {
	multiline, err := regexp.Compile(`(?m)^{ *([a-z]+)(?: *\"(.*)\")? *}\n?((?:.|\n)*?)\n?{/}$`)
	if err != nil {
		return []byte{}, err
	}
	log.Printf("%s\n", multiline.FindAll(md, -1))

	singleline, err := regexp.Compile(`(?m)^{ *([a-z]+)(?: *\"(.*)\")? */}$`)
	if err != nil {
		return []byte{}, err
	}
	log.Printf("%s\n", singleline.FindAll(md, -1))

	return md, nil
}
