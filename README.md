# Hexifer
```
 ___ ___                    ___  ___  ___            
|   Y   .-----.--.--.______|   .'  _.'  _.-----.----.
|.  1   |  -__|_   _|______|.  |   _|   _|  -__|   _|
|.  _   |_____|__.__|      |.  |__| |__| |_____|__|  
|:  |   |                  |:  |                     
|::.|:. |                  |::.|                     
`--- ---'                  `---'                     
                                          
```
Small file hex dumper written in golang, something small i thought i would do

# about 

there is not much to it, it will log the output to a file labeled `log.txt`<br>
will take filename as a flag <br>
will customize the byte size based off the file so the output does not create a buffer overflow or stack overflow <br> 
will parse name and size of the file before waiting two seconds for user read <br>

# usage 

```
go run hex.go -file <filename/path> or go run hex.go -h <for help> 
```
