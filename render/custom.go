package render

import (
	"bytes"
	"log"
)

func (r *RealRenderer) renderCustom(md []byte, config *RenderConfig) []byte {
	md = r.renderSingleline(md, config)
	md = r.renderMultiline(md, config)
	return md
}

func (r *RealRenderer) renderSingleline(md []byte, config *RenderConfig) []byte {
	match := r.singleline.FindSubmatch(md)
	if match == nil {
		return md
	}

	html := r.renderSingleHTML(match[1], match[2], config)

	md = bytes.Replace(md, match[0], html, 1)
	return r.renderSingleline(md, config)
}

func (r *RealRenderer) renderSingleHTML(tag, param []byte, config *RenderConfig) []byte {
	log.Printf("\ntag: %s\nparam: %s\n", tag, param)
	return []byte("Jesus Christ")
}

func (r *RealRenderer) renderMultiline(md []byte, config *RenderConfig) []byte {
	match := r.multiline.FindSubmatch(md)
	if match == nil {
		return md
	}

	var html []byte
	if len(match) == 3 {
		html = r.renderMultiHTML(match[1], nil, match[2], config)
	} else if len(match) == 4 {
		html = r.renderMultiHTML(match[1], match[2], match[3], config)
	}

	md = bytes.Replace(md, match[0], html, 1)
	return r.renderMultiline(md, config)
}

func (r *RealRenderer) renderMultiHTML(tag, param, body []byte, config *RenderConfig) []byte {
	log.Printf("\ntag: %s\nparam: %s\nbody: %s\n", tag, param, body)
	return []byte("Jesus Christ")
}
