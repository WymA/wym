package internal

const WORD_INDEX_FOOTER string = `
<div class="starter-template">

</div>

</main><!-- /.container -->
<div>
<footer>
  <div class="text-center p-3" style="background-color: rgba(0, 0, 0, 0.2);">
	TO MY DEAR WIFE PHYLLIS and to Toga my cat.© 2024
  </div>
  <!-- Copyright -->
</footer>
</div>
<!-- End of .container -->

<script src="https://cdnjs.cloudflare.com/ajax/libs/jquery/3.6.0/jquery.min.js" referrerpolicy="no-referrer"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/2.11.5/umd/popper.min.js" integrity="sha512-8cU710tp3iH9RniUh6fq5zJsGnjLzOWLWdZqBMLtqaoZUA6AWIE34lwMB3ipUNiTBP5jEZKY95SfbNnQ8cCKvA==" crossorigin="anonymous" referrerpolicy="no-referrer"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/bootstrap/4.6.1/js/bootstrap.min.js" integrity="sha512-UR25UO94eTnCVwjbXozyeVd6ZqpaAE9naiEUBK/A+QDbfSTQFhPGj5lOR6d8tsgbBk84Ggb5A3EkjsOgPRPcKA==" crossorigin="anonymous" referrerpolicy="no-referrer"></script>
<script src="https://matthias2wym.com/assets/js/main.js"></script>
</body>
</html>
`

const WORD_INDEX_HEADER string = `
<!doctype html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    <meta name="description" content="vocabulary builder, 12 new words everyday">
    <meta name="author" content="matthias2wym, wym, matthias, matthias.eth, WymA, vocabulary builder">
    <link rel="icon" href="assets/img/favicon.ico">

    <title>Wym's home | 12 new words everyday</title>

    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/bootstrap/4.6.1/css/bootstrap.min.css" integrity="sha512-T584yQ/tdRR5QwOpfvDfVQUidzfgc2339Lc8uBDtcp/wYu80d7jwBgAxbyMh0a9YM9F8N3tdErpFI8iaGx6x5g==" crossorigin="anonymous" referrerpolicy="no-referrer" />

    <link  rel="stylesheet" href="https://matthias2wym.com/assets/css/cards.css"></style>
    <style>
        body {
        padding-top: 5rem;
        }
        .starter-template {
        padding: 3rem 1.5rem;
        text-align: center;
        }
    </style>
    <!-- Global site tag (gtag.js) - Google Analytics -->
    <script async src="https://www.googletagmanager.com/gtag/js?id=G-3Z4DCRGB21"></script>
    <script>
      window.dataLayer = window.dataLayer || [];
      function gtag(){dataLayer.push(arguments);}
      gtag('js', new Date());

      gtag('config', 'G-3Z4DCRGB21');
    </script>
    {{range .Items}}
    <meta name="{{.Word}}" content="{{.Definition}}">
    {{end}}
  </head>

  <body>
    <nav class="navbar navbar-expand-md navbar-dark bg-dark fixed-top">
      <a class="navbar-brand" href="#">Wym</a>
      <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarsExampleDefault" aria-controls="navbarsExampleDefault" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
      </button>

      <div class="collapse navbar-collapse" id="navbarsExampleDefault">
        <ul class="navbar-nav mr-auto">
          <li class="nav-item active">
            <a class="nav-link" href="https://matthias2wym.com/index.html">Home <span class="sr-only">(current)</span></a>
          </li>
          <li class="nav-item">
            <a class="nav-link" href="https://github.com/WymA">github</a>
          </li>
          <li class="nav-item">
            <a class="nav-link" href="https://matthias2wym.com/history/history.html">History</a>
          </li>
        </ul>
        <!-- <form class="form-inline my-2 my-lg-0">
          <input class="form-control mr-sm-2" type="text" placeholder="Search" aria-label="Search">
          <button class="btn btn-outline-success my-2 my-sm-0" type="submit">Search</button>
        </form> -->
      </div>
    </nav>

    <main role="main" class="container">
`

