package main

import (
	"fmt"
	"os"

	"github.com/go-json-experiment/json"
	"github.com/go-json-experiment/json/jsontext"
)

func main() {
	data := []byte(`[
	{ "width": 5 },
	{ "height": 10 },
	{ "name": "example" }
]`)

	var attrs []Attr

	opts := json.WithUnmarshalers(json.UnmarshalFuncV2(unmarshalAttr))

	err := json.Unmarshal(data, &attrs, opts)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
	fmt.Println(attrs)
}

func unmarshalAttr(dec *jsontext.Decoder, v *Attr, opts json.Options) error {
	var proxy struct {
		Width  *int    `json:"width"`
		Height *int    `json:"height"`
		Name   *string `json:"name"`
	}
	err := json.UnmarshalDecode(dec, &proxy)
	if err != nil {
		return err
	}
	switch {
	case proxy.Width != nil:
		*v = Width(*proxy.Width)
	case proxy.Height != nil:
		*v = Height(*proxy.Height)
	case proxy.Name != nil:
		*v = Name(*proxy.Name)
	}
	return nil
}

type Attr interface {
	isAttr()
}

type Width int

func (Width) isAttr() {}

type Height int

func (Height) isAttr() {}

type Name string

func (Name) isAttr() {}

func (n Name) String() string {
	return string(n)
}
