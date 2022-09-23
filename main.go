package main

import (
	"filesystem"
	"renderer"
)

func main() {
	document := renderer.ParseHTMLDocument(`
		<div  id="container" name='QQxQQ' age=18 >
			<ul class="ul">
				<!-- 注释 -->
				<li style="background: red;">1</li>
				<li style="background-color: green;">2</li>
				<li style="background: blue;">3</li>
			</ul>
		</div>
	`)
	println(document)

	cssRules := renderer.ParseCSS(`
		.className {
			margin: 10px;
			padding: 10px;
		}
		#id {
			color: green;
		}
		span {
			background-color: red;
		}
	`)
	println(cssRules)

	resource := filesystem.GetResource("https://www.bilibili.com/")
	println(resource)
}
