
package main

import (
	_ "embed"
	"log"

	webview "github.com/webview/webview_go"
)

type Res struct {
	Sum int `json:"sum"` 
	Diff int `json:"diff"`
}

//go:embed scripts/script.js
var scriptjs string
func main() {

	w := webview.New(true)
	defer w.Destroy()
	w.SetTitle("Hello")
	w.Bind("noop", func() string {
		log.Println("hello")
		return "hello"
	})
	w.Bind("add", func(a, b int) int {
		return a + b
	})
	w.Bind("quit", func() {
		w.Terminate()
	})
	w.Bind("op", func(a, b int) *Res {
		return &Res{
			Sum: a+b,
			Diff: a-b,
		}
	})
	w.SetHtml(`<!doctype html>
		<html>
			<body>hello</body>
			<script defer> ` + scriptjs + `</script>
			<script>
				window.onload = async function() {
					document.body.innerText = ` + "`hello, ${navigator.userAgent}`" + `;
					const res_noop = await noop()
					console.log('noop res', res_noop)
					const result = await add(4, 6)	
					console.log('add result', result)
					const result_c = await op(4, 6)	
					console.log('op result', JSON.stringify(result_c))
					utils.sayHello()
				};
			</script>
		</html>
	)`)
	w.Run()
}
