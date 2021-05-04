# SimpleGOServer
Simple Server using Golang
This server is a time server and will return current epoch
[GET] : /time := will return current time epoch in milliseconds
[GEt] : /localtime/{timezone} := will return current local time (https://en.wikipedia.org/wiki/List_of_tz_database_time_zones)
[GEt] : /datetime := will return beautified current time (RFC850 = "Monday, 02-Jan-06 15:04:05 MST")

timezone list :
