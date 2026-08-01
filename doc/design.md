## create a software that:
get address and tags and save finded text between tags in a file
1. user type in CMD: ./app http://address.com .css_selector
2. app run and scrape the address
3. system save returned string from 2 in a file

## layerd design
core: save text in a file
business: scrape: input: (address_web string) (css_selectors string)
				  output: (result bool)
				  steps: 
UI: get data from user and show result to user
save_text
### core functions
func save_text([][]string) int
steps:  1. ready text
		2. save text
		3. return result

