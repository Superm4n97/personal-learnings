## Resources:
Shell Scripting - http://www.linuxintro.org/wiki/Shell_scripting_tutorial <br>
Special Parameters - https://www.gnu.org/software/bash/manual/html_node/Special-Parameters.html#Special-Parameters <br>

```shell
read VAR # reads variable VAR
``` 
### Function
two ways to declare a function. The first way
```shell
function foo() {
    ...
}
```
2nd way,
```shell
foo(){
  ...
}
```
#### Calling the function:
The functions can be called simply by their names:
```shell
foo
```
Arguments can be passed after the name:
```shell
foo "arg1"
```
the argument can be caught from the function by:
```shell
foo(){
  '$1 is the arg1'
}
```

### $@, $* and $?
**$@** and **$*** both are used to access the full list of positional parameter. If they are used **without double quotes**, both will behave identical.<br>
**$?** gives the most exit status of most recently executed command.<br>
More about special parameter in bash is given in the resource section above.
