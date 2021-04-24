package main

const homeHtml = `
<!DOCTYPE html>
<html lang="en" class="h-100">
	<head>
		<meta charset="utf-8">
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<meta http-equiv="X-UA-Compatible" content="IE-edge">
		<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.0.0-beta3/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-eOJMYsd53ii+scO/bJGFsiCZc+5NDVN2yr8+0RDqr0Ql0h+rP48ckxlpbzKgwra6" crossorigin="anonymous">
		<title>Postgremon</title>
	</head>
	<body class="d-flex h-100 text-center text-white bg-dark" style="text-shadow: 0 .05rem .1rem rgba(0, 0, 0, .5);">
		<div class="d-flex w-100 h-100 p-3 mx-auto flex-column" style="max-width: 42em;">
			<header class="mb-auto">
				<div>
					<h3 class="float-md-start mb-0">Postgremon</h3>
				</div>
			</header>
			<main class="px-3">
				<h1>Postgremon</h1>
				<p class="lead">
					A Pokémon search page written in Go, using PostgreSQL&apos;s full text search functionality.
				</p>
				<form action="/" method="GET">
					<input class="form-control" style="color: #fff; background-color: #212529; border-color:#6c757d" placeholder="Search for a Pokémon..." class="form-control" autofocus name="search" maxlength="51" type="text">
				</form>
				<br/>
				<p>
					Try searching for a Pokémon above using one of the following attributes: <br/> Pokédex Number, Name, Type, Description or Abilities:</p>
				</p>
			</main>
			<footer class="mt-auto text-white-50">
				<p>
					Check out the code on <a href="https://github.com/dcolesDEV/postgremon" target="_blank">GitHub</a>
					made by <a href="https://github.com/dcolesDEV" target="_blank">Daniel Coles</a>
				</p>
			</footer>
		</div>
	</body>
</html>
`

const resultsHtml = `
<!DOCTYPE html>
<html lang="en" class="h-100">
	<head>
		<head>
		<meta charset="utf-8">
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<meta http-equiv="X-UA-Compatible" content="IE-edge">
		<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.0.0-beta3/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-eOJMYsd53ii+scO/bJGFsiCZc+5NDVN2yr8+0RDqr0Ql0h+rP48ckxlpbzKgwra6" crossorigin="anonymous">
		<title>Postgremon</title>
	</head>
	<body class="d-flex flex-column h-100">
		<header>
			<nav class="navbar navbar-expand-md navbar-dark fixed-top bg-dark">
				<div class="container-fluid">
					<div class="col-sm-2">
						<a class="navbar-brand" href="/">Postgremon</a>
					</div>
					<div class="col-sm-4 me-auto">
						<form class="d-flex align-middle" style="width: 100%" action="/" method="GET">
							<input class="form-control" autofocus name="search" maxlength=51 type="text" value="{{.Search}}">
						</form>
					</div>
				</div>
			</nav>
		</header>
		<main class="flex-shrink-0" style="padding: 60px 15px 0px">
			<div class="container">
				<h1 class="mt-5">Results</h1>
				{{range .Results}}
				<a href="/details?id={{.Id}}" style="display: block; text-decoration: none; color: black;">
					<div class="row" style="padding: 1.5em; border-bottom: 1px solid #eee">
						<div class="col-sm-2">
							<img class="img-fluid" src="https://assets.pokemon.com/assets/cms2/img/pokedex/full/{{.Id}}.png" alt="{{.Name}}" />
						</div>
						<div class="col-sm-9 col-sm-offset-1">
							<div>
								<h4 class="text-muted">#{{.Id}}</h4>
								<h1>{{.Name}}</h1>
							</div>
							<div class="pt-4">
								<h4>Matched:<br/><span class="text-muted">{{.Snippet}}</span></h4>
							</div>
						</div>
					</div>
				</a>
				{{end}}
			</div>
		</main>
		<footer class="footer mt-auto py-3 bg-light text-center">
			<div class="container">
				<p>
					Check out the code on <a href="https://github.com/dcolesDEV/postgremon" target="_blank">GitHub</a>
					made by <a href="https://github.com/dcolesDEV" target="_blank">Daniel Coles</a>
				</p>
			</div>
		</footer>
		<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.0.0-beta3/dist/js/bootstrap.bundle.min.js" integrity="sha384-JEW9xMcG8R+pH31jmWH6WWP0WintQrMb4s7ZOdauHnUtxwoG2vI5DkLtS3qm9Ekf" crossorigin="anonymous"></script>
	</body>
</html>
`

const detailsHtml = `
<!DOCTYPE html>
<html lang="en" class="h-100">
	<head>
		<head>
		<meta charset="utf-8">
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<meta http-equiv="X-UA-Compatible" content="IE-edge">
		<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.0.0-beta3/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-eOJMYsd53ii+scO/bJGFsiCZc+5NDVN2yr8+0RDqr0Ql0h+rP48ckxlpbzKgwra6" crossorigin="anonymous">
		<title>{{.Name}} | Postgremon</title>
	</head>
	<body class="d-flex flex-column h-100">
		<header>
			<nav class="navbar navbar-expand-md navbar-dark fixed-top bg-dark">
				<div class="container-fluid">
					<div class="col-sm-2">
						<a class="navbar-brand" href="/">Postgremon</a>
					</div>
					<div class="col-sm-4 me-auto">
						<form  class="d-flex align-middle" style="width: 100%" action="/" method="GET">
							<input class="form-control" autofocus name="search" maxlength=51 type="text" value="{{.Search}}">
						</form>
					</div>
				</div>
			</nav>
		</header>
		<main class="flex-shrink-0" style="padding: 60px 15px 0px">
			<div class="container">
				<div class="row">
					<div class="col-sm-4">
						<img class="img-fluid" src="https://assets.pokemon.com/assets/cms2/img/pokedex/full/{{.Id}}.png" alt="{{.Name}}" />
					</div>
					<div class="col-sm-7 offset-sm-1">
						<div class="row py-4">
							<h1>#{{.Id}} - {{.Name}}</h1>
						</div>
						<div class="row">
							<h3>Types: {{.Types}}</h3>
						</div>
						<div class="row">
							<h3>Classification: {{.Classification}}</h3>
						</div>
						<div class="row">
							<h3>Abilities: {{.Abilities}}</h3>
						</div>
					</div>
				</div>
			</div>
		</main>
		<footer class="footer mt-auto py-3 bg-light text-center">
			<div class="container">
				<p>
					Check out the code on <a href="https://github.com/dcolesDEV/postgremon" target="_blank">GitHub</a>
					made by <a href="https://github.com/dcolesDEV" target="_blank">Daniel Coles</a>
				</p>
			</div>
		</footer>
		<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.0.0-beta3/dist/js/bootstrap.bundle.min.js" integrity="sha384-JEW9xMcG8R+pH31jmWH6WWP0WintQrMb4s7ZOdauHnUtxwoG2vI5DkLtS3qm9Ekf" crossorigin="anonymous"></script>
	</body>
</html>
`
