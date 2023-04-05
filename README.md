# lem-in

### General information

This project is a digital version of an ant farm.
This program (lem-in) reads from a file (describing the ants and the colony) given in the arguments.
Upon successfully finding the quickest path, lem-in will display the content of the file passed as argument and each move the ants make from room to room.

At the beginning of the game, all the ants are in the room `##start`. The goal is to bring them to the room `##end` with as few moves as possible.

Some colonies will have many rooms and many links, but no path between `##start` and `##end`.
Some will have rooms that link to themselves, sending your path-search spinning in circles. Some will have too many/too few ants, no `##start` or `##end`, duplicated rooms, links to unknown rooms, rooms with invalid coordinates and a variety of other invalid or poorly-formatted input.
In this case you will receive an error message.

The results are displayed on the standard output in the following format:

    number_of_ants
    the_rooms
    the_links

    Lx-y Lz-w Lr-o ...

- x, z, r represents the ants numbers (going from 1 to number_of_ants) and y, w, o represents the rooms names.

- A room is defined by `"name coord_x coord_y"`, and like `"Room 1 2", "nameoftheroom 1 6", "4 6 7"`.

- The links are defined by `"name1-name2"` and look like `"1-2"`, `"2-5"`.

#### Example input:

    10    
    ##start
    1 23 3
    2 16 7
    #comment
    3 16 3
    4 16 5
    5 9 3
    6 1 5
    7 4 8
    ##end
    0 9 5
    0-4
    0-6
    1-3
    4-3
    5-2
    3-5
    #another comment
    4-2
    2-1
    7-6
    7-2
    7-4
    6-5

#### Which corresponds to the following representation :

             _________________
            /                 \
      ____[5]----[3]--[1]      |
     /            |    /       |
    [6]---[0]----[4]  /        |
     \   ________/|  /         |
      \ /        [2]/_________/
      [7]_________/


### Technologies

The language used to write the program is **Go** version **1.17** and was written on **linux/amd64**.
- The backend is written in **Go**.
- The code respects the good [practices](https://git.01.kood.tech/root/public/src/branch/master/subjects/good-practices/README.md).
- The program doesn't present an ant farm visualizer.
- Only the [standard Go](https://golang.org/pkg/) packages are used.

### How to test / audit

One by one:
1. `go run . <filename>` (`go run . ./examples/example00.txt`)
2. Compare the result to the given one.

All the audit examples together:
1. `./test.sh`
2. Compare the results.

Audit questions can be found [here](https://github.com/01-edu/public/tree/master/subjects/lem-in/audit).
All the audit examples are in the folder `examples`.

To see how long each run takes, delete `//` from the beginning of the lines `23` and `34` in `main.go`.

### Developers
- Willem Kuningas / *thinkpad*
- Samuel Uzoagba / *suzoagba*