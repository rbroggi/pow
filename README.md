# pow
 
Proof of Work in Go

This is a simple implementation of a proof of work algorithm in Go. 

The algorithm implemented solves the following problem:

Given a difficulty level `y` where `y` is the number of leading zeros bits in the resulting hash, 
find a nonce `x` such that the hash of the nonce concatenated with the input block has `y` leading zeros.

## Complexity of the algorithm

As `y` grows the probability of finding `x` by picking a random nonce decreases exponentially on every independent try.

To estimate the number of tries needed to find a valid nonce, we start by computing the probability of finding a valid nonce on 
an independent try:

The probability is equal to the number of valid hashes divided by the total number of possible hashes.

$$p = \frac{2^{n-y}}{2^n} = \frac{1}{2^y}$$

The cumulative probability $\delta$ of finding the nonce after $k$ tries is given by the formula[^1]:

$$\delta = 1 - (1 - p)^k$$

By solving for $k$ we get:

$$k = \lceil\frac{\log(1 - \delta)}{\log(1 - p)}\rceil$$

Which we can easily put into a spreadsheet and compute some numbers:

| y \\ delta | 50,00%         | 70,00%           | 80,00%           | 90,00%           | 95,00%           |
| ---------- | -------------- | ---------------- | ---------------- | ---------------- | ---------------- |
| 1          | 1,00           | 2,00             | 3,00             | 4,00             | 5,00             |
| 2          | 3,00           | 5,00             | 6,00             | 9,00             | 11,00            |
| 3          | 6,00           | 10,00            | 13,00            | 18,00            | 23,00            |
| 4          | 11,00          | 19,00            | 25,00            | 36,00            | 47,00            |
| 5          | 22,00          | 38,00            | 51,00            | 73,00            | 95,00            |
| 6          | 45,00          | 77,00            | 103,00           | 147,00           | 191,00           |
| 7          | 89,00          | 154,00           | 206,00           | 294,00           | 382,00           |
| 8          | 178,00         | 308,00           | 412,00           | 589,00           | 766,00           |
| 9          | 355,00         | 616,00           | 824,00           | 1.178,00         | 1.533,00         |
| 10         | 710,00         | 1.233,00         | 1.648,00         | 2.357,00         | 3.067,00         |
| 11         | 1.420,00       | 2.466,00         | 3.296,00         | 4.715,00         | 6.134,00         |
| 12         | 2.839,00       | 4.931,00         | 6.592,00         | 9.431,00         | 12.270,00        |
| 13         | 5.678,00       | 9.863,00         | 13.184,00        | 18.862,00        | 24.540,00        |
| 14         | 11.357,00      | 19.726,00        | 26.369,00        | 37.725,00        | 49.081,00        |
| 15         | 22.713,00      | 39.452,00        | 52.738,00        | 75.450,00        | 98.163,00        |
| 16         | 45.426,00      | 78.903,00        | 105.476,00       | 150.902,00       | 196.327,00       |
| 17         | 90.852,00      | 157.807,00       | 210.952,00       | 301.804,00       | 392.656,00       |
| 18         | 181.705,00     | 315.614,00       | 421.904,00       | 603.608,00       | 785.312,00       |
| 19         | 363.409,00     | 631.228,00       | 843.809,00       | 1.207.217,00     | 1.570.625,00     |
| 20         | 726.818,00     | 1.262.457,00     | 1.687.618,00     | 2.414.435,00     | 3.141.252,00     |
| 21         | 1.453.635,00   | 2.524.914,00     | 3.375.236,00     | 4.828.870,00     | 6.282.505,00     |
| 22         | 2.907.270,00   | 5.049.828,00     | 6.750.472,00     | 9.657.741,00     | 12.565.011,00    |
| 23         | 5.814.540,00   | 10.099.656,00    | 13.500.943,00    | 19.315.483,00    | 25.130.023,00    |
| 24         | 11.629.080,00  | 20.199.312,00    | 27.001.887,00    | 38.630.967,00    | 50.260.046,00    |
| 25         | 23.258.160,00  | 40.398.623,00    | 54.003.775,00    | 77.261.934,00    | 100.520.094,00   |
| 26         | 46.516.320,00  | 80.797.247,00    | 108.007.550,00   | 154.523.869,00   | 201.040.189,00   |
| 27         | 93.032.640,00  | 161.594.494,00   | 216.015.100,00   | 309.047.739,00   | 402.080.378,00   |
| 28         | 186.065.280,00 | 323.188.989,00   | 432.030.200,00   | 618.095.479,00   | 804.160.758,00   |
| 29         | 372.130.559,00 | 646.377.977,00   | 864.060.400,00   | 1.236.190.958,00 | 1.608.321.517,00 |
| 30         | 744.261.118,00 | 1.292.755.955,00 | 1.728.120.799,00 | 2.472.381.917,00 | 3.216.643.035,00 |


## References

[^1]: [Bernoulli trial](https://en.wikipedia.org/wiki/Bernoulli_trial)
