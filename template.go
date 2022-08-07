package main

var defaultPageTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Book</title>

    <style>
    div {
      background-color: lightgrey;
      width: 600px;
      border: 2px solid grey;
      padding: 40px;
      margin: auto;
    }
    </style>

</head>
<body>

    <div>
    <H3>{{ .Title }}</H3>  
        {{ range .Story}}             
           <p>{{.}}</p>           
        {{ end }}   
    <ul>
        {{ range .Options}} 
            <li><a href="{{.Arc}}">{{.Text}}</a></li>          
        {{end}}
    </ul>  
    </div>
</body>
</html>
`
