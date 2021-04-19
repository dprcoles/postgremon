package main

const homeHtml = `
<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="utf-8">
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<meta http-equiv="X-UA-Compatible" content="IE-edge">
		<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.0.0-beta3/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-eOJMYsd53ii+scO/bJGFsiCZc+5NDVN2yr8+0RDqr0Ql0h+rP48ckxlpbzKgwra6" crossorigin="anonymous">
		<title>Postgremon</title>
	</head>
	<body>
		<div class="container">
			<div class="row" style="text-align:center">
				<h1 style="padding-top: 2em; font-size: 80px">Postgremon</h1>
			</div>
			<div class="row" style="padding-top: 2em">
				<div class="col-sm-4 col-sm-offset-4">
					<form action="/" method="GET">
						<input class="form-control" autofocus name="query" maxlength=51 type="text">
					</form>
				</div>
			</div>
			<div class="row" style="text-align: center; padding-top: 4em">
				<div class="col-sm-6 col-sm-offset-4">
						Check out the code on <a href="https://github.com/dcolesdev/postgremon" target="_blank">GitHub</a>
				</div>
			</div>
		</div>
	</body>
</html>
`

const resultsHtml = `
<!DOCTYPE html>
<html lang="en">
	<head>
		<head>
		<meta charset="utf-8">
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<meta http-equiv="X-UA-Compatible" content="IE-edge">
		<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.0.0-beta3/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-eOJMYsd53ii+scO/bJGFsiCZc+5NDVN2yr8+0RDqr0Ql0h+rP48ckxlpbzKgwra6" crossorigin="anonymous">
		<title>Postgremon</title>
	</head>
	<body>
		<div class="container-fluid">
			<div class="row" style="text-align: center; background-color: #e0e0e0; padding: 1em 0">
				<div class="col-sm-1" style="font-size: 32px">
					<a href="/" style="text-decoration: none; color: #000">Postgremon</a>
				</div>
				<div class="col-sm-4">
					<form action="/" method="GET>
						<input class="form-control" autofocus name="query" maxlength=51 type="text" value"{{.Query}}">
					</form>
				</div>
			</div>
			{{range .Results}}
			<div class="row" style="padding: 1.5em; border-bottom: 1px solid #eee">
				<div class="col-sm-11 col-sm-offset-1">
					<div style="font-size: 8px">
						<a href="/details?id={{.Id}}">#{{.Id}} - {{.Name}}</a>
					</div>
					{{.Snippet}}
				</div>
			</div>
			{{end}}
			<div class="row" style="padding: 1.5em; background-color: #e0e0e0">
				Check out the code on <a href="https://github.com/dcolesdev/postgremon" target="_blank">GitHub</a>
			</div>
		</div>
	</body>
</html>
`

const detailsHtml = `
<!DOCTYPE html>
<html lang="en">
	<head>
		<head>
		<meta charset="utf-8">
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<meta http-equiv="X-UA-Compatible" content="IE-edge">
		<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.0.0-beta3/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-eOJMYsd53ii+scO/bJGFsiCZc+5NDVN2yr8+0RDqr0Ql0h+rP48ckxlpbzKgwra6" crossorigin="anonymous">
		<title>#{{.Id}} - {{.Name}} | Postgremon</title>
	</head>
	<body>
		<div class="container-fluid">
			<div class="row" style="text-align: center; background-color: #f1f1f1; padding: 0.5em 0; font-size: 40px">
				#{{.Id}} - {{.Name}}
			</div>
			<div class="row" style="text-align: center; padding 1em 0 0 0; font-size: 24px">
				Types: {{.Types}}
			</div>
			<div class="row" style="text-align: center; padding 1em 0 0 0; font-size: 24px">
				Classification: {{.Classification}}
			</div>
			<div class="row" style="text-align: center; padding 1em 0 0 0; font-size: 24px">
				Abilities: {{.Abilities}}
			</div>
			<div class="row" style="padding: 1.5em; background-color: #f1f1f1">
				Check out the code on <a href="https://github.com/dcolesdev/postgremon" target="_blank">GitHub</a>
			</div>
		</div>
	</body>
</html>
`
