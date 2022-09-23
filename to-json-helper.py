
if __name__ == "__main__":
    input_str = ' '
    lines = []
    while input_str != 'quit' and input_str != 'exit' and input_str != '':
        input_str = input('Enter line: ')
        for i, x in enumerate(input_str.split()):
            if i >= 3:
                break
            lines.append(x)
    for i in range(len(lines)//3):
        atom1 =  lines[i*3]
        atom2 =  lines[i*3+1]
        order = lines[i*3+2]
        if int(atom2) < int(atom1):
            atom1, atom2 = atom2, atom1
        print('{')
        print('\t\"atom1\": {},'.format(atom1))
        print('\t\"atom2\": {},'.format(atom2))
        print('\t\"bond-strength\": {}'.format(order))
        print('},')
