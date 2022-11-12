# Grundlagen

## Typen
| Type           | Description                       |
| -------------- | --------------------------------- |
| Integer        | Whole numbers                     |
| Floating Point | Decimal numbers                   |
| String         | Text                              |
| Boolean        | "truthy" values                   |
| None           | Special value for null / no value |

## Operatoren
| Operator | Description               | Types                    |
| -------- | ------------------------- | ------------------------ |
| +        | Addition                  | Integer, Float, String   |
| -        | Subtraction               | Integer, Float           |
| *        | Multiplication            | Integer, Float, (String) |
| /        | Division                  | Integer, Float           |
| //       | Quotient / Floor Division | Integer, Float           |
| %        | Remainder                 | Integer, Float           | 
| **       | Power                     | Integer, Float           |
| and      | Logical AND | Boolean |
| or       | Logical OR  | Boolean | 
| not      | Logical NOT | Boolean |
| ==       | Equal                 | Integer, Float, String |
| is       | Equal                 | Boolean, None          | 
| !=       | Not equal             | Integer, Float, String |
| is not   | Not equal             | Boolean, None          | 
| <        | Less than             | Integer, Float         |
| <=       | Less or equal than    | Integer, Float         | 
| >        | Greater than          | Integer, Float         | 
| >=       | Greater or equal than | Integer, Float         |

## Comments
```python
# let's define a variable
my_variable = 10 # this is an integer
```

## Flow Control
### If
```python
if condition:
    pass
elif condition:
    pass
elif condition:
    pass
else:
    pass
```
### While
```python
while condition:
    pass
```

```python
while True:  # endless looping
    ...
    if some_other_condition:
        break  # stop the loop
```

```python
while condition:
    ...
    if condition:
        continue  # jump to the next step of the loop

    # code here isn't executed if the "continue" was run
    ...
```

### For
```python
for variable in statement:
    ...
```

### Imports
#### Standard Import
```python
import utils
```
#### Import specific Fuction
```python
from utils import bla
```
#### Rename import
```python
from utils import bla as blafunc
```

### Naming Convention


## Unterlagen und Quellen
- [Invoking the shell](https://docs.python.org/3/tutorial/interpreter.html)
- [Types and operators](https://docs.python.org/3/library/stdtypes.html)
- [Expressions](https://docs.python.org/3/reference/expressions.html)
- [Basic Data Types in Python](https://realpython.com/python-data-types/)
- [Variables in Python](https://realpython.com/python-variables/)
- [Operators and Expressions in Python](https://realpython.com/python-operators-expressions/)
