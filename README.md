# pinf

`pinf` is a tool that will **P**rint the input to the terminal output **I**f **N**ot **F**ound in the file given as an argument. 

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

Note that Pinf will not add new lines to the input file. For this, you can use [anew](github.com/tomnomnom/anew).


# Install

```
▶ go get -u github.com/nkxxkn/pinf
```

# Credits

This tool was greatly inspired from anew from [TomNomNom](https://github.com/tomnomnom).