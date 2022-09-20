package tests

import (
	"renderer"
	"testing"
)

func Test1(t *testing.T) {
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
}
