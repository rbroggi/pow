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

| y \\ delta | 50%         | 70%           | 80%           | 90%           | 95%           |
| ---------- | ----------- | ------------- | ------------- | ------------- | ------------- |
| 1          | 1           | 2             | 3             | 4             | 5             |
| 2          | 3           | 5             | 6             | 9             | 11            |
| 3          | 6           | 10            | 13            | 18            | 23            |
| 4          | 11          | 19            | 25            | 36            | 47            |
| 5          | 22          | 38            | 51            | 73            | 95            |
| 6          | 45          | 77            | 103           | 147           | 191           |
| 7          | 89          | 154           | 206           | 294           | 382           |
| 8          | 178         | 308           | 412           | 589           | 766           |
| 9          | 355         | 616           | 824           | 1.178         | 1.533         |
| 10         | 710         | 1.233         | 1.648         | 2.357         | 3.067         |
| 11         | 1.420       | 2.466         | 3.296         | 4.715         | 6.134         |
| 12         | 2.839       | 4.931         | 6.592         | 9.431         | 12.270        |
| 13         | 5.678       | 9.863         | 13.184        | 18.862        | 24.540        |
| 14         | 11.357      | 19.726        | 26.369        | 37.725        | 49.081        |
| 15         | 22.713      | 39.452        | 52.738        | 75.450        | 98.163        |
| 16         | 45.426      | 78.903        | 105.476       | 150.902       | 196.327       |
| 17         | 90.852      | 157.807       | 210.952       | 301.804       | 392.656       |
| 18         | 181.705     | 315.614       | 421.904       | 603.608       | 785.312       |
| 19         | 363.409     | 631.228       | 843.809       | 1.207.217     | 1.570.625     |
| 20         | 726.818     | 1.262.457     | 1.687.618     | 2.414.435     | 3.141.252     |
| 21         | 1.453.635   | 2.524.914     | 3.375.236     | 4.828.870     | 6.282.505     |
| 22         | 2.907.270   | 5.049.828     | 6.750.472     | 9.657.741     | 12.565.011    |
| 23         | 5.814.540   | 10.099.656    | 13.500.943    | 19.315.483    | 25.130.023    |
| 24         | 11.629.080  | 20.199.312    | 27.001.887    | 38.630.967    | 50.260.046    |
| 25         | 23.258.160  | 40.398.623    | 54.003.775    | 77.261.934    | 100.520.094   |
| 26         | 46.516.320  | 80.797.247    | 108.007.550   | 154.523.869   | 201.040.189   |
| 27         | 93.032.640  | 161.594.494   | 216.015.100   | 309.047.739   | 402.080.378   |
| 28         | 186.065.280 | 323.188.989   | 432.030.200   | 618.095.479   | 804.160.758   |
| 29         | 372.130.559 | 646.377.977   | 864.060.400   | 1.236.190.958 | 1.608.321.517 |
| 30         | 744.261.118 | 1.292.755.955 | 1.728.120.799 | 2.472.381.917 | 3.216.643.035 |

So for a difficulty level of 20 leading zeros, the probability of finding a valid nonce is 80% after 1.687.618 tries.

## References

[^1]: [Bernoulli trial](https://en.wikipedia.org/wiki/Bernoulli_trial)