const DIV_NEW_WORD_ITEM_TEMPLATE = `
      <div class="starter-template">
        <h1>12 new words everyday</h1>
        <h3>{{.TodayDate}}</h2>
      </div>
      <div class="container">
        <div id="new-words" class="row">
        {{range .Items}}
          <div id="{{.Index}}" class="col-sm-4">
            <div class="card card-flip h-100">
              <div class="card-front text-white bg-dark">
                <div class="card-body">
                  <i class="fa fa-search fa-5x float-right"></i>
                  <h3 class="card-title">#{{.Index}}</h3>
                  <p class="card-text">{{.Word}}</p>
                </div>
              </div>
              <div class="card-back bg-white">
                <div class="card-body">
                  <h3 class="card-title">{{.Word}}</h3>
                  <p class="card-text">{{.Definition}}</p>
                </div>
              </div>
            </div>
          </div>
          {{else}}{{end}}
        </div>
      </div>
`

const HISTORY_HEADER string = `
<!doctype html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    <meta name="description" content="vocabulary builder, 12 new words everyday">
    <meta name="author" content="wym, matthias, WymA, vocabulary builder">
    <link rel="icon" href="https://matthias2wym.com/assets/img/favicon.ico">

    <title>Wym's home | 12 new words everyday</title>

    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/bootstrap/4.6.1/css/bootstrap.min.css" integrity="sha512-T584yQ/tdRR5QwOpfvDfVQUidzfgc2339Lc8uBDtcp/wYu80d7jwBgAxbyMh0a9YM9F8N3tdErpFI8iaGx6x5g==" crossorigin="anonymous" referrerpolicy="no-referrer" />

    <link  rel="stylesheet" href="https://matthias2wym.com/assets/css/cards.css"></style>
    <style>
        body {
        padding-top: 5rem;
        }
        .starter-template {
        padding: 3rem 1.5rem;
        text-align: center;
        }
    </style>
    <!-- Global site tag (gtag.js) - Google Analytics -->
    <script async src="https://www.googletagmanager.com/gtag/js?id=G-3Z4DCRGB21"></script>
    <script>
      window.dataLayer = window.dataLayer || [];
      function gtag(){dataLayer.push(arguments);}
      gtag('js', new Date());

      gtag('config', 'G-3Z4DCRGB21');
    </script>
  </head>

  <body>
    <nav class="navbar navbar-expand-md navbar-dark bg-dark fixed-top">
      <a class="navbar-brand" href="#">Wym</a>
      <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarsExampleDefault" aria-controls="navbarsExampleDefault" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
      </button>

      <div class="collapse navbar-collapse" id="navbarsExampleDefault">
        <ul class="navbar-nav mr-auto">
          <li class="nav-item active">
            <a class="nav-link" href="https://matthias2wym.com/index.html">Home <span class="sr-only">(current)</span></a>
          </li>
          <li class="nav-item">
            <a class="nav-link" href="https://github.com/WymA">github</a>
          </li>
          <li class="nav-item">
            <a class="nav-link" href="https://matthias2wym.com/history/history.html">History</a>
          </li>
        </ul>
        <!-- <form class="form-inline my-2 my-lg-0">
          <input class="form-control mr-sm-2" type="text" placeholder="Search" aria-label="Search">
          <button class="btn btn-outline-success my-2 my-sm-0" type="submit">Search</button>
        </form> -->
      </div>
    </nav>

    <main role="main" class="container">
`

const DIV_ALL_HISTORY_ITEMS = `
  <div class="starter-template">
    <h1>History</h1>
  </div>
  <div class="container">
    <div class="row">
      {{range .Items}}
      <div class="col-lg-3 col-md-4 col-sm-6 mb-3">
        <a href="` + BASE_HISTORY_URL + `{{.Filename}}.html" class="d-block p-3 border rounded text-center" style="font-size: 1.25rem; text-decoration: none;">{{.Date}}</a>
      </div>
      {{else}}{{end}}
    </div>
  </div>
`

const BASE_HISTORY_URL = "https://matthias2wym.com/history/"

const PUBLIC_PATH = "../public/"
const HISTORY_PATH = "../public/history/"
