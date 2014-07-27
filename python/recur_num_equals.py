#!/usr/bin/python
import sys

#numarr = [7, 4, 3, 9]
numarr = [1, 2, 3, 4, 2, 3, 5]

# 1 sort array
numarr.sort()

# 2 get array sum
_sum = sum(numarr)

# 3 divi 2
if (_sum % 2) == 1:
    print 'NO'
    sys.exit(128)

half_sum = _sum / 2

# 4 recursion
def numarr_equals(numarr, index, value):
    if numarr[index] == value:
        return True
    if index == (len(numarr) - 1):
        return False
    if numarr[index] > value:
        return False
    value_index = numarr[index]
    index = index + 1
    if numarr_equals(numarr, index, value-value_index):
        return True
    if numarr_equals(numarr, index, value):
        return True
    return False

print numarr_equals(numarr, 0, half_sum)

