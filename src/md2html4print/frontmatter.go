package md2html4print

type frontMatter struct {
	Title string `yaml:"title"`
	Slug  string `yaml:"slug"`
}

func (f frontMatter) SlugStuff() string {
	if f.Slug != "" {
		return f.Slug
	}

	return f.Title
}
