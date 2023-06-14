package webhook

type Message struct {
	Username  string  `json:"username,omitempty"`
	AvatarUrl string  `json:"avatar_url,omitempty"`
	Content   string  `json:"content,omitempty"`
	Embeds    []Embed `json:"embeds,omitempty"`
}

type Embed struct {
	Author      *Author   `json:"author,omitempty"`
	Title       string    `json:"title,omitempty"`
	Url         string    `json:"url,omitempty"`
	Description string    `json:"description,omitempty"`
	Color       uint32    `json:"color,omitempty"`
	Fields      []Field   `json:"fields,omitempty"`
	Thumbnail   *Media    `json:"thumbnail,omitempty"`
	Image       *Media    `json:"image,omitempty"`
	Video       *Media    `json:"video,omitempty"`
	Footer      *Footer   `json:"footer,omitempty"`
	Provider    *Provider `json:"provider,omitempty"`
}

type Author struct {
	Name    string `json:"name,omitempty"`
	Url     string `json:"url,omitempty"`
	IconUrl string `json:"icon_url,omitempty"`
}

type Field struct {
	Name   string `json:"name,omitempty"`
	Value  string `json:"value,omitempty"`
	Inline bool   `json:"inline,omitempty"`
}

type Media struct {
	Url      string `json:"url,omitempty"`
	ProxyUrl string `json:"proxy_url,omitempty"`
	Width    int    `json:"width,omitempty"`
	Height   int    `json:"height,omitempty"`
}

type Footer struct {
	Text    string `json:"text,omitempty"`
	IconUrl string `json:"icon_url,omitempty"`
}

type Provider struct {
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`
}
