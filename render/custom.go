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
	match := r.singleline.Find(md)
	if match == nil {
		return md
	}

	submatches := r.singleline.FindSubmatch(match)
	html := r.renderSingleHTML(submatches[0], submatches[1], config)

	md = bytes.Replace(md, match, html, 1)
	return r.renderSingleline(md, config)
}

func (r *RealRenderer) renderSingleHTML(tag, param []byte, config *RenderConfig) []byte {
	log.Printf("tag: %s\nparam: %s\n", tag, param)
	return []byte("Jesus Christ")
}

func (r *RealRenderer) renderMultiline(md []byte, config *RenderConfig) []byte {
	match := r.multiline.Find(md)
	if match == nil {
		return md
	}

	submatches := r.multiline.FindSubmatch(match)
	var html []byte
	if len(submatches) == 2 {
		html = r.renderMultiHTML(submatches[0], nil, submatches[1], config)
	} else if len(submatches) == 3 {
		html = r.renderMultiHTML(submatches[0], submatches[1], submatches[2], config)
	}

	md = bytes.Replace(md, match, html, 1)
	return r.renderMultiline(md, config)
}

func (r *RealRenderer) renderMultiHTML(tag, param, body []byte, config *RenderConfig) []byte {
	log.Printf("\ntag: %s\nparam: %s\nbody: %s\n", tag, param, body)
	return []byte("Jesus Christ")
}
