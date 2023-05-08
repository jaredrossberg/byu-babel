# BYU-Babel

This program was created by the [Ess Research Lab](https://esslab.byu.edu/) at Brigham Young University to assist with chemistry teaching app [iORA](https://github.com/DanielEss-lab/iORA).

While building iORA, it was found that [OpenBabel](http://openbabel.org/wiki/Main_Page) did not support partial bonds. BYU-Babel aims to replicate the functionality needed from OpenBabel while also supporting partial bonds.

# Usage

1. [Download](https://www.python.org/downloads/) and install the Python programming language

2. Download or clone this repository

3. Open cmd or a terminal and navigate to the folder containing the repository

4. Optionally (but highly recommended), make a copy of `config.json` as `configs/<filename>.json`

    - Note that `<filename>` represents the name of the input file with the file extension removed.

    - If no custom config file is found in [configs/](https://github.com/jaredrossberg/byu-babel/tree/master/configs) the default [config.json](https://github.com/jaredrossberg/byu-babel/blob/master/config.json) file will be used instead.

5. Modify the config file as necessary.

    - `transition-state` is the frame containing the transition state of the reaction
    
    - `window` is the number of frames both before and after the transition state where bonds will begin to form and break
    
    - `frozen-forever` is a list of bonds that will never change during the reaction
    
    - `frozen-before` is a list of bonds that are frozen only before the transition state window
    
    - `frozen-after` is a list of bonds that are frozen only after the transition state window
    
    - `distance-criteria` is a list of atom to atom pairs containing the criteria for a bond to form if it is not frozen

    - If a bonds are forming before or after they are expected or if they are forming the wrong bond types, adjust the distances in `distance-criteria`.
    
    - Note that in order to prevent a bond from being formed between two specific atoms, it must be frozen to a value of `-1`. A value of `0` will have no effect.
    
    - The value of `bond-strength` can be `-1`, `0.5`, `1`, `1.5`, `2`, `2.5`, or `3`. Any other value will have undefined behavor.

    - In order to speed up the creation of custom config files, [to-json-helper.py](https://github.com/jaredrossberg/byu-babel/blob/master/to-json-helper.py) was created. It is optional, but it can be used as described below:
    
        1. In cmd or a terminal, navigate to the folder containing the repository

        2. Run the program: `python to-json-helper.py`

        3. You will be prompted to `Enter before TS`. Open the SDF file containing the reaction and copy and paste the lines containing the bond information of the first frame into the prompt. The lines must be in the format of `<atom 1 number> <atom 2 number> <bond strength>` Any information on each line after those 3 values will be ignored and does not need to be removed. While any frame before the transition state window could be included, it is recommended to use the first frame.

        4. If needed, press enter once or twice until you are prompted to `Enter after TS`. In the same format as the previous step, copy and paste the bonds from the last frame of the reaction into the prompt.

        5. Press enter once or twice until prompted to `Enter TS`. Enter the number of the frame containing the transition state. This should usually be zero-indexed.

        6. When prompted to to `Enter window`, enter the approximate number of frames during which the transition takes place. Note that if `x` is entered as the window, the transition begins `x` frames before the transition state for a total window size of `2 * x`.

        7. The output will be copied to the device clipboard, but can also be printed to the console if desired. Note that auto copy to clipboard has only been tested on MacOS, but should also be supported on Windows.

        8. Paste the output into the config file.

6. Run this program using the following example formats: 
    - `python . --input <path/to/input/file> --output <path/to/output/file>`
    - `python __main__.py -i <path/to/input/file> -o <path/to/output/file>`
    - `python . -i <path/to/input/file>`
        - If no output is specified, you will be asked whether to overwrite the input file or to print to stdout.

7. Place the generated SDF file in iORA's [sdfFiles/](https://github.com/DanielEss-lab/iORA/tree/main/iORA/iORA/sdfFiles) folder. 

8. Modify iORA's [ReactionSectionView.swift](https://github.com/DanielEss-lab/iORA/blob/main/iORA/iORA/ReactionSelectionView.swift) to include the new reaction.

9. In XCode, compile and run iORA to test that the reaction shows as expected. Repeat steps 5-9 as needed.

10. If desired, the new build of iORA can be pushed to Apple Testflight from XCode by changing the output target to `Any iOS device (arm64)` and then navigating to `Product` > `Archive` in the menu. Select `Distribute App` and follow the steps to upload the build. The uploaded build can be sent to Testflight groups or submitted to the App Store from [App Store Connect](https://appstoreconnect.apple.com/apps).


## Important Note:
- All included python programs were developed and tested using python3 and should not be run using python2.


## Supported Filetypes
BYU-Babel currently supports the following file formats
- Inputs
    - xyz
    - sdf
- Outputs
    - sdf
