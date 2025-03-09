# Go simple http interview code sample

### Instructions 
Write a program to perform a GET request on the route htttp://coderbyte.com/api/challenges/json/wizard-list and then calculate the average age of wizards by house. Create a new array of objects with a house property and an average_age property. Ignore any items with an empty string as the house property. At the end convert the average_age property from a float to an integer rounding to the nearest whole number.

Finally, output the array of objects as a JSON string.

### Example Input
`[{"name":"EmilySmith","age":25,"house":"California","country":"USA","wand":{"wood":"Maple","core":"Unicornhair","length":12},"friends":[{"name":"Alice","age":24,"house":"Florida","wand": { "wood": "", "core": "", "length": null },"occupation":"nurse"}],"enemies":["Eve"],"patronus":"Dolphin"}]`

### Example Output
`[{"house":"California","average_age":25.0}]`