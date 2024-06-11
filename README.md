# Math-Skills

## Objectives

The purpose of this project is to develop a statistical calculator capable of calculating the following statistical measures:

- Average
- Median
- Variance
- Standard Deviation

## Features
1. **[Mean](https://learn.zone01kisumu.ke/git/weakinyi/math-skills/src/branch/master/mathskills/average.go)** This is function that calculates the average of the population.
2. **[Median](https://learn.zone01kisumu.ke/git/weakinyi/math-skills/src/branch/master/mathskills/median.go):** It is the middle value in a set of numbers when they are arranged in order from smallest to largest
3. **[Variance](https://learn.zone01kisumu.ke/git/weakinyi/math-skills/src/branch/master/mathskills/variance.go):** This function checks the variance of the population and calculates the variance of the data using the mean and the data itself.
4. **[Standard Deviation](https://learn.zone01kisumu.ke/git/weakinyi/math-skills/src/branch/master/mathskills/standarddeviation.go):** This function uses the variance function to compute the standard deviation of the population data.

## Instructions

To clone this repository: use the command below on your terminal:
```
git clone >paste the link
```

The program should be able to read data from a file and compute each of the forementioned statistics. The input data will be provided in a file specified as a command-line argument. Each line in the file contains one value, representing a statistical population.

To execute the program, users will use a command below
```bash
go run . data.txt
```
OR:

```bash
go run main.go data.txt
```
After reading the file, your program should compute each of the requested statistics and print the results as follows (values are rounded integers):

The output should look like the example below:
```bash
Average: 35
Median: 4
Variance: 5
Standard Deviation: 65
```


## Testing

This project has test files which test the fuctionality of the functions that calculates the average, median, variance and standard deviation. to run the test file:

```bash
go test -v
```
We run the above command in the console or terminal to test the files that are in the mathskills directory

## Learning Objectives

This project is designed to help learn about:

- Statistics and Mathematics
- Average
- Median
- Variance
- Standard Deviation

This project was done by:
[Wendy Akinyi](https://github.com/Wendy-Tabitha/MathMystique)