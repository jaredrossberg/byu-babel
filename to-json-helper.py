from copy import copy
import subprocess 
import platform

if __name__ == "__main__":
    # Read before TS
    input_str = ' '
    lines = []
    before = []
    print('Enter before TS: ', end='')
    while input_str != '' and input_str != '\n':
        input_str = input()
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
        before.append((atom1, atom2, order))

    # Read after TS
    input_str = ' '
    lines = []
    after = []
    print('Enter after TS: ', end='')
    while input_str != '' and input_str != '\n':
        input_str = input()
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
        after.append((atom1, atom2, order))

    # Find bonds in both
    always = [x for x in before if x in after]
    for x in always:
        before.remove(x)
        after.remove(x)
    
    always.sort()
    before.sort()
    after.sort()

    
    ts = input('Enter TS: ')
    window = input('Enter window: ')

    output = '"transition-state": '+ts+','+'\n'
    output += '"window": '+window+','+'\n'

    # Print always
    output += '"frozen-forever": ['+'\n'
    for i, x in enumerate(always):
        output += '\t{'+'\n'
        output += '\t\t\"atom1\": {},'.format(x[0])+'\n'
        output += '\t\t\"atom2\": {},'.format(x[1])+'\n'
        output += '\t\t\"bond-strength\": {}'.format(x[2])+'\n'
        if i != len(always)-1:
            output += '\t},'+'\n'
        else:
            output += '\t}'+'\n'
    output += '],'+'\n'

    # Print before
    output += '"frozen-before": ['+'\n'
    for i, x in enumerate(before):
        output += '\t{'+'\n'
        output += '\t\t\"atom1\": {},'.format(x[0])+'\n'
        output += '\t\t\"atom2\": {},'.format(x[1])+'\n'
        output += '\t\t\"bond-strength\": {}'.format(x[2])+'\n'
        if i != len(before)-1:
            output += '\t},'+'\n'
        else:
            output += '\t}'+'\n'
    output += '],'+'\n'

    # Print after
    output += '"frozen-after": ['+'\n'
    for i, x in enumerate(after):
        output += '\t{'+'\n'
        output += '\t\t\"atom1\": {},'.format(x[0])+'\n'
        output += '\t\t\"atom2\": {},'.format(x[1])+'\n'
        output += '\t\t\"bond-strength\": {}'.format(x[2])+'\n'
        if i != len(after)-1:
            output += '\t},'+'\n'
        else:
            output += '\t}'+'\n'
    output += '],'+'\n'

    # Copy or print
    copy_command = ''
    if 'Windows' in platform.uname():
        copy_command = 'clip'
    else:
        copy_command = 'pbcopy'
    subprocess.run(copy_command, universal_newlines=True, input=output)
    print_to_terminal = input('\nResult copied to clipboard. Print to terminal? (y/n): ')
    if print_to_terminal == 'y':
        print()
        print(output)

    
    
