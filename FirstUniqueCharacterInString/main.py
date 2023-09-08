def isUnique(string: str) -> bool:
    # we will need to check if the string is greater than 128 in ASCII value
    # the reason for this is because, if the length is greater than 128, it means
    # that there are repeating characters within the string.
    # If the length of the string is greater than 128, we will return false
    if len(string) > 128:
        return False
    
    # We will need to create an array of character flags
    # This flag will represent the ASCII value of the character using the ord()
    # Doing so will allow the alogorithm to check if a character already exist
    char_flags = [False] * 128
    
    # Iterate over the string for each of the character
    for char in string:
        ascii_val = ord(char)
        # If it does, we need to return False, which means it is not unqiue
        # Note that ascii_val is an integer, which means char_flags[ascii_val] gets the value of either True or False on its index
        if char_flags[ascii_val]:
            return False
        char_flags[ascii_val] = True
    return True


print(isUnique("heloWord"))

# check permutation
# for this to return True, we must meet the following criteria
# check if both strings contains the same length
# ALSO check if each character are the same count
def checkPermutation(string1: str, string2: str) -> bool:
    if len(string1) != len(string2):
        return False

    # iterate over string1; if a character already exist in the map,
    # then we increment by 1
    char_dict = {}
    for char in string1:
        if  char in char_dict:
            char_dict[char] += 1
        else:
            char_dict[char] = 1

    # print(char_dict)

    # iterate over string2; if a char already exist in the map, we decrement by 1
    # if char_dict's values returns 0, this would mean that it is a permutation
    # if char from string2 does not exust in char_dict, then we return false
    for char in string2:
        if char in char_dict:
            char_dict[char] -= 1
        else:
            return False
    
    # use the all function to return if All elements in a given iterable are  true
    print(char_dict)
    return all(value == 0 for value in char_dict.values())

print(checkPermutation("abc", "cab"))
print(checkPermutation("abcc", "cabc"))