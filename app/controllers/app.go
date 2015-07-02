package controllers

import "github.com/revel/revel"
import "net/http"
import "encoding/json"

type App struct {
	*revel.Controller
}

type GiphyImageType struct {
	Url string `json:"url"`
}

type GiphyImages struct {
	FixedHeight GiphyImageType `json:"fixed_height"`
	Original    GiphyImageType `json:"original"`
}

type GiphyFiles struct {
	Images GiphyImages `json:"images"`
}

type GiphyJson struct {
	Data []GiphyFiles `json:"data"`
}

func (c App) Index() revel.Result {
	return c.Render("Hello something")
}

func (c App) Gif(q string) revel.Result {
	resp, err := http.Get("http://api.giphy.com/v1/gifs/search?api_key=dc6zaTOxFJmzC&limit=12&q=" + q)
	if err != nil {
		return c.RenderHtml("error")
	}

	var giphy GiphyJson
	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(&giphy)

	list := giphy.Data
	return c.Render(list, q)
}
