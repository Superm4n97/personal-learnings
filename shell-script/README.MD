## Resources:
Shell Scripting - http://www.linuxintro.org/wiki/Shell_scripting_tutorial <br>
Special Parameters - https://www.gnu.org/software/bash/manual/html_node/Special-Parameters.html#Special-Parameters <br>

### $@, $* and $?
**$@** and **$*** both are used to access the full list of positional parameter. If they are used **without double quotes**, both will behave identical.<br>
**$?** gives the most exit status of most recently executed command.<br>
More about special parameter in bash is given in the resource section above.

### Script Name
```shell
script_name=${0##*/}
```
script_name stores the name of the script

### Notes
* bash script support 3 wildcards:(https://www.howtogeek.com/439199/15-special-characters-you-need-to-know-for-bash/)
  * `?`represent exactly one character after
  * `*` represent any sequence of character including no character
  * `[]` square bracket, the relevant character of the file name must match at least one of the character of the square box.
* a `local` variable can only be accessed with the function, and it's child functions.
* `$(uname -m)` gives the architecture of the system. (example in local_variable.sh file)
* **Environment vs regular shell variables**: Shell variables is local to a particular shell instance, and environment variables are inherited through the all the scripts.