package main

const temp1 = `{{.TotalCount}} issues:
{{range .Item}}-----------------------------------
Number: {{.Numebr}}
User: {{.User.Login}}
Title: {{.Title | printf "%.64s"}}
Age: {{.CreateAt | daysAgo}} days
{{end}}`
