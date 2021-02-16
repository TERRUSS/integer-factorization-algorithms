from decimal import *
import time


def fermat(n):
	t = time.process_time_ns()

	x = n.sqrt().to_integral_exact(rounding=ROUND_CEILING) #ceil(sqrt(n))
	y = x**2 - n

	it_count = 1
	while not y.sqrt().to_integral_exact(rounding=ROUND_CEILING)**2 == y:
		x += 1
		y = x**2 - n
		
		it_count += 1
	
	return int(x + y.sqrt()), int(x - y.sqrt()), it_count, time.process_time_ns() - t
