package printer

//Color is a string for terminal colors
type Color string

const Red Color = "\033[0;31m"
const BoldRed Color = "\033[1;31m"
const Green Color = "\033[0;32m"
const BoldGreen Color = "\033[1;32m"
const Yellow Color = "\033[0;33m"
const BoldYellow Color = "\033[1;33m"
const Blue Color = "\033[0;34m"
const BoldBlue Color = "\033[1;34m"
const Magenta Color = "\033[0;35m"
const BoldMagenta Color = "\033[1;35m"
const Cyan Color = "\033[0;36m"
const BoldCyan Color = "\033[1;36m"
const ResetColor Color = "\033[0m"
