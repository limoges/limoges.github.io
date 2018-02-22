---
title: "A history of encodings before UTF-8"
date: 2009-10-12
description: "An explanation of the needs of encoding"
draft: false
categories: ["computers", "history"]
---
Google announced this year that UTF-8 had finally overtook all other character
encoding, accounting for the most used character encoding on the internet.

This is a quite important event for the world, one which may help communication
across the world. Still not sure? Follow me down the character encoding rabbit
hole.

### The internet, circa 1999.
If you've used your computer to browse the internet in the last 10 years, you've
probably visited a page which had all kinds of weird squigly characters. 

If you speak multiple languages, it is likely that you've experienced it many times.

If you spoke only one language, and that language was not english, you would probably
have experienced this daily in the past.

Seeing those weird character sequence means one thing: you've had an experience
with a wrong character encoding.

Of course if you live in North America and only speak english, these experiences
would probably have been minimal. 

### Representing characters
The issue is simple: how do we represent written characters in a computer. By
characters, I mean the alphabets, digits, symbols, written and unwritten.

You see, to have characters on a computer, we take a value, say 65 and give it
a special meaning, say 'A'. That means that when we encounter the value 65 in
a file or in a transmission, we assume that it means the capital letter A. 

Of course, speaking english, we figure out that we only need about a hundred 
symbols to represent english, including numbers and punctuation. Being pragmatic,
we use the the smallest unit available on most computers, a byte, 8 bits.

What we've done is build a character encoding. In that encoding, each value maps
to a particular symbol.

And this is what happened. Systems were built and a standard encoding emerged 
and was used within most computing hardware in North America.

### The internet
And then, the internet happened. Suddenly, homogenous networks from all around
the world were being connected together. Networks with different languages,
different culture, and so, different encodings.

The people around the world had developed their own encoding to support their
language and cultures. You see, while there is an alphabet with 26 characters in
english, there are many languages using characters which are not found in the
english alphabet.

So while the value 64 in english means 'A', in another language, it might mean
something completely different in another language.
(edit: Turns out, most encodings use the same meaning for the first 127 values.
The real problems end up presenting themselves in between 127 and 255 for 
"European languages".)

And there was the problem. A computer could no longer "stupidly" look at a number
and translate it using the table it was using before. Files would need to present
their encoding first to be read properly. And that's if you had access to the 
encoding referred in the file in the first place.

Being a bit counter-intuitive to most, this process was omitted or relied on
guesswork, leading to encoding problems.

## Enter unicode

In the late 80s, a proposal for the creation of an encoding which could be
used languages and cultures was drafted. This character encoding would be the
end of encodings, a universal encoding, one encoding to rule them all.

It would use 2 bytes instead of 1 to encode all of the world's characters.
And this is how the Universal Character Set, UCS-2 was born (2 as in 2 bytes).

Turns out, there were a few problems with UCS-2.
1. It has some complexity with how it dealt with bits
2. It isn't large enough to represent all the characters we need.
3. A good character set might want to be extendable, with newer and older symbols.
4. For people who only need english characters, every file becomes twice as big.
(The last point is really petty, in my opinion, but it turns out to be important.)

And this led to a few different encodings to be designed, in order to attempt
to solve these issues.

## UTF-8
One that stuck, was UTF-8. UTF stands for Unicode Transformation Format, 8 bits.

The strength of the design was that unlike other encodings, UTF-8 was a variable
width character encoding. This means that instead of each symbol being represented
by a single byte, a UTF-8 symbol would be encoded in a string of multiple bytes.

This was interesting concept because it solved all of the issues at once. And yes,
english file would not double in space! This very fact probably made the adoption 
of UTF-8 a much bigger success than its compatriot UCS-2.

## Adoption in across the world
There was still a couple problems with UTF-8 adoption in several regions of the world.
Most notably, the regions which do not use european latin alphabets very much.

While utf-8 is a great technology to communicate across international boudaries,
it was still considered a hassle for many.

1. Unless one uses many languages in a single file, it is not necessary
to use UTF-8 to encode that file. Encoding generally support english characters
and another language, so, unless you wanted, for example, Chinese and French,
you could just stick to your local encoding.

2. Some languages are much more costly to encode than others. For example,
Chinese can be 3 or 4 bytes per character in UTF-8. This argument was strong enough
for english speakers to dislike UTF-16, which doubled the file size. Now imagine
quadrupling the size.

3. increases the bandwidth that you must use to transmit when speaking
one of those more costly languages.

To be fair, point 2 and 3 are the same, but they are very important. This is why 
UTF-8 adoption is still a work in progress. While it is very good at encoding 
everything, it isn't free. In many situations where space is at a premium, it
is much cheaper to just support your local encoding.

Cheers to the success of UTF-8!
