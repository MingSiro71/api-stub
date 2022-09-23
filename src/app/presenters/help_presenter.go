package presenters

import (
	"fmt"
	"net/http"
)

type HelpPresenter interface {
	success()
}

type helpPresenter struct {
	w http.ResponseWriter 
}

func NewHelpPresenter(w http.ResponseWriter) (*helpPresenter) {
	return &helpPresenter{w: w}
}

func (p *helpPresenter) Success() {
	html := `
	<html>
	<head>
	<title>help of api-stub</title>
	</head>
	<body>
		<h1>api-stub</h1>
		<ul>
			<li>
				<h2>Set</h2>
				<p>url: /set/{your id}</p>
				<p>Set new stub response to your queue.<p>
			<li>
		</ul>
	</body>
	</html>
	`
	fmt.Fprint(p.w, html)
}
