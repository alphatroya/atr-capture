package tags

type Item struct {
	Name  string
	Alias string
}

func RequestTags() []Item {
	return []Item{
		{Alias: "🛍️ Book to buy", Name: "books-to-buy"},
		{Alias: "🍿 Movie", Name: "movies"},
		{Alias: "🤔 Ideas", Name: "ideas"},
		{Alias: "✍️ Blog", Name: "blog"},
		{Alias: "Golang", Name: "golang"},
		{Alias: "iOS", Name: "ios"},
		{Alias: "Swift", Name: "swift"},
		{Alias: "SwiftUI", Name: "swiftui"},
		{Alias: "Rust", Name: "rust"},
	}
}
