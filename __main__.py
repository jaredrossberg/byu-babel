import sys
import os
from byubabel import BYUBabel

# Print error messages to stderr
def eprint(*args, **kwargs):
    print(*args, file=sys.stderr, **kwargs)

# Entrypoint for file run as program
def main():
    input_file = ''
    output_file = ''
    translation = {'x': 0, 'y': 0, 'z': 0}
    
    # Read and save command line arguments
    for i, arg in enumerate(sys.argv):
        if arg == '-i' or arg == '--input':
            if input_file != '':
                eprint('Input file already provided')
                sys.exit(1)
            elif i+1 >= len(sys.argv):
                eprint('Input file must be provided')
                sys.exit(1)
            input_file = sys.argv[i+1]
        elif arg == '-o' or arg == '--output':
            if output_file != '':
                eprint('Output file already provided')
                sys.exit(1)
            elif i+1 >= len(sys.argv):
                eprint('Output file must be provided')
                sys.exit(1)
            output_file = sys.argv[i+1]
        elif arg == '-x':
            if translation['x'] != 0:
                eprint('X axis translation already provided')
                sys.exit(1)
            elif i+1 >= len(sys.argv):
                eprint('X axis translation not provided')
                sys.exit(1)
            translation['x'] = float(sys.argv[i+1])
        elif arg == '-y':
            if translation['y'] != 0:
                eprint('Y axis translation already provided')
                sys.exit(1)
            elif i+1 >= len(sys.argv):
                eprint('Y axis translation not provided')
                sys.exit(1)
            translation['y'] = float(sys.argv[i+1])
        elif arg == '-z':
            if translation['z'] != 0:
                eprint('Z axis translation already provided')
                sys.exit(1)
            elif i+1 >= len(sys.argv):
                eprint('Z axis translation not provided')
                sys.exit(1)
            translation['z'] = float(sys.argv[i+1])
        
    while input_file == '':
        input_file = input('Input file: ')
    if output_file == '':
        yes_or_no = input('No output file provided. Write to input file? (y/n): ')
        if yes_or_no.lower() == 'y' or yes_or_no.lower() == 'yes':
            output_file = input_file
        else:
            print('No output provided. Exiting...')
            sys.exit(0)

    
    config_file = os.path.dirname(__file__) + '/configs/' + os.path.splitext(os.path.basename(input_file))[0] + '.json'
    if not os.path.exists(config_file):
        print('Custom config file not found. Using default config')
        config_file = os.path.dirname(__file__) + '/config.json'

    # Run program on file and print to output file
    try:
        BYUBabel(input_file, config_file, translation).calculate().write_reaction(output_file)
    except Exception as err:
        print(err)
        print(err.__traceback__)
        sys.exit(1)

    sys.exit(0)

if __name__ == "__main__":
    main()
