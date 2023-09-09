# when we compare the two strings, it cannot be greater than length of string + 2
# if this was the case, then we will need to return false
# otherwise we will need to iterate over each of char of string
# for string1, we convert it into a list
# for string2, we iterate over each char and compare it with the string1 list
# if char exist in string1 list, then we remove that char from the list until it is either none is left or 1 is left which will return true; otherwise it will return false

def oneWay(string1: str, string2: str) -> bool :
    longest_string = ""
    shortes_string = ""
    if len(string1) > len(string2):
        longest_string = string1
        shortes_string = string2
    else:
        longest_string = string2
        shortes_string = string1

    char_dict = {}
    dif_list = []
    for char in shortes_string:
        if char in char_dict:
            char_dict[char] += 1
        else:
            char_dict[char] = 1
    
    for char in longest_string:
        if char in char_dict:
            char_dict[char] -= 1
        else:
            dif_list.append(char)

    if len(dif_list) > 1:
        return False
    else:
        return True

def oneWayBetter(string1: str, string2: str) -> bool:
    # If the lengths differ by more than 1, return False
    if abs(len(string1) - len(string2)) > 1:
        return False

    # Pointers for iteration
    i, j = 0, 0
    # Count the differences
    diff_count = 0
    
    while i < len(string1) and j < len(string2):
        # If characters are not the same
        if string1[i] != string2[j]:
            # If we've already found a difference
            if diff_count:
                return False
            diff_count += 1

            # If one string is longer than the other, move the pointer of the longer string
            if len(string1) > len(string2):
                i += 1
            elif len(string1) < len(string2):
                j += 1
            else:  # If the lengths are the same, move both pointers
                i += 1
                j += 1
        else:  # If characters are the same, move both pointers
            i += 1
            j += 1

    return True
# in general, when comparing two string, it is good to use two pointers

res = oneWay("pale", "bake")
print(res)

res = oneWayBetter("pale", "bake")
print(res)