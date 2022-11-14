# Strukturen

## Listen
Eine Liste ist eine Folge von Elementen. Der Hauptvorteil von Listen ist, 
dass Sie nicht im Voraus wissen müssen, wie viele Elemente die Liste enthalten wird. 
Listen sind veränderbar: Sie können z. B. Elemente hinzufügen und löschen, und alle Verweise auf dieselbe Liste werden die neuen Daten "sehen".
### Erstellung
```python
numbers = list()
```
### Index
```python
num_list = numbers[0]
# Reverse Access
test = num_list[-1] # -1 = last Item -> -2,..,--10
```
### Add Element
```python
num_list.append(5)

#Specific index
wizards.insert(3, 5)
```

### Remove Element
```python
num_list.remove(5)

#Oder
del num_list[2]
```

### Listen Addieren
```python
concatenated_list = [1,2,3] + num_list
```

### Listen Repetieren
```python
repeated_list = 10 * num_list
```

#### Index Error
```python
num_list = list(range(11))
# Reverse Access
test = num_list[-100]
```
Ergibt einen "IndexError: list index out of range"

### Element in Liste
```python
if bla in num_list:
  #Do stuff
  
if 30 not in num_list:
  #Do stuff
```

### Index von einem Element
```python
element_index = num_list.index("test"):
```

### Iteration
```python
for num in num_list:
  #Do stuff
  print(num)

#With index
for index, num in enumerate(num_list):
  #Do stuff
  print(num) 
```
### Comprehension
```python
number_list_incr = [num+1 for num in num_list]

#Syntax
result = [EXPRESSION for item in original_list if condition]
```

### Slices
       +---+---+---+---+---+---+
       | a | b | c | d | e | f |
       +---+---+---+---+---+---+
Slice: 0   1   2   3   4   5   6
Index:   0   1   2   3   4   5
```python
two_numbers = num_list[:2]
#Oder 
two_numbers = num_list[0:2]
```
## Tuples
Tuples sind immutable!
```python
# List
subjects1 = ["Transformation", "Potions", "Divination"]  # This is a list
# Tubple
subjects2 = ("Transformation", "Potions", "Divination")  # This is a tuple
```

## Dictionaries

### Declare
```python
dict_text = dict()
dict_test = {
    "hair": "black",
    "eyes": "green",
    "feature": "scar",
}
eyecolor = dict_test["eyes"]
dict_test["eyes"] = "brown"
```
### Get Keys
```python
dict_test.keys()
```
### Get Values
```python
dict_test.values()
```
### Get Items
```python
dict_test.items()
```
### Check if key exists
```python
if "key" in dict_test:
...
```
### Get Value (get defined value if it does not exist)
```python
test = dict_test.get("testkey", "value")
```
### Set without overriding
```python
dict_test.setdefault("key", 69)
```
### Delete Item
```python
dict_test.pop("hair")
del dict_test["size"]
```
### Comprehension
```python
new_dict = {key: value for (key, value) in old_dict.items() if condition}
```
### Loop
```python
if key, val in dict_test.items():
...
```

## Sets
A set is an unordered, immutable collection of unique elements.
```python
set_test = {
    "bla",
    "bla1",
...
}
```

## Unterlagen und Quellen
- [Lists](https://docs.python.org/3/tutorial/introduction.html#lists)
- [More on Lists](https://docs.python.org/3/tutorial/datastructures.html#more-on-lists)
- [Lists and Tubles](https://realpython.com/python-lists-tuples/)
- [Facts and myths about Python names and values](https://nedbatchelder.com/text/names.html)
