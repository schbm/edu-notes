# APIs

## Pasing Arguments

### Argparse

#### Import
```Python
import argparse
```

#### Standard Setup
```Python
# allow_abbrev --> allows a arguments to like --input to be shortened at input e.g --in
my_parser = argparse.ArgumentParser(prog='myls',
                                    usage='%(prog)s [options] path',
                                    description='List the content of a folder',
                                    epilog='Enjoy the program! :)',
                                    prefix_chars='-',
                                    allow_abbrev=True,
                                    add_help=True)
```

#### Add Arguments
```Python
# Property of namespace var: named by default as the first argument passed to
# .add_argument() for the positional argument and as the long option string for optional arguments. 
# metavar --> value name for the help message
# Positional Argument
my_parser.add_argument('Path',
                       metavar='path',
                       type=str,
                       help='the path to list',
                       action='store',
                       nargs=1,
                       default='default value',
                       required=False,
                       dest='name_of_var_in_namespace')
# Optional Argument (Option)
my_parser.add_argument('-l',
                       '--long',
                       type=int,
                       action='store',
                       help='enable the long listing format',
                       action='store',
                       nargs=1,
                       default=1,
                       choices=range(1, 5),
                       required=False,
                       dest='name_of_var_in_namespace_2')
```

### Actions

- **store stores** the input value to the Namespace object. (This is the default action.)
- **store_const** stores a constant value when the corresponding optional arguments are specified.
- **store_true** stores the Boolean value True when the corresponding optional argument is specified and stores a False elsewhere.
- **store_false** stores the Boolean value False when the corresponding optional argument is specified and stores True elsewhere.
- **append stores** a list, appending a value to the list each time the option is provided.
- **append_const** stores a list appending a constant value to the list each time the option is provided.
- **count stores** an int that is equal to the times the option has been provided.
- **help** shows a help text and exits.
- **version** shows the version of the program and exits.

### Number of Values in Arguments

- ?: a single value, which can be optional
- *: a flexible number of values, which will be gathered into a list
- +: like *, but requiring at least one value
- argparse.REMAINDER: all the values that are remaining in the command line


## Quellen und Unterlagen:
- [How to Build Command Line Interfaces in Python With argparse](https://realpython.com/command-line-interfaces-python-argparse/)
