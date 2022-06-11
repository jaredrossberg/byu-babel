
if __name__ == "__main__":
    input_str = ' '
    lines = []
    while input_str != 'quit' and input_str != 'exit' and input_str != '':
        input_str = input('Enter line: ')
        for x in input_str.split():
            lines.append(x)
    for i in range(len(lines)//3):
        print('{')
        print('\t\"atom1\": {},'.format(lines[i*3]))
        print('\t\"atom2\": {},'.format(lines[i*3+1]))
        print('\t\"bond-strength\": {}'.format(lines[i*3+2]))
        print('},')
