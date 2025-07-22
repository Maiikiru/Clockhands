# Clock Angle

<figure>
  <img src="https://media1.giphy.com/media/v1.Y2lkPTc5MGI3NjExbHVidDN6M2hldzZ5ZzRqdDdycGxvbW10Z2toMzF2eWs0MnhjYzVpbyZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9Zw/eKggqtzMsD9OW5HPzV/giphy.gif" alt="Clock with hands moving" style="width:50%">
  <figcaption>Source: <a href="https://media1.giphy.com/media/v1.Y2lkPTc5MGI3NjExbHVidDN6M2hldzZ5ZzRqdDdycGxvbW10Z2toMzF2eWs0MnhjYzVpbyZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9Zw/eKggqtzMsD9OW5HPzV/giphy.gif"> Tenor</a> </figcaption>
</figure>

While I was at work, a fellow coworker told me about their interview problem where they had to find the angle between 2 arbitary clock hands and I thought that
the problem was interesting and pretty easy to implement so I wrote a quick go script to find the solution to the problem.

I think I came up with a pretty performant solution as it uses no trig functions, just basic add/multiply with abs and min. All trivially quick operations.
Technically speaking, since there are only 12 * 60 = 720 possible clock locations, you could also just create a massive precomputed lookup table which would possibly be faster.

## Usage
```
go run clock.go Hour:Minute

>>> go run clock.go 12:15
Your angle between clock hands is 90 degrees

>>> go run clock.go 2:47
Your angle between clock hands is 138 degrees
```
