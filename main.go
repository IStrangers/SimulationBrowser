package main

import (
	"gitee.com/QQXQQ/Aix/renderer"
)

func main() {
	nodeDOM := renderer.ParseHTML(`
		<div  id="container" name='QQxQQ' age=18 >
			<ul class="ul">
				<li>1</li>
				<li>2</li>
				<li>3</li>
			</ul>
		</div>
	`)
	println(nodeDOM)

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
}
