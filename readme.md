# ge: go edit

## Code structure
Each package contains a `test` sub package. If the package is mostly an interface with several implementations (the `ui` package for example), then it will contain an `impl` subpackage as well.

## Design

### buf
Contains the Buffer struct which has methods to actually manipulate the text and to obtain Line string or Lines []string views into the text.

### ge
Handles editor functions such as opening/closing/saving files and loading configuration.

### ui
Creates an abstraction to interface with Ui libraries. The Ui interface is modeled after the tcell library. 
Abstracting the Ui is important for cross platform support. It enables any graphics platform to be used and allows for easy testing via Ui mocks.

### win
Uses the Ui to render buffers, display metadata, and handle commands.
