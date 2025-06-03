package message

import "fmt"

func ComingSoonTextBody(menu string) string {
	return fmt.Sprintf(`
<b>%s is currently not available.</b>

For more information or inquiries, feel free to email us at hello@astta.xyz.

Thank you for your patience!
`, menu)
}
