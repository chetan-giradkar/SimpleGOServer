# SimpleGOServer
Simple Server using Golang
This server is a time server and will return current epoch
[GET] : /time := will return current time epoch in milliseconds
[GEt] : /localtime/{timezone} := will return current local time (The data should be in the format of a 
&nbsp;&nbsp;standard IANA time zone file)
[GEt] : /datetime := will return beautified current time (RFC850 = "Monday, 02-Jan-06 15:04:05 MST")

timezone list :
