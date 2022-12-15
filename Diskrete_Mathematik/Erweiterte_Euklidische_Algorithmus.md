
| Step | Description |
|------|-------------|
| 1    | Write the two numbers whose greatest common divisor you want to find as `a` and `b`, with `a` greater than or equal to `b`. |
| 2    | Divide `a` by `b` to obtain a quotient `q` and a remainder `r`. |
| 3    | If `r` is equal to 0, then `b` is the greatest common divisor of `a` and `b`. If `r` is not equal to 0, repeat the same process with `b` and `r` as the two numbers. |
| 4    | Repeat the process until a remainder of 0 is obtained, at which point the previous value of `b` will be the greatest common divisor of `a` and `b`. |


Suppose we want to find the greatest common divisor of 12 and 18. We can apply the steps of the algorithm as follows:

1. Write the two numbers as a and b, with a greater than or equal to b: a = 12, b = 18
2. Divide a by b to obtain a quotient q and a remainder r: 12 / 18 = 0 remainder 12
3. If r is not equal to 0, repeat the same process with b and r as the two numbers: 18 / 12 = 1 remainder 6
4. Repeat the process until a remainder of 0 is obtained: 12 / 6 = 2 remainder 0

In this case, the greatest common divisor of 12 and 18 is 6.

|  a  |  b  |  q  |  r  |  gcd(a, b)  |
|-----|-----|-----|-----|-------------|
| 12  | 18  |  0  |  12 |       -     |
| 18  | 12  |  1  |   6 |       -     |
| 12  |  6  |  2  |   0 |       6     |


