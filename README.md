# Brick Dataset Generator

The Brick Dataset Generator is a program designed to ease up the making of test-data for the [LEGO brick sorting machine](https://github.com/ElectricCoffee/SW5-Sorting) which was made in conjunction with a 5th semester university project at Aalborg University.

## Files & Their Contents
This program is split into a variety of files, some less obvious than others.

### brick.go
This file contains the `Brick` struct, and related functions.
### input-data.go
This file contains the `InputData` struct, and related functions.
### data-generation.go
This file contains functions responsible for generating the output data.
### utility-functions.go
This file contains various data conversion and convenience functions that don't really belong in the other files.
### main.go
Just has the `main` function.

# Background
As a group, we decided on using a simple interchange format for use in our embedded system that looks like this:

```
COL:<32-bit integer value> LEN:<integer-value>\n
```

The COL value is a representation of what a hexadecimal colour-value would be, had it been written as a decimal value. As an example the colour \#FF00CC would be 16711884

The LEN value is the length of a LEGO-brick in the amount of milliseconds it takes the brick to pass in front of a photointerrupter on the machine's conveyor belt (super weird I know)

So to reliably generate a large amount of bricks, an automated solution was necessary.

# The Program

The generator takes an input [JSON](http://json.org)-file, in the format

```json
{
  "colors": array of number strings,
  "sizes": array of brick sizes in studs between 2 and 8,
  "amount": the number of bricks to be generated
}
```

The colors field accepts string representations of numbers in all the formats that C does. In other words, decimals, octals when prefixed with a 0, and hexadecimals when prefixed with 0x.

The sizes field takes a brick length in studs, "studs" being the little knobs on top of a LEGO brick. All the bricks are assumed to be 2 sutds wide, and n studs long.

If sizes or colors aren't supplied (or left empty), the generator will generate a set of random values that fit within the parameters of the application.

An example of an input file is supplied as part of this repo as "test-data.json".

# Licence
This software is released under an MIT Licence.
