# Text processing kata

## Introduction
As a developer that writes blog posts I want a tool that helps me to understand better the text I am writing. For that I need a way to know the following:

What are the most common words used in the text?
1. How many characters does the text have?
2. Here goes an example of an interface to get started with:

```
interface Processor {
   analyse(text: string);
}
```
The usage of such interface is not required.

## First challenge
Given the following text:
```
Hello, this is an example for you to practice. You should grab
this text and make it as your test case.
```

The output should be:
```
Those are the top 10 words used:

1. you
2. this
3. your
4. to
5. text
6. test
7. should
8. practice
9. make
10. it

The text has in total 21 words
```
Some side notes:

1. As you may have noticed, the example assumes that You and you are the same, in other words, it's not case sensitive.
2. There is no order in which the words that have the same number are listed. For example, this and it, appear once, and they are not alphabetically ordered.

Next up, the kata starts to be a bit more complex. Make sure to complete this challenge first before going on into the second.

## Second challenge
Besides the previous features, the text processing also should have:

1. A way to ignore a given piece of text to analyse (For example, a code snippets - the code snippet is in between ` ` `javascript ` ` ` anything inside should be ignored)
2. A way to offer stop words and remove them from the analysis

Given the example for 1, this would be a text with code snippets:
```
Hello, this is an example for you to practice. You should grab
this text and make it as your test case:

` ` `javascript
if (true) {
  console.log('should should should')
}
```
The text processing output should ignore the code snippet. Meaning, that the output should be:
```
Those are the top 10 words used:

1. you
2. this
3. your
4. to
5. text
6. test
7. should
8. practice
9. make
10. it

The text has in total 21 words
```
Note that, the word should is the same, and it does not goes up in the list as should appears four times (more than the word you).

