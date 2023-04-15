<h1 align="center"> Go Reloaded </h1>

This program accesses the content of a file through the file name given as first argument.

It then converts the given byte table to a entire string to be iterated and formated:

Firstly, it fixes spaces by deleting extra spaces.

Secondly, it fixes the parentheses adding spaces oustide and removing spaces inside each pair of parentheses.

Thirdly, it fixes the punctuation: a punctuation should have space on its right and no space on its left, so good so long as there is a word on its left.

Once those three parameters fixed, instances can be applied without difficulty since they now all have the same shape

Instances will be applied only on compatible words before:

- The program will separate the punctuation before applying any instance, so the instance will be applied to the word after the so-fixed punctuation.

- Invalid Instances such as nested instances will be just considered as normal strings.

After applying instances, the program fixes the punctuation once again, due to the side effect of instances deletion.

Then it fixes the apostrophic mark removing space on the odd marks' right and the even marks' left.

And then, it fixes the a/an cases in front of vowel starting word, adding the letter "n" to the "a" in the corresponding case

Finally having fixed all, the program create a new file through the file name given as second argument and write the content of the first file (first argument) in the second file (second argument).

## Usage

```console
    go run . sample.txt result.txt
```
