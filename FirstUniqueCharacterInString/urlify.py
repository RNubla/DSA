def urlify(string: str, trueLength: int) -> str:
    # we will set string to a new version of itself which excludes 
    # extra spaces after the inital space at the end of the string
    # we will need to iterate over the string characters and if we find a 
    # space character, we set that to a %20.
    # at the end of the string, it has a space, we will remove that space and return the final string
    string = list(string[:trueLength])
    index = 0
    for  char in string:
        if char == ' ':
            string[index] = "%20"
        index += 1

    # print("".join(string))
    return "".join(string)

def urlifyBetter(string: str, true_length: int) -> str:
    if true_length > len(string):
        raise ValueError('true_length is greater than the actual string length')

    result = []

    for i in range(true_length):
        if string[i] == ' ':
            result.append('%20')
        else:
            result.append(string[i])

    return "".join(result)

# res = urlify("Mr John Smith        ", 13)
res = urlifyBetter("Mr John Smith        ", 14)

print(res)