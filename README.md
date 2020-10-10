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
1. Process any number
1. Detect Errors
    1. I first worked finding unrecognized characters, or going backwards (like `IM`)
    1. The I worked the more than three units `IIII` or `CCCC` for instance
    1. Then I focused in managing wrong minus one structures like `IVIV` or `IIX`
    1. At this point I do believe that code works, goes through the test but it's a bit ugly.