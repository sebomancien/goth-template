package context

import "context"

const (
	themeKey = "theme"
)

func GetTheme(ctx context.Context) string {
	return ctx.Value(themeKey).(string)
}

func SetTheme(ctx context.Context, value string) context.Context {
	return context.WithValue(ctx, themeKey, value)
}
