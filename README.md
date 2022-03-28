# Calc, async(signer), uniq

# Part 1

## Part 1. Uniq

A utility has been implemented with which you can display or filter
repeated lines in a file (analogous to the UNIX `uniq` utility). And recurring
input strings must not be recognized unless they strictly follow each other.
The utility itself has a set of parameters that must be supported.

### Parameters

`-c` - count the number of occurrences of the string in the input.
Output this number before the string separated by a space.

`-d` - output only those lines that are repeated in the input.

`-u` - output only those lines that are not repeated in the input.

`-f num_fields` Ignore the first `num_fields` fields in a line.
A field in a string is a non-empty set of characters separated by a space.

`-s num_chars` ignore the first `num_chars` characters in the string.
When used with the `-f` option, first characters are counted
after `num_fields` fields (ignoring space delimiter after
last field).

`-i` - do not take into account the case of letters.

### Usage

`uniq [-c | -d | -u] [-i] [-f num] [-s chars] [input_file [output_file]]`

1. All parameters are optional. Utility behaviors without parameters --
   simple derivation of unique strings from the input.

2. Parameters c, d, u are interchangeable. Should be considered,
   that in parallel these parameters do not make any sense. At
   passing one along with the other needs to be displayed to the user
   proper use of the utility

3. If input_file is not passed, then consider stdin as the input stream

4. If output_file is not passed, then consider stdout as the output stream

### Work example

<details>
    <summary>No parameters</summary>

```bash
$cat input.txt
I love music.
I love music.
I love music.

I love music of Kartik.
I love music of Kartik.
thanks.
I love music of Kartik.
I love music of Kartik.
$cat input.txt | go run uniq.go
I love music.

I love music of Kartik.
thanks.
I love music of Kartik.
```

</details>

<details>
    <summary>With input_file</summary>

```bash
$cat input.txt
I love music.
I love music.
I love music.

I love music of Kartik.
I love music of Kartik.
thanks.
I love music of Kartik.
I love music of Kartik.
$go run uniq.go input.txt
I love music.

I love music of Kartik.
thanks.
I love music of Kartik.
```

</details>

<details>
    <summary>With input_file and output_file parameters</summary>

```bash
$cat input.txt
I love music.
I love music.
I love music.

I love music of Kartik.
I love music of Kartik.
thanks.
I love music of Kartik.
I love music of Kartik.
$go run uniq.go input.txt output.txt
$cat output.txt
I love music.

I love music of Kartik.
thanks.
I love music of Kartik.
```

</details>

<details>
    <summary>With the -c option</summary>

```bash
$cat input.txt
I love music.
I love music.
I love music.

I love music of Kartik.
I love music of Kartik.
thanks.
I love music of Kartik.
I love music of Kartik.
$cat input.txt | go run uniq.go -c
3 I love music.
one
2 I love music of Kartik.
1 thanks.
2 I love music of Kartik.
```

</details>

<details>
    <summary>With the -d option</summary>

```bash
$cat input.txt
I love music.
I love music.
I love music.

I love music of Kartik.
I love music of Kartik.
thanks.
I love music of Kartik.
I love music of Kartik.
$cat input.txt | go run uniq.go -d
I love music.
I love music of Kartik.
I love music of Kartik.
```

</details>

<details>
    <summary>With the -u option</summary>

```bash
$cat input.txt
I love music.
I love music.
I love music.

I love music of Kartik.
I love music of Kartik.
thanks.
I love music of Kartik.
I love music of Kartik.
$cat input.txt | go run uniq.go -u

thanks.
```

</details>

<details>
    <summary>With the -i option</summary>

```bash
$cat input.txt
I LOVE MUSIC.
I love music.
I LoVe MuSiC.

I love MuSIC of Kartik.
I love music of kartik.
thanks.
I love music of kartik.
I love MuSIC of Kartik.
$cat input.txt | go run uniq.go -i
I LOVE MUSIC.

I love MuSIC of Kartik.
thanks.
I love music of kartik.
```

</details>

<details>
    <summary>With -f num option</summary>

```bash
$cat input.txt
We love music.
I love music.
They love music.

I love music of Kartik.
We love music of Kartik.
thanks.
$cat input.txt | go run uniq.go -f 1
We love music.

I love music of Kartik.
thanks.
```

</details>

<details>
    <summary>With -s num option</summary>

