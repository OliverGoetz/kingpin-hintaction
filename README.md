# kingpin-hintaction

Demo how hintaction breaks arg order

Both commands have two arguments with a HintAction at the first one. 

The first HintAction returns a static slice and the arguments stay in the intended order.

The second HintAction returns a slice from a variable and the order of arguments is broken.

See the following output:

```
$ ./kingpin-hintaction 
usage: kingpin-hintaction [<flags>] <command> [<args> ...]

demo how hintaction breaks arg order

Flags:
  -h, --help        Show context-sensitive help (also try --help-long and --help-man).
      --argset="x"  set of arguments to use
      --version     Show application version.

Commands:
  help [<command>...]
    Show help.

  correctorder <arg1> <arg2>
    Command with correct arg order

  brokenorder <barg2> <barg1>
    Command with broken arg order
```
