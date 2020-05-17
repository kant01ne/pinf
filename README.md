# pinf

`pinf` is a tool that will **P**ipe an input to the terminal output **I**f **N**ot **F**ound in the file given as an argument. 

# Usage example

```
▶ cat "things.txt"
Zero
One
Two
Three

▶ cat otherthings.txt
One
Four
Five

▶ cat things.txt | pinf otherthings.txt
Zero
Two
Three

▶ cat otherthings.txt
One
Four
Five
```

Note that Pinf will not add new lines to the input file. If that's what you want to achieve, you can use [anew](github.com/tomnomnom/anew).


# Install

```
▶ go get -u github.com/NkxxkN/pinf
```

# Credits

This tool was 100% inspired from `anew` from [TomNomNom](https://github.com/tomnomnom).