```bash
$cat input.txt
I love music.
A love music.
C love music.

I love music of Kartik.
We love music of Kartik.
thanks.
$cat input.txt | go run uniq.go -s 1
I love music.

I love music of Kartik.
We love music of Kartik.
thanks.
```

</details>

### Testing

You need to test the behavior of the written functionality
with different settings. For testing, you need to write unit tests
for this functionality. Tests are needed both for successful cases,
as well as for the unsuccessful.

## Part 2. Calc

We need to write a calculator that can calculate the expression given to STDIN.

It is enough to implement addition, subtraction, multiplication, division and parentheses support.

Tests are also needed here 🙂 Tests need to cover all operations.

### Work example

```bash
    $ go run calc.go "(1+2)-3"
    0

    $ go run calc.go "(1+2)*3"
    nine
```


# Part 2

------

In this task, we are writing an analogue of the unix pipeline, something like:
```
grep 127.0.0.1 | awk '{print $2}' | sort | uniq -c | sort-nr
```

When the STDOUT of one program is passed as STDIN to another program

But in our case, these roles are performed by channels that we pass from one function to another.

*This is a difficult task, feel free to ask for help, it is not done immediately, but when it clicks in your head, everything becomes very simple*

*Assignment when using lecture materials. Everything you need is in the lecture code*

The task itself essentially consists of two parts.
* Writing an ExecutePipeline function that provides us with pipeline processing of worker functions that do something.
* Writing several functions that consider us some kind of conditional hash sum from the input data

The calculation of the hash sum is implemented by the following chain:
* SingleHash considers the value of crc32(data)+"~"+crc32(md5(data)) (concatenation of two strings through ~), where data is what came to the input (in fact, numbers from the first function)
* MultiHash considers the crc32(th+data)) value (the concatenation of the digit cast to the string and the string), where th=0..5 ( i.e. 6 hashes for each input value ), then takes the concatenation of the results in the order of calculation ( 0..5), where data is what came to the input (and went to the output from SingleHash)
* CombineResults gets all results, sorts (https://golang.org/pkg/sort/), combines sorted result with _ (underscore character) into one string
* crc32 is read through the DataSignerCrc32 function
* md5 is read through DataSignerMd5

What's the catch:
* DataSignerMd5 can only be called once at a time, counts as 10ms. If several start up at the same time, there will be an overheat for 1 second
* DataSignerCrc32, counted as 1 sec
* We have 3 seconds for all calculations.
* If you do it linearly - for 7 elements it will take almost 57 seconds, so you need to somehow parallelize it

The results that are displayed if you send 2 values ​​​​(commented out in the test):

```
0 SingleHash data 0
0 SingleHash md5(data) cfcd208495d565ef66e7dff9f98764da
0 SingleHash crc32(md5(data)) 502633748
0 SingleHash crc32(data) 4108050209
0 SingleHash result 4108050209~502633748
4108050209~502633748 MultiHash: crc32(th+step1)) 0 2956866606
4108050209~502633748 MultiHash: crc32(th+step1)) 1 803518384
4108050209~502633748 MultiHash: crc32(th+step1)) 2 1425683795
4108050209~502633748 MultiHash: crc32(th+step1)) 3 3407918797
4108050209~502633748 MultiHash: crc32(th+step1)) 4 2730963093
4108050209~502633748 MultiHash: crc32(th+step1)) 5 1025356555
4108050209~502633748 MultiHash result: 29568666068035183841425683795340791879727309630931025356555

1 SingleHash data 1
1 SingleHash md5(data) c4ca4238a0b923820dcc509a6f75849b
1 SingleHash crc32(md5(data)) 709660146
1 SingleHash crc32(data) 2212294583
1 SingleHash result 2212294583~709660146
2212294583~709660146 MultiHash: crc32(th+step1)) 0 495804419
2212294583~709660146 MultiHash: crc32(th+step1)) 1 2186797981
2212294583~709660146 MultiHash: crc32(th+step1)) 2 4182335870
2212294583~709660146 MultiHash: crc32(th+step1)) 3 1720967904
2212294583~709660146 MultiHash: crc32(th+step1)) 4 259286200
2212294583~709660146 MultiHash: crc32(th+step1)) 5 2427381542
2212294583~709660146 MultiHash result: 4958044192186797981418233587017209679042592862002427381542

CombineResults
```

Run as `go test -v -race`

