 Pokedex-cli an interactive Pokedex Cli tool written in Go
 ==========================================================

This project was created as part of the Boot.dev Backend Developer Course
--------------------------------------------------------------------------

A simple cli tool that allows the user to navigate to different pokemon regions and review the pokemon that reside within them using the PokeAPI. This project was primarily made for me to improve in writing Go for real world projects, and develop a better understanding of interacting with API's and REPL concepts.

In this project, I foused on designing unit tests, improving my documentation and git work flow, as well as developing a clear and maintable go project in the hopes of developing skills needed to undertake new independent project ideas.

Use
------------------

pokedex-cli has a built in help function that explains all the commands available to the user. It can be called by simply typing help.

Current commands are:
```bash
help
map
mapb
explore [location-name]
catch
exit
```

Installation
-----------------
Begin by cloning the repository:
```bash
git clone https://github.com/Hi-Im-Yuri/pokedex-cli
```
Then navigate to the directory where the project was cloned and with go installed:
```bash
go install
```

Run with:
```bash
pokedex-cli
```
