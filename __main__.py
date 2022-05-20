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
        
    while input_file == '':
        input_file = input('Input file: ')
    if output_file == '':
        output_file = input('Output file (press ENTER to print to console): ')

    config_file = os.path.dirname(__file__) + '/config.json'
    
    # Run program on file and print to output file
    try:
        BYUBabel(input_file, config_file).calculate().write_reaction(output_file)
    except Exception as err:
        print(err)
        print(err.__traceback__)
        sys.exit(1)

    sys.exit(0)

if __name__ == "__main__":
    main()
