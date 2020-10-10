# RomanNumeralsKata
Going through the Roman Numerals cata from Coding Dojo in Go.

## Parts

### Part 1. Convert from Digits to Numerals

I've done the next steps (each one with its own commit and following TDD):
1. Throw and error if input is higher than 3000
1. Work with numbers lower than 10.
    1. Refactor for reusing the process of a digit logic
1. Work with numbers lower that 100.
1. Work with any number (a)
    1. Refactor for avoid personalized functions per digit, and move charaters to a constant array.

### Part 2. Convert from Numerals to Digitals
1. Process thousands (Romans having only `M`)
1. Process hundreds (`M`, `D` and `C` supported)