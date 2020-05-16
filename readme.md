# gt: go text
This is a vim like text utility written in golang.

## Code structure
Each package contains a `test` sub package. If the package is mostly an interface with several implementations (the `ui` package for example), then it will contain an `impl` subpackage as well.

## Design

### buf
Contains `Buffer` which has methods to manipulate the text and to obtain a line `string` or lines `[]string` views into the text. All text manipulations should be done through `Buffer` methods.

### gt
Handles editor functions such as opening/closing/saving files and loading configuration.

### ui
Creates an abstraction to interface with Ui libraries. The Ui interface is modeled after the tcell library. 
Abstracting the Ui is important for cross platform support. It enables any graphics platform to be used and allows for easy testing via Ui mocks.

### win
Uses the Ui to render buffers, display metadata, and handle commands.


### v1 roadmap
- all of vim's normal mode commands
- 2 ui: tcell and mock
- config as source file for keymapping to buffer commands
- Go syntax and error highlighting
- Go autocomplete
- Go fmt on save
