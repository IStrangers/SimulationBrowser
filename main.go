package main

import (
	"gitee.com/QQXQQ/Aix/renderer"
)

func main() {
	nodeDOM := renderer.ParserHTML(`
		<div  id="container" name='QQxQQ' age=18 >
			<ul class="ul">
				<li>1</li>
				<li>2</li>
				<li>3</li>
			</ul>
		</div>
	`)
	println(nodeDOM)

	styleSheet := renderer.ParserCSS(`
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
	println(styleSheet)
}
