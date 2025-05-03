package render

import (
	"bytes"
	"fmt"
	"io"
)

func (r *RealRenderer) renderCustom(md []byte, config *RenderConfig) ([]byte, error) {
	var err error
	md, err = r.renderSingleline(md, config)
	if err != nil {
		return []byte{}, err
	}
	md, err = r.renderMultiline(md, config)
	if err != nil {
		return []byte{}, err
	}

	return md, nil
}

func (r *RealRenderer) renderSingleline(md []byte, config *RenderConfig) ([]byte, error) {
	match := r.singleline.FindSubmatch(md)
	if match == nil {
		return md, nil
	}

	total, tag, param := match[0], match[1], match[2]

	var newMarkdown bytes.Buffer
	switch string(tag) {
	case "include":
		include, err := r.source.LoadInclude(string(param))
		if err != nil {
			return []byte{}, err
		}
		newMarkdown.Write(include)
	default:
		io.WriteString(&newMarkdown, fmt.Sprintf("Tag %s does not exist\n", tag))
	}

	md = bytes.Replace(md, total, newMarkdown.Bytes(), 1)
	return r.renderSingleline(md, config)
}

func (r *RealRenderer) renderMultiline(md []byte, config *RenderConfig) ([]byte, error) {
	match := r.multiline.FindSubmatch(md)
	if match == nil {
		return md, nil
	}

	total, tag, body := match[0], match[1], match[2]

	var newMarkdown bytes.Buffer
	switch string(tag) {
	case "warning":
		io.WriteString(&newMarkdown, "Warning:\n")
		newMarkdown.Write(body)
	default:
		io.WriteString(&newMarkdown, fmt.Sprintf("Tag %s does not exist\n", tag))
		newMarkdown.Write(body)
	}

	md = bytes.Replace(md, total, newMarkdown.Bytes(), 1)
	return r.renderMultiline(md, config)
}
