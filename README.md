221125
get tmperature +? details
hide api key from main
    supply apiKey from env
supply location as cmd line arg


221208
move get api key from main (? test)
get key out of git repo
more readable/pretty response

TODO (CRITICAL)

Handle locations with spaces
    need all of os.Args
    need to escape spaces in URL - use url.QueryEscape?
    Pass TestFormatURL_EscapesSpacesInLocation
    
TODO (NICE TO HAVE)

Report temperature in Fahrenheit
Report actual location (because it might not be the one the user wanted)

NOT GOING TODO (Out of scope)

Animated weather map
Surf forecast
