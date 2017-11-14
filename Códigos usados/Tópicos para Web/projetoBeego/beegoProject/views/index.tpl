<!DOCTYPE html>

<html>
<head>
  <title>Página inicial</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
</head>

<body>
  <header>
    <h1 class="logo">Minha página</h1>
    <div class="description">
     Descrição
    </div>
  </header>
  <div>
    <h4>Pessoas</h4>
    <p>{{.Pessoas}}</p>
  </div>
  <footer>
    <div class="author">
      Official website:
      <a href="http://{{.Website}}">{{.Website}}</a> /
      Contact me:
      <a class="email" href="mailto:{{.Email}}">{{.Email}}</a>
    </div>
  </footer>
  <div class="backdrop"></div>

  <script src="/static/js/reload.min.js"></script>
</body>
</html>
