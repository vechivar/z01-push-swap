# Push-swap

Push swap is a 2 people early cursus exercise where you have to sort small stacks of numbers using very specific operations. Detailed subject can be found in subject.md. Though the original exercise was not very hard, an extra question required to sort stacks of 100 random numbers in less than 700 operations.

The algorithm we created is explained (in french) in algorithme.md. Our first "naive" algorithm produced around 1500 operations. With a different global algorithm and several micro-optimizations, we were able to consistently solve the 100 numbers sorting with less than 700 operations (10 fails over 10 000 tries).

## How to use

`./build.sh` to compile programs

`./checker --generaterandom N` to generate N random numbers

`./push-swap "[random numbers]"` to generate operations

For example :

```
./checker --generaterandom 100; ARG="$(cat random-100)";
./push-swap "$ARG" | wc -l
./push-swap "$ARG" | ./checker "$ARG"
```